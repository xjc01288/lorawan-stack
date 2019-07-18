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

package store

import (
	"context"
	"fmt"
	"reflect"
	"runtime/trace"
	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/jinzhu/gorm"
	"go.thethings.network/lorawan-stack/pkg/rpcmiddleware/warning"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

// GetClusterStore returns an ClusterStore on the given db (or transaction).
func GetClusterStore(db *gorm.DB) ClusterStore {
	return &clusterStore{store: newStore(db)}
}

type clusterStore struct {
	*store
}

// selectClusterFields selects relevant fields (based on fieldMask) and preloads details if needed.
func selectClusterFields(ctx context.Context, query *gorm.DB, fieldMask *types.FieldMask) *gorm.DB {
	if fieldMask == nil || len(fieldMask.Paths) == 0 {
		return query.Preload("Attributes")
	}
	var clusterColumns []string
	var notFoundPaths []string
	for _, path := range ttnpb.TopLevelFields(fieldMask.Paths) {
		switch path {
		case "ids", "created_at", "updated_at":
			// always selected
		case attributesField:
			query = query.Preload("Attributes")
		default:
			if columns, ok := clusterColumnNames[path]; ok {
				clusterColumns = append(clusterColumns, columns...)
			} else {
				notFoundPaths = append(notFoundPaths, path)
			}
		}
	}
	if len(notFoundPaths) > 0 {
		warning.Add(ctx, fmt.Sprintf("unsupported field mask paths: %s", strings.Join(notFoundPaths, ", ")))
	}
	return query.Select(cleanFields(append(append(modelColumns, "cluster_id"), clusterColumns...)...))
}

func (s *clusterStore) CreateCluster(ctx context.Context, cls *ttnpb.Cluster) (*ttnpb.Cluster, error) {
	defer trace.StartRegion(ctx, "create cluster").End()
	clsModel := Cluster{
		ClusterID: cls.ClusterID, // The ID is not mutated by fromPB.
	}
	clsModel.fromPB(cls, nil)
	if err := s.createEntity(ctx, &clsModel); err != nil {
		return nil, err
	}
	var clsProto ttnpb.Cluster
	clsModel.toPB(&clsProto, nil)
	return &clsProto, nil
}

func (s *clusterStore) FindClusters(ctx context.Context, ids []*ttnpb.ClusterIdentifiers, fieldMask *types.FieldMask) ([]*ttnpb.Cluster, error) {
	defer trace.StartRegion(ctx, "find clusters").End()
	idStrings := make([]string, len(ids))
	for i, id := range ids {
		idStrings[i] = id.GetClusterID()
	}
	query := s.query(ctx, Cluster{}, withClusterID(idStrings...))
	query = selectClusterFields(ctx, query, fieldMask)
	if limit, offset := limitAndOffsetFromContext(ctx); limit != 0 {
		countTotal(ctx, query.Model(Cluster{}))
		query = query.Limit(limit).Offset(offset)
	}
	var clsModels []Cluster
	query = query.Find(&clsModels)
	setTotal(ctx, uint64(len(clsModels)))
	if query.Error != nil {
		return nil, query.Error
	}
	clsProtos := make([]*ttnpb.Cluster, len(clsModels))
	for i, clsModel := range clsModels {
		clsProto := &ttnpb.Cluster{}
		clsModel.toPB(clsProto, fieldMask)
		clsProtos[i] = clsProto
	}
	return clsProtos, nil
}

func (s *clusterStore) GetCluster(ctx context.Context, id *ttnpb.ClusterIdentifiers, fieldMask *types.FieldMask) (*ttnpb.Cluster, error) {
	defer trace.StartRegion(ctx, "get cluster").End()
	query := s.query(ctx, Cluster{}, withClusterID(id.GetClusterID()))
	query = selectClusterFields(ctx, query, fieldMask)
	var clsModel Cluster
	if err := query.First(&clsModel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errNotFoundForID(id)
		}
		return nil, err
	}
	clsProto := &ttnpb.Cluster{}
	clsModel.toPB(clsProto, fieldMask)
	return clsProto, nil
}

func (s *clusterStore) UpdateCluster(ctx context.Context, cls *ttnpb.Cluster, fieldMask *types.FieldMask) (updated *ttnpb.Cluster, err error) {
	defer trace.StartRegion(ctx, "update cluster").End()
	query := s.query(ctx, Cluster{}, withClusterID(cls.GetClusterID()))
	query = selectClusterFields(ctx, query, fieldMask)
	var clsModel Cluster
	if err = query.First(&clsModel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errNotFoundForID(cls.ClusterIdentifiers)
		}
		return nil, err
	}
	if err := ctx.Err(); err != nil { // Early exit if context canceled
		return nil, err
	}
	oldAttributes := clsModel.Attributes
	columns := clsModel.fromPB(cls, fieldMask)
	if err = s.updateEntity(ctx, &clsModel, columns...); err != nil {
		return nil, err
	}
	if !reflect.DeepEqual(oldAttributes, clsModel.Attributes) {
		if err = s.replaceAttributes(ctx, "cluster", clsModel.ID, oldAttributes, clsModel.Attributes); err != nil {
			return nil, err
		}
	}
	updated = &ttnpb.Cluster{}
	clsModel.toPB(updated, fieldMask)
	return updated, nil
}

func (s *clusterStore) DeleteCluster(ctx context.Context, id *ttnpb.ClusterIdentifiers) error {
	defer trace.StartRegion(ctx, "delete cluster").End()
	return s.deleteEntity(ctx, id)
}
