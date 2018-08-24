/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"time"
)

// 200 ok object
type GetCharactersCharacterIdOk struct {

	// The character's alliance ID
	AllianceId int32 `json:"alliance_id,omitempty"`

	// ancestry_id integer
	AncestryId int32 `json:"ancestry_id,omitempty"`

	// Creation date of the character
	Birthday time.Time `json:"birthday"`

	// bloodline_id integer
	BloodlineId int32 `json:"bloodline_id"`

	// The character's corporation ID
	CorporationId int32 `json:"corporation_id"`

	// description string
	Description string `json:"description,omitempty"`

	// ID of the faction the character is fighting for, if the character is enlisted in Factional Warfare
	FactionId int32 `json:"faction_id,omitempty"`

	// gender string
	Gender string `json:"gender"`

	// name string
	Name string `json:"name"`

	// race_id integer
	RaceId int32 `json:"race_id"`

	// security_status number
	SecurityStatus float32 `json:"security_status,omitempty"`
}
