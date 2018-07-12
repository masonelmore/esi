/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// victim object
type GetKillmailsKillmailIdKillmailHashVictim struct {

	// alliance_id integer
	AllianceId int32 `json:"alliance_id,omitempty"`

	// character_id integer
	CharacterId int32 `json:"character_id,omitempty"`

	// corporation_id integer
	CorporationId int32 `json:"corporation_id,omitempty"`

	// How much total damage was taken by the victim 
	DamageTaken int32 `json:"damage_taken"`

	// faction_id integer
	FactionId int32 `json:"faction_id,omitempty"`

	// items array
	Items []GetKillmailsKillmailIdKillmailHashItem `json:"items,omitempty"`

	Position *GetKillmailsKillmailIdKillmailHashPosition `json:"position,omitempty"`

	// The ship that the victim was piloting and was destroyed 
	ShipTypeId int32 `json:"ship_type_id"`
}
