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
	"github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

// Cluster model.
type Cluster struct {
	Model
	SoftDelete

	// BEGIN common fields
	ClusterID   string       `gorm:"unique_index:cluster_id_index;type:VARCHAR(36);not null"`
	Name        string       `gorm:"type:VARCHAR"`
	Description string       `gorm:"type:TEXT"`
	Attributes  []Attribute  `gorm:"polymorphic:Entity;polymorphic_value:cluster"`
	Memberships []Membership `gorm:"polymorphic:Entity;polymorphic_value:client"`
	// END common fields

	ClusterSecret string `gorm:"type:VARCHAR"`
}

func init() {
	registerModel(&Cluster{})
}

// functions to set fields from the cluster model into the cluster proto.
var clusterPBSetters = map[string]func(*ttnpb.Cluster, *Cluster){
	nameField:        func(pb *ttnpb.Cluster, cls *Cluster) { pb.Name = cls.Name },
	descriptionField: func(pb *ttnpb.Cluster, cls *Cluster) { pb.Description = cls.Description },
	attributesField:  func(pb *ttnpb.Cluster, cls *Cluster) { pb.Attributes = attributes(cls.Attributes).toMap() },
	secretField:      func(pb *ttnpb.Cluster, cls *Cluster) { pb.Secret = cls.ClusterSecret },
}

// functions to set fields from the cluster proto into the cluster model.
var clusterModelSetters = map[string]func(*Cluster, *ttnpb.Cluster){
	nameField:        func(cls *Cluster, pb *ttnpb.Cluster) { cls.Name = pb.Name },
	descriptionField: func(cls *Cluster, pb *ttnpb.Cluster) { cls.Description = pb.Description },
	attributesField: func(cls *Cluster, pb *ttnpb.Cluster) {
		cls.Attributes = attributes(cls.Attributes).updateFromMap(pb.Attributes)
	},
	secretField: func(cls *Cluster, pb *ttnpb.Cluster) { cls.ClusterSecret = pb.Secret },
}

// fieldMask to use if a nil or empty fieldmask is passed.
var defaultClusterFieldMask = &types.FieldMask{}

func init() {
	paths := make([]string, 0, len(clusterPBSetters))
	for path := range clusterPBSetters {
		paths = append(paths, path)
	}
	defaultClusterFieldMask.Paths = paths
}

// fieldmask path to column name in clusters table.
var clusterColumnNames = map[string][]string{
	attributesField:  {},
	contactInfoField: {},
	nameField:        {nameField},
	descriptionField: {descriptionField},
	secretField:      {"cluster_secret"},
}

func (cls Cluster) toPB(pb *ttnpb.Cluster, fieldMask *types.FieldMask) {
	pb.ClusterIdentifiers.ClusterID = cls.ClusterID
	pb.CreatedAt = cleanTime(cls.CreatedAt)
	pb.UpdatedAt = cleanTime(cls.UpdatedAt)
	if fieldMask == nil || len(fieldMask.Paths) == 0 {
		fieldMask = defaultClusterFieldMask
	}
	for _, path := range fieldMask.Paths {
		if setter, ok := clusterPBSetters[path]; ok {
			setter(pb, &cls)
		}
	}
}

func (cls *Cluster) fromPB(pb *ttnpb.Cluster, fieldMask *types.FieldMask) (columns []string) {
	if fieldMask == nil || len(fieldMask.Paths) == 0 {
		fieldMask = defaultClusterFieldMask
	}
	for _, path := range fieldMask.Paths {
		if setter, ok := clusterModelSetters[path]; ok {
			setter(cls, pb)
			if columnNames, ok := clusterColumnNames[path]; ok {
				columns = append(columns, columnNames...)
			}
			continue
		}
	}
	return
}
