/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// active_total object
type GetFwLeaderboardsActiveTotalActiveTotal1 struct {

	// Amount of victory points
	Amount int32 `json:"amount,omitempty"`

	// faction_id integer
	FactionId int32 `json:"faction_id,omitempty"`
}
