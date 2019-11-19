/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// character object
type GetCharactersCharacterIdStatsCharacter struct {

	// days_of_activity integer
	DaysOfActivity int64 `json:"days_of_activity,omitempty"`

	// minutes integer
	Minutes int64 `json:"minutes,omitempty"`

	// sessions_started integer
	SessionsStarted int64 `json:"sessions_started,omitempty"`
}
