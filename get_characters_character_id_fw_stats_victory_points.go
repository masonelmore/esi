/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// Summary of victory points gained by the given character for the enlisted faction
type GetCharactersCharacterIdFwStatsVictoryPoints struct {

	// Last week's victory points gained by the given character
	LastWeek int32 `json:"last_week"`

	// Total victory points gained since the given character enlisted
	Total int32 `json:"total"`

	// Yesterday's victory points gained by the given character
	Yesterday int32 `json:"yesterday"`
}
