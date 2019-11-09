/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetIncursions200Ok struct {

	// The constellation id in which this incursion takes place
	ConstellationId int32 `json:"constellation_id"`

	// The attacking faction's id
	FactionId int32 `json:"faction_id"`

	// Whether the final encounter has boss or not
	HasBoss bool `json:"has_boss"`

	// A list of infested solar system ids that are a part of this incursion
	InfestedSolarSystems []int32 `json:"infested_solar_systems"`

	// Influence of this incursion as a float from 0 to 1
	Influence float32 `json:"influence"`

	// Staging solar system for this incursion
	StagingSolarSystemId int32 `json:"staging_solar_system_id"`

	// The state of this incursion
	State string `json:"state"`

	// The type of this incursion
	Type_ string `json:"type"`
}
