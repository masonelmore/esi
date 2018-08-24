/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetUniverseRegionsRegionIdOk struct {

	// constellations array
	Constellations []int32 `json:"constellations"`

	// description string
	Description string `json:"description,omitempty"`

	// name string
	Name string `json:"name"`

	// region_id integer
	RegionId int32 `json:"region_id"`
}
