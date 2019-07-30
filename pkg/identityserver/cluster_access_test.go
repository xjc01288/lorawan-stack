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
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/util/test"
	"google.golang.org/grpc"
)

func init() {
	clusterAccessUser.Admin = false
	clusterAccessUser.State = ttnpb.STATE_APPROVED
	for _, apiKey := range userAPIKeys(&clusterAccessUser.UserIdentifiers).APIKeys {
		apiKey.Rights = []ttnpb.Right{
			ttnpb.RIGHT_CLUSTER_ALL,
		}
	}
}

func TestClusterAccessRightsPermissionDenied(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		userID, creds := clusterAccessUser.UserIdentifiers, userCreds(clusterAccessUserIdx)
		clusterID := userClusters(&userID).Clusters[0].ClusterIdentifiers
		collaboratorID := collaboratorUser.UserIdentifiers.OrganizationOrUserIdentifiers()

		reg := ttnpb.NewClusterAccessClient(cc)

		_, err := reg.SetCollaborator(ctx, &ttnpb.SetClusterCollaboratorRequest{
			ClusterIdentifiers: clusterID,
			Collaborator: ttnpb.Collaborator{
				OrganizationOrUserIdentifiers: *collaboratorID,
				Rights:                        []ttnpb.Right{ttnpb.RIGHT_ALL},
			},
		}, creds)

		a.So(err, should.NotBeNil)
		a.So(errors.IsPermissionDenied(err), should.BeTrue)
	})
}

func TestClusterAccessPermissionDenied(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		userID := defaultUser.UserIdentifiers
		collaboratorID := collaboratorUser.UserIdentifiers.OrganizationOrUserIdentifiers()
		clusterID := userClusters(&userID).Clusters[0].ClusterIdentifiers

		reg := ttnpb.NewClusterAccessClient(cc)

		rights, err := reg.ListRights(ctx, &clusterID)

		a.So(rights, should.NotBeNil)
		a.So(rights.Rights, should.BeEmpty)
		a.So(err, should.BeNil)

		collaborators, err := reg.ListCollaborators(ctx, &ttnpb.ListClusterCollaboratorsRequest{
			ClusterIdentifiers: clusterID,
		})

		a.So(collaborators, should.BeNil)
		a.So(err, should.NotBeNil)
		a.So(errors.IsUnauthenticated(err), should.BeTrue)

		_, err = reg.SetCollaborator(ctx, &ttnpb.SetClusterCollaboratorRequest{
			ClusterIdentifiers: clusterID,
			Collaborator: ttnpb.Collaborator{
				OrganizationOrUserIdentifiers: *collaboratorID,
				Rights:                        []ttnpb.Right{ttnpb.RIGHT_CLUSTER_ALL},
			},
		})

		a.So(err, should.NotBeNil)
		a.So(errors.IsPermissionDenied(err), should.BeTrue)
	})
}

func TestClusterAccessClusterAuth(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		userID := defaultUser.UserIdentifiers
		clusterID := userClusters(&userID).Clusters[0].ClusterIdentifiers

		reg := ttnpb.NewClusterAccessClient(cc)

		rights, err := reg.ListRights(ctx, &clusterID, is.WithClusterAuth())

		a.So(rights, should.NotBeNil)
		a.So(ttnpb.AllClusterRights.Intersect(ttnpb.AllClusterRights).Sub(rights).Rights, should.BeEmpty)
		a.So(err, should.BeNil)
	})
}

func TestClusterAccessCRUD(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		userID, creds := defaultUser.UserIdentifiers, userCreds(defaultUserIdx)
		collaboratorID := collaboratorUser.UserIdentifiers.OrganizationOrUserIdentifiers()
		clusterID := userClusters(&userID).Clusters[0].ClusterIdentifiers

		reg := ttnpb.NewClusterAccessClient(cc)

		rights, err := reg.ListRights(ctx, &clusterID, creds)

		a.So(rights, should.NotBeNil)
		a.So(rights.Rights, should.Contain, ttnpb.RIGHT_CLUSTER_ALL)
		a.So(err, should.BeNil)

		modifiedClusterID := clusterID
		modifiedClusterID.ClusterID += "mod"

		rights, err = reg.ListRights(ctx, &modifiedClusterID, creds)
		a.So(rights, should.NotBeNil)
		a.So(rights.Rights, should.BeEmpty)
		a.So(err, should.BeNil)

		collaborators, err := reg.ListCollaborators(ctx, &ttnpb.ListClusterCollaboratorsRequest{
			ClusterIdentifiers: clusterID,
		}, creds)

		a.So(collaborators, should.NotBeNil)
		a.So(collaborators.Collaborators, should.NotBeEmpty)
		a.So(err, should.BeNil)

		_, err = reg.SetCollaborator(ctx, &ttnpb.SetClusterCollaboratorRequest{
			ClusterIdentifiers: clusterID,
			Collaborator: ttnpb.Collaborator{
				OrganizationOrUserIdentifiers: *collaboratorID,
				Rights:                        []ttnpb.Right{ttnpb.RIGHT_CLUSTER_ALL},
			},
		}, creds)

		a.So(err, should.BeNil)

		res, err := reg.GetCollaborator(ctx, &ttnpb.GetClusterCollaboratorRequest{
			ClusterIdentifiers:            clusterID,
			OrganizationOrUserIdentifiers: *collaboratorID,
		}, creds)

		a.So(err, should.BeNil)
		a.So(res.Rights, should.Resemble, []ttnpb.Right{ttnpb.RIGHT_CLUSTER_ALL})
	})
}
