// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identityserver

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/jinzhu/gorm"
	"go.thethings.network/lorawan-stack/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/pkg/identityserver/blacklist"
	"go.thethings.network/lorawan-stack/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/unique"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)


func (is *IdentityServer) createCluster(ctx context.Context, req *ttnpb.CreateClusterRequest) (cls *ttnpb.Cluster, err error) {
	createdByAdmin := is.IsAdmin(ctx)
	if err = blacklist.Check(ctx, req.ClusterID); err != nil {
		return nil, err
	}
	if usrIDs := req.Collaborator.GetUserIDs(); usrIDs != nil {
		if err = rights.RequireUser(ctx, *usrIDs, ttnpb.RIGHT_USER_CLUSTERS_CREATE); err != nil {
			return nil, err
		}
	} else if orgIDs := req.Collaborator.GetOrganizationIDs(); orgIDs != nil {
		if err = rights.RequireOrganization(ctx, *orgIDs, ttnpb.RIGHT_ORGANIZATION_CLUSTERS_CREATE); err != nil {
			return nil, err
		}
	}
	if err := validateContactInfo(req.Cluster.ContactInfo); err != nil {
		return nil, err
	}
	if !createdByAdmin {
		req.Cluster.State = ttnpb.STATE_REQUESTED
	}
	err = is.withDatabase(ctx, func(db *gorm.DB) (err error) {
		cls, err = store.GetClusterStore(db).CreateCluster(ctx, &req.Cluster)
		if err != nil {
			return err
		}
		if err = store.GetMembershipStore(db).SetMember(
			ctx,
			&req.Collaborator,
			cls.ClusterIdentifiers,
			ttnpb.RightsFrom(ttnpb.RIGHT_ALL),
		); err != nil {
			return err
		}
		if len(req.ContactInfo) > 0 {
			cleanContactInfo(req.ContactInfo)
			cls.ContactInfo, err = store.GetContactInfoStore(db).SetContactInfo(ctx, cls.ClusterIdentifiers, req.ContactInfo)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	is.invalidateCachedMembershipsForAccount(ctx, &req.Collaborator)

	return cls, nil
}

func (is *IdentityServer) getCluster(ctx context.Context, req *ttnpb.GetClusterRequest) (cls *ttnpb.Cluster, err error) {
	if err = is.RequireAuthenticated(ctx); err != nil {
		return nil, err
	}
	req.FieldMask.Paths = cleanFieldMaskPaths(ttnpb.ClusterFieldPathsNested, req.FieldMask.Paths, getPaths, nil)
	if err = rights.RequireCluster(ctx, req.ClusterIdentifiers, ttnpb.RIGHT_CLUSTER_ALL); err != nil {
		if ttnpb.HasOnlyAllowedFields(req.FieldMask.Paths, ttnpb.PublicClusterFields...) {
			defer func() { cls = cls.PublicSafe() }()
		} else {
			return nil, err
		}
	}
	err = is.withDatabase(ctx, func(db *gorm.DB) (err error) {
		cls, err = store.GetClusterStore(db).GetCluster(ctx, &req.ClusterIdentifiers, &req.FieldMask)
		if err != nil {
			return err
		}
		if ttnpb.HasAnyField(req.FieldMask.Paths, "contact_info") {
			cls.ContactInfo, err = store.GetContactInfoStore(db).GetContactInfo(ctx, cls.ClusterIdentifiers)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return cls, nil
}

func (is *IdentityServer) listClusters(ctx context.Context, req *ttnpb.ListClustersRequest) (clss *ttnpb.Clusters, err error) {
	req.FieldMask.Paths = cleanFieldMaskPaths(ttnpb.ClusterFieldPathsNested, req.FieldMask.Paths, getPaths, nil)
	var clsRights map[string]*ttnpb.Rights
	if req.Collaborator == nil {
		callerRights, _, err := is.getRights(ctx)
		if err != nil {
			return nil, err
		}
		clsRights = make(map[string]*ttnpb.Rights, len(callerRights))
		for ids, rights := range callerRights {
			if ids.EntityType() == "cluster" {
				clsRights[unique.ID(ctx, ids)] = rights
			}
		}
		if len(clsRights) == 0 {
			return &ttnpb.Clusters{}, nil
		}
	}
	if usrIDs := req.Collaborator.GetUserIDs(); usrIDs != nil {
		if err = rights.RequireUser(ctx, *usrIDs, ttnpb.RIGHT_USER_CLUSTERS_LIST); err != nil {
			return nil, err
		}
	} else if orgIDs := req.Collaborator.GetOrganizationIDs(); orgIDs != nil {
		if err = rights.RequireOrganization(ctx, *orgIDs, ttnpb.RIGHT_ORGANIZATION_CLUSTERS_LIST); err != nil {
			return nil, err
		}
	}
	var total uint64
	ctx = store.WithPagination(ctx, req.Limit, req.Page, &total)
	defer func() {
		if err == nil {
			setTotalHeader(ctx, total)
		}
	}()
	clss = &ttnpb.Clusters{}
	err = is.withDatabase(ctx, func(db *gorm.DB) error {
		if clsRights == nil {
			rights, err := store.GetMembershipStore(db).FindMemberRights(ctx, req.Collaborator, "cluster")
			if err != nil {
				return err
			}
			clsRights = make(map[string]*ttnpb.Rights, len(rights))
			for ids, rights := range rights {
				clsRights[unique.ID(ctx, ids)] = rights
			}
		}
		if len(clsRights) == 0 {
			return nil
		}
		clsIDs := make([]*ttnpb.ClusterIdentifiers, 0, len(clsRights))
		for uid := range clsRights {
			clsID, err := unique.ToClusterID(uid)
			if err != nil {
				continue
			}
			clsIDs = append(clsIDs, &clsID)
		}
		clss.Clusters, err = store.GetClusterStore(db).FindClusters(ctx, clsIDs, &req.FieldMask)
		if err != nil {
			return err
		}
		for i, cls := range clss.Clusters {
			if rights.RequireCluster(ctx, cls.ClusterIdentifiers, ttnpb.RIGHT_CLUSTER_ALL) != nil {
				clss.Clusters[i] = cls.PublicSafe()
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return clss, nil
}

func (is *IdentityServer) updateCluster(ctx context.Context, req *ttnpb.UpdateClusterRequest) (cls *ttnpb.Cluster, err error) {
	if err = rights.RequireCluster(ctx, req.ClusterIdentifiers, ttnpb.RIGHT_CLUSTER_ALL); err != nil {
		return nil, err
	}
	req.FieldMask.Paths = cleanFieldMaskPaths(ttnpb.ClusterFieldPathsNested, req.FieldMask.Paths, nil, getPaths)
	if len(req.FieldMask.Paths) == 0 {
		req.FieldMask.Paths = updatePaths
	}
	if ttnpb.HasAnyField(req.FieldMask.Paths, "contact_info") {
		if err := validateContactInfo(req.Cluster.ContactInfo); err != nil {
			return nil, err
		}
	}
	updatedByAdmin := is.IsAdmin(ctx)

	if !updatedByAdmin {
		for _, path := range req.FieldMask.Paths {
			switch path {
			case "state":
				return nil, errUpdateUserAdminField.WithAttributes("field", path)
			}
		}
	}

	err = is.withDatabase(ctx, func(db *gorm.DB) (err error) {
		cls, err = store.GetClusterStore(db).UpdateCluster(ctx, &req.Cluster, &req.FieldMask)
		if err != nil {
			return err
		}
		if ttnpb.HasAnyField(req.FieldMask.Paths, "contact_info") {
			cleanContactInfo(req.ContactInfo)
			cls.ContactInfo, err = store.GetContactInfoStore(db).SetContactInfo(ctx, cls.ClusterIdentifiers, req.ContactInfo)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cls, nil
}

func (is *IdentityServer) deleteCluster(ctx context.Context, ids *ttnpb.ClusterIdentifiers) (*types.Empty, error) {
	if err := rights.RequireCluster(ctx, *ids, ttnpb.RIGHT_CLUSTER_ALL); err != nil {
		return nil, err
	}
	err := is.withDatabase(ctx, func(db *gorm.DB) error {
		return store.GetClusterStore(db).DeleteCluster(ctx, ids)
	})
	if err != nil {
		return nil, err
	}
	return ttnpb.Empty, nil
}

func (is *IdentityServer) getClusterIdentifiersForAddress(context.Context, *ttnpb.GetClusterIdentifiersForAddressRequest) (*ttnpb.ClusterIdentifiers, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "not implemented")
}

type clusterRegistry struct {
	*IdentityServer
}

func (cr *clusterRegistry) Create(ctx context.Context, req *ttnpb.CreateClusterRequest) (*ttnpb.Cluster, error) {
	return cr.createCluster(ctx, req)
}

func (cr *clusterRegistry) Get(ctx context.Context, req *ttnpb.GetClusterRequest) (*ttnpb.Cluster, error) {
	return cr.getCluster(ctx, req)
}

func (cr *clusterRegistry) List(ctx context.Context, req *ttnpb.ListClustersRequest) (*ttnpb.Clusters, error) {
	return cr.listClusters(ctx, req)
}

func (cr *clusterRegistry) Update(ctx context.Context, req *ttnpb.UpdateClusterRequest) (*ttnpb.Cluster, error) {
	return cr.updateCluster(ctx, req)
}

func (cr *clusterRegistry) Delete(ctx context.Context, req *ttnpb.ClusterIdentifiers) (*types.Empty, error) {
	return cr.deleteCluster(ctx, req)
}

func (cr *clusterRegistry) GetIdentifiersForAddress(ctx context.Context, req *ttnpb.GetClusterIdentifiersForAddressRequest) (*ttnpb.ClusterIdentifiers, error) {
	return cr.getClusterIdentifiersForAddress(ctx, req)
}
