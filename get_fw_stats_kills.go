/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// Summary of kills against an enemy faction for the given faction
type GetFwStatsKills struct {

	// Last week's total number of kills against enemy factions
	LastWeek int32 `json:"last_week"`

	// Total number of kills against enemy factions since faction warfare began
	Total int32 `json:"total"`

	// Yesterday's total number of kills against enemy factions
	Yesterday int32 `json:"yesterday"`
}
