/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"time"
)

// 200 ok object
type GetCharactersCharacterIdMedals200Ok struct {

	// corporation_id integer
	CorporationId int32 `json:"corporation_id"`

	// date string
	Date time.Time `json:"date"`

	// description string
	Description string `json:"description"`

	// graphics array
	Graphics []GetCharactersCharacterIdMedalsGraphic `json:"graphics"`

	// issuer_id integer
	IssuerId int32 `json:"issuer_id"`

	// medal_id integer
	MedalId int32 `json:"medal_id"`

	// reason string
	Reason string `json:"reason"`

	// status string
	Status string `json:"status"`

	// title string
	Title string `json:"title"`
}
