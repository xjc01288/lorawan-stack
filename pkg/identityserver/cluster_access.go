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
	"go.thethings.network/lorawan-stack/pkg/email"
	"go.thethings.network/lorawan-stack/pkg/events"
	"go.thethings.network/lorawan-stack/pkg/identityserver/emails"
	"go.thethings.network/lorawan-stack/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/pkg/log"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

var (
	evtUpdateClusterCollaborator = events.Define(
		"cluster.collaborator.update", "update cluster collaborator",
		ttnpb.RIGHT_CLUSTER_ALL,
		ttnpb.RIGHT_USER_CLUSTERS_LIST,
	)
	evtDeleteClusterCollaborator = events.Define(
		"cluster.collaborator.delete", "delete cluster collaborator",
		ttnpb.RIGHT_CLUSTER_ALL,
		ttnpb.RIGHT_USER_CLUSTERS_LIST,
	)
)

func (is *IdentityServer) listClusterRights(ctx context.Context, ids *ttnpb.ClusterIdentifiers) (*ttnpb.Rights, error) {
	clsRights, err := rights.ListCluster(ctx, *ids)
	if err != nil {
		return nil, err
	}
	return clsRights.Intersect(ttnpb.AllClusterRights), nil
}

func (is *IdentityServer) getClusterCollaborator(ctx context.Context, req *ttnpb.GetClusterCollaboratorRequest) (*ttnpb.GetCollaboratorResponse, error) {
	if err := rights.RequireCluster(ctx, req.ClusterIdentifiers, ttnpb.RIGHT_CLUSTER_ALL); err != nil {
		return nil, err
	}
	res := &ttnpb.GetCollaboratorResponse{
		OrganizationOrUserIdentifiers: req.OrganizationOrUserIdentifiers,
	}
	err := is.withDatabase(ctx, func(db *gorm.DB) error {
		rights, err := store.GetMembershipStore(db).GetMember(
			ctx,
			&req.OrganizationOrUserIdentifiers,
			req.ClusterIdentifiers,
		)
		if err != nil {
			return err
		}
		res.Rights = rights.GetRights()
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (is *IdentityServer) setClusterCollaborator(ctx context.Context, req *ttnpb.SetClusterCollaboratorRequest) (*types.Empty, error) {
	// Require that caller has rights to manage collaborators.
	if err := rights.RequireCluster(ctx, req.ClusterIdentifiers, ttnpb.RIGHT_CLUSTER_ALL); err != nil {
		return nil, err
	}
	// Require that caller has at least the rights we're giving to the collaborator.
	if err := rights.RequireCluster(ctx, req.ClusterIdentifiers, req.Collaborator.Rights...); err != nil {
		return nil, err
	}
	err := is.withDatabase(ctx, func(db *gorm.DB) error {
		return store.GetMembershipStore(db).SetMember(
			ctx,
			&req.Collaborator.OrganizationOrUserIdentifiers,
			req.ClusterIdentifiers,
			ttnpb.RightsFrom(req.Collaborator.Rights...),
		)
	})
	if err != nil {
		return nil, err
	}
	if len(req.Collaborator.Rights) > 0 {
		events.Publish(evtUpdateClusterCollaborator(ctx, ttnpb.CombineIdentifiers(req.ClusterIdentifiers, req.Collaborator), nil))
		err = is.SendContactsEmail(ctx, req.EntityIdentifiers(), func(data emails.Data) email.MessageData {
			data.SetEntity(req.EntityIdentifiers())
			return &emails.CollaboratorChanged{Data: data, Collaborator: req.Collaborator}
		})
		if err != nil {
			log.FromContext(ctx).WithError(err).Error("Could not send collaborator updated notification email")
		}
	} else {
		events.Publish(evtDeleteClusterCollaborator(ctx, ttnpb.CombineIdentifiers(req.ClusterIdentifiers, req.Collaborator), nil))
	}
	is.invalidateCachedMembershipsForAccount(ctx, &req.Collaborator.OrganizationOrUserIdentifiers)
	return ttnpb.Empty, nil
}

func (is *IdentityServer) listClusterCollaborators(ctx context.Context, req *ttnpb.ListClusterCollaboratorsRequest) (collaborators *ttnpb.Collaborators, err error) {
	if err = is.RequireAuthenticated(ctx); err != nil { // Cluster collaborators can be seen by all authenticated users.
		return nil, err
	}
	var total uint64
	ctx = store.WithPagination(ctx, req.Limit, req.Page, &total)
	defer func() {
		if err == nil {
			setTotalHeader(ctx, total)
		}
	}()
	err = is.withDatabase(ctx, func(db *gorm.DB) error {
		memberRights, err := store.GetMembershipStore(db).FindMembers(ctx, req.ClusterIdentifiers)
		if err != nil {
			return err
		}
		collaborators = &ttnpb.Collaborators{}
		for member, rights := range memberRights {
			collaborators.Collaborators = append(collaborators.Collaborators, &ttnpb.Collaborator{
				OrganizationOrUserIdentifiers: *member,
				Rights:                        rights.GetRights(),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return collaborators, nil
}

type clusterAccess struct {
	*IdentityServer
}

func (ca *clusterAccess) ListRights(ctx context.Context, req *ttnpb.ClusterIdentifiers) (*ttnpb.Rights, error) {
	return ca.listClusterRights(ctx, req)
}

func (ca *clusterAccess) GetCollaborator(ctx context.Context, req *ttnpb.GetClusterCollaboratorRequest) (*ttnpb.GetCollaboratorResponse, error) {
	return ca.getClusterCollaborator(ctx, req)
}

func (ca *clusterAccess) SetCollaborator(ctx context.Context, req *ttnpb.SetClusterCollaboratorRequest) (*types.Empty, error) {
	return ca.setClusterCollaborator(ctx, req)
}

func (ca *clusterAccess) ListCollaborators(ctx context.Context, req *ttnpb.ListClusterCollaboratorsRequest) (*ttnpb.Collaborators, error) {
	return ca.listClusterCollaborators(ctx, req)
}
