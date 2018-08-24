/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// last_week object
type GetFwLeaderboardsCharactersLastWeekLastWeek1 struct {

	// Amount of victory points
	Amount int32 `json:"amount,omitempty"`

	// character_id integer
	CharacterId int32 `json:"character_id,omitempty"`
}
