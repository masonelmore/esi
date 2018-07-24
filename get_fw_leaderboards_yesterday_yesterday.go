/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// yesterday object
type GetFwLeaderboardsYesterdayYesterday struct {

	// Amount of kills
	Amount int32 `json:"amount,omitempty"`

	// faction_id integer
	FactionId int32 `json:"faction_id,omitempty"`
}
