/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"time"
)

// 200 ok object
type GetKillmailsKillmailIdKillmailHashOk struct {

	// attackers array
	Attackers []GetKillmailsKillmailIdKillmailHashAttacker `json:"attackers"`

	// ID of the killmail
	KillmailId int32 `json:"killmail_id"`

	// Time that the victim was killed and the killmail generated 
	KillmailTime time.Time `json:"killmail_time"`

	// Moon if the kill took place at one
	MoonId int32 `json:"moon_id,omitempty"`

	// Solar system that the kill took place in 
	SolarSystemId int32 `json:"solar_system_id"`

	Victim *GetKillmailsKillmailIdKillmailHashVictim `json:"victim"`

	// War if the killmail is generated in relation to an official war 
	WarId int32 `json:"war_id,omitempty"`
}
