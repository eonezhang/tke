/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client
import (
	"time"
)

type HarborLabel struct {
	// The ID of the label
	Id int64 `json:"id,omitempty"`
	// The name the label
	Name string `json:"name,omitempty"`
	// The description the label
	Description string `json:"description,omitempty"`
	// The color the label
	Color string `json:"color,omitempty"`
	// The scope the label
	Scope string `json:"scope,omitempty"`
	// The ID of project that the label belongs to
	ProjectId int64 `json:"project_id,omitempty"`
	// The creation time the label
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The update time of the label
	UpdateTime time.Time `json:"update_time,omitempty"`
}