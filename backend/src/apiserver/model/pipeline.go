// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "fmt"

// PipelineStatus a label for the status of the Pipeline.
// This is intend to make pipeline creation and deletion atomic.
type PipelineStatus string

const (
	PipelineCreating PipelineStatus = "CREATING"
	PipelineReady    PipelineStatus = "READY"
	PipelineDeleting PipelineStatus = "DELETING"
)

type Pipeline struct {
	UUID           string `gorm:"column:UUID; not null; primary_key"`
	CreatedAtInSec int64  `gorm:"column:CreatedAtInSec; not null"`
	Name           string `gorm:"column:Name; not null; unique"`
	Description    string `gorm:"column:Description; not null"`
	/* Set size to 65535 so it will be stored as longtext. https://dev.mysql.com/doc/refman/8.0/en/column-count-limit.html */
	Parameters string         `gorm:"column:Parameters; not null; size:65535"`
	Status     PipelineStatus `gorm:"column:Status; not null"`
}

func (p Pipeline) GetValueOfPrimaryKey() string {
	return fmt.Sprint(p.UUID)
}

func GetPipelineTablePrimaryKeyColumn() string {
	return "UUID"
}

// PrimaryKeyColumnName returns the primary key for model Pipeline.
func (p *Pipeline) PrimaryKeyColumnName() string {
	return "UUID"
}

// DefaultSortField returns the default sorting field for model Pipeline.
func (p *Pipeline) DefaultSortField() string {
	return "CreatedAtInSec"
}

var pipelineAPIToModelFieldMap = map[string]string{
	"id":          "UUID",
	"name":        "Name",
	"created_at":  "CreatedAtInSec",
	"description": "Description",
}

// APIToModelFieldMap returns a map from API names to field names for model
// Pipeline.
func (p *Pipeline) APIToModelFieldMap() map[string]string {
	return pipelineAPIToModelFieldMap
}
