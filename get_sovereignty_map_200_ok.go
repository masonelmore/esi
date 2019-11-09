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
type GetSovereigntyMap200Ok struct {

	// alliance_id integer
	AllianceId int32 `json:"alliance_id,omitempty"`

	// corporation_id integer
	CorporationId int32 `json:"corporation_id,omitempty"`

	// faction_id integer
	FactionId int32 `json:"faction_id,omitempty"`

	// system_id integer
	SystemId int32 `json:"system_id"`
}
