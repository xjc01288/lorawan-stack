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
	"fmt"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/util/test"
	"google.golang.org/grpc"
)

func TestClustersUnauthenticated(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	canManageClusters = func(_ context.Context) bool {
		return false
	}

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClusterRegistryClient(cc)

		_, err := reg.Create(ctx, &ttnpb.CreateClusterRequest{
			Cluster: ttnpb.Cluster{
				ClusterIdentifiers: ttnpb.ClusterIdentifiers{ClusterID: "foo-cls"},
			},
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.Update(ctx, &ttnpb.UpdateClusterRequest{
			Cluster: ttnpb.Cluster{
				ClusterIdentifiers: ttnpb.ClusterIdentifiers{ClusterID: "foo-cls"},
				Name:               "Updated Name",
			},
			FieldMask: types.FieldMask{Paths: []string{"name"}},
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.Delete(ctx, &ttnpb.ClusterIdentifiers{ClusterID: "foo-cls"})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}
	})
}

func TestClustersPermissionDenied(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	canManageClusters = func(_ context.Context) bool {
		return false
	}

	creds := userCreds(defaultUserIdx, "key without rights")

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClusterRegistryClient(cc)

		_, err := reg.Create(ctx, &ttnpb.CreateClusterRequest{
			Cluster: ttnpb.Cluster{
				ClusterIdentifiers: ttnpb.ClusterIdentifiers{ClusterID: "foo-cls"},
			},
		}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Update(ctx, &ttnpb.UpdateClusterRequest{
			Cluster: ttnpb.Cluster{
				ClusterIdentifiers: ttnpb.ClusterIdentifiers{ClusterID: "foo-cls"},
				Name:               "Updated Name",
			},
			FieldMask: types.FieldMask{Paths: []string{"name"}},
		}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Delete(ctx, &ttnpb.ClusterIdentifiers{ClusterID: "foo-cls"}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}
	})
}

func TestClustersCRUD(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	creds := userCreds(adminUserIdx, "") // TODO: Replace with cluster management creds.

	canManageClusters = func(_ context.Context) bool {
		return true
	}

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClusterRegistryClient(cc)

		created, err := reg.Create(ctx, &ttnpb.CreateClusterRequest{
			Cluster: ttnpb.Cluster{
				ClusterIdentifiers: ttnpb.ClusterIdentifiers{ClusterID: "foo"},
				Name:               "Foo Cluster",
			},
		}, creds)

		a.So(err, should.BeNil)
		a.So(created.Name, should.Equal, "Foo Cluster")

		got, err := reg.Get(ctx, &ttnpb.GetClusterRequest{
			ClusterIdentifiers: created.ClusterIdentifiers,
			FieldMask:          types.FieldMask{Paths: []string{"name"}},
		}, creds)

		a.So(err, should.BeNil)
		a.So(got.Name, should.Equal, created.Name)

		_, err = reg.Get(ctx, &ttnpb.GetClusterRequest{
			ClusterIdentifiers: created.ClusterIdentifiers,
			FieldMask:          types.FieldMask{Paths: []string{"name"}},
		})
		a.So(err, should.BeNil)

		updated, err := reg.Update(ctx, &ttnpb.UpdateClusterRequest{
			Cluster: ttnpb.Cluster{
				ClusterIdentifiers: created.ClusterIdentifiers,
				Name:               "Updated Name",
			},
			FieldMask: types.FieldMask{Paths: []string{"name"}},
		}, creds)

		a.So(err, should.BeNil)
		a.So(updated.Name, should.Equal, "Updated Name")

		list, err := reg.List(ctx, &ttnpb.ListClustersRequest{
			FieldMask: types.FieldMask{Paths: []string{"name"}},
		}, creds)
		a.So(err, should.BeNil)
		if a.So(list.Clusters, should.NotBeEmpty) {
			var found bool
			for _, item := range list.Clusters {
				if item.ClusterIdentifiers == created.ClusterIdentifiers {
					found = true
					a.So(item.Name, should.Equal, updated.Name)
				}
			}
			a.So(found, should.BeTrue)
		}

		_, err = reg.List(ctx, &ttnpb.ListClustersRequest{
			FieldMask: types.FieldMask{Paths: []string{"name"}},
		})
		a.So(err, should.BeNil)

		_, err = reg.Delete(ctx, &created.ClusterIdentifiers, creds)
		a.So(err, should.BeNil)
	})
}

func TestClustersPagination(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	creds := userCreds(adminUserIdx, "") // TODO: Replace with cluster management creds.

	canManageClusters = func(_ context.Context) bool {
		return true
	}

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewClusterRegistryClient(cc)

		for i := 0; i < 3; i++ {
			reg.Create(ctx, &ttnpb.CreateClusterRequest{
				Cluster: ttnpb.Cluster{
					ClusterIdentifiers: ttnpb.ClusterIdentifiers{ClusterID: fmt.Sprintf("foo-%d", i)},
					Name:               fmt.Sprintf("Foo Cluster %d", i),
				},
			}, creds)
		}

		list, err := reg.List(test.Context(), &ttnpb.ListClustersRequest{
			FieldMask: types.FieldMask{Paths: []string{"name"}},
			Limit:     2,
			Page:      1,
		}, creds)

		a.So(list, should.NotBeNil)
		a.So(list.Clusters, should.HaveLength, 2)
		a.So(err, should.BeNil)

		list, err = reg.List(test.Context(), &ttnpb.ListClustersRequest{
			FieldMask: types.FieldMask{Paths: []string{"name"}},
			Limit:     2,
			Page:      2,
		}, creds)

		a.So(list, should.NotBeNil)
		a.So(list.Clusters, should.HaveLength, 1)
		a.So(err, should.BeNil)

		list, err = reg.List(test.Context(), &ttnpb.ListClustersRequest{
			FieldMask: types.FieldMask{Paths: []string{"name"}},
			Limit:     2,
			Page:      3,
		}, creds)

		a.So(list, should.NotBeNil)
		a.So(list.Clusters, should.BeEmpty)
		a.So(err, should.BeNil)
	})
}
