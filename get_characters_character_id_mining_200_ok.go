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
type GetCharactersCharacterIdMining200Ok struct {

	// date string
	Date string `json:"date"`

	// quantity integer
	Quantity int64 `json:"quantity"`

	// solar_system_id integer
	SolarSystemId int32 `json:"solar_system_id"`

	// type_id integer
	TypeId int32 `json:"type_id"`
}
