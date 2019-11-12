/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetFwLeaderboardsCharactersOk struct {

	Kills *GetFwLeaderboardsCharactersKills `json:"kills"`

	VictoryPoints *GetFwLeaderboardsCharactersVictoryPoints `json:"victory_points"`
}
