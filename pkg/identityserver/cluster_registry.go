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
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/identityserver/blacklist"
	"go.thethings.network/lorawan-stack/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// TODO: Write func to determine if caller can manage clusters.
var canManageClusters = func(_ context.Context) bool {
	return false
}

var errNoClusterAdminRights = errors.DefinePermissionDenied("no_cluster_admin_rights", "no cluster admin rights")

func (is *IdentityServer) createCluster(ctx context.Context, req *ttnpb.CreateClusterRequest) (cls *ttnpb.Cluster, err error) {
	if err = blacklist.Check(ctx, req.ClusterID); err != nil {
		return nil, err
	}
	if err = is.RequireAuthenticated(ctx); err != nil {
		return nil, err
	}
	if !canManageClusters(ctx) {
		return nil, errNoClusterAdminRights
	}
	if err := validateContactInfo(req.Cluster.ContactInfo); err != nil {
		return nil, err
	}
	err = is.withDatabase(ctx, func(db *gorm.DB) (err error) {
		cls, err = store.GetClusterStore(db).CreateCluster(ctx, &req.Cluster)
		if err != nil {
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
	return cls, nil
}

func (is *IdentityServer) getCluster(ctx context.Context, req *ttnpb.GetClusterRequest) (cls *ttnpb.Cluster, err error) {
	if !canManageClusters(ctx) {
		defer func() { cls = cls.PublicSafe() }()
	}
	req.FieldMask.Paths = cleanFieldMaskPaths(ttnpb.ClusterFieldPathsNested, req.FieldMask.Paths, getPaths, nil)
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
	if !canManageClusters(ctx) {
		defer func() {
			for i, cls := range clss.Clusters {
				clss.Clusters[i] = cls.PublicSafe()
			}
		}()
	}
	req.FieldMask.Paths = cleanFieldMaskPaths(ttnpb.ClusterFieldPathsNested, req.FieldMask.Paths, getPaths, nil)
	var total uint64
	ctx = store.WithPagination(ctx, req.Limit, req.Page, &total)
	defer func() {
		if err == nil {
			setTotalHeader(ctx, total)
		}
	}()
	clss = &ttnpb.Clusters{}
	err = is.withDatabase(ctx, func(db *gorm.DB) (err error) {
		clss.Clusters, err = store.GetClusterStore(db).FindClusters(ctx, nil, &req.FieldMask)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return clss, nil
}

func (is *IdentityServer) updateCluster(ctx context.Context, req *ttnpb.UpdateClusterRequest) (cls *ttnpb.Cluster, err error) {
	if err = is.RequireAuthenticated(ctx); err != nil {
		return nil, err
	}
	if !canManageClusters(ctx) {
		return nil, errNoClusterAdminRights
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
	if err := is.RequireAuthenticated(ctx); err != nil {
		return nil, err
	}
	if !canManageClusters(ctx) {
		return nil, errNoClusterAdminRights
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
