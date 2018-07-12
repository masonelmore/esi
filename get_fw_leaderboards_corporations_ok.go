/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetFwLeaderboardsCorporationsOk struct {

	Kills *GetFwLeaderboardsCorporationsKills `json:"kills"`

	VictoryPoints *GetFwLeaderboardsCorporationsVictoryPoints `json:"victory_points"`
}
