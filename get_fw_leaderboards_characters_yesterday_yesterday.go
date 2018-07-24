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
type GetFwLeaderboardsCharactersYesterdayYesterday struct {

	// Amount of kills
	Amount int32 `json:"amount,omitempty"`

	// character_id integer
	CharacterId int32 `json:"character_id,omitempty"`
}
