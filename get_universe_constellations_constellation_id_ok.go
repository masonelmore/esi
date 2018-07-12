/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetUniverseConstellationsConstellationIdOk struct {

	// constellation_id integer
	ConstellationId int32 `json:"constellation_id"`

	// name string
	Name string `json:"name"`

	Position *GetUniverseConstellationsConstellationIdPosition `json:"position"`

	// The region this constellation is in
	RegionId int32 `json:"region_id"`

	// systems array
	Systems []int32 `json:"systems"`
}
