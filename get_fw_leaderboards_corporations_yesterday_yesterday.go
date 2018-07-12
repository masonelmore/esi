/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// yesterday object
type GetFwLeaderboardsCorporationsYesterdayYesterday struct {

	// Amount of kills
	Amount int32 `json:"amount,omitempty"`

	// corporation_id integer
	CorporationId int32 `json:"corporation_id,omitempty"`
}
