/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetCharactersCharacterIdFittings200Ok struct {

	// description string
	Description string `json:"description"`

	// fitting_id integer
	FittingId int32 `json:"fitting_id"`

	// items array
	Items []GetCharactersCharacterIdFittingsItem `json:"items"`

	// name string
	Name string `json:"name"`

	// ship_type_id integer
	ShipTypeId int32 `json:"ship_type_id"`
}
