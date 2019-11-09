/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// last_week object
type GetFwLeaderboardsCharactersLastWeekLastWeek struct {

	// Amount of kills
	Amount int32 `json:"amount,omitempty"`

	// character_id integer
	CharacterId int32 `json:"character_id,omitempty"`
}
