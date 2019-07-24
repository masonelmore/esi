/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// Summary of kills done by the given corporation against enemy factions
type GetCorporationsCorporationIdFwStatsKills struct {

	// Last week's total number of kills by members of the given corporation against enemy factions
	LastWeek int32 `json:"last_week"`

	// Total number of kills by members of the given corporation against enemy factions since the corporation enlisted
	Total int32 `json:"total"`

	// Yesterday's total number of kills by members of the given corporation against enemy factions
	Yesterday int32 `json:"yesterday"`
}
