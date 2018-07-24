/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetUniverseStructuresStructureIdOk struct {

	// The full name of the structure
	Name string `json:"name"`

	// The ID of the corporation who owns this particular structure
	OwnerId int32 `json:"owner_id"`

	Position *GetUniverseStructuresStructureIdPosition `json:"position,omitempty"`

	// solar_system_id integer
	SolarSystemId int32 `json:"solar_system_id"`

	// type_id integer
	TypeId int32 `json:"type_id,omitempty"`
}
