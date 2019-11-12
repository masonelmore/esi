/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// item object
type PostCharactersCharacterIdFittingsItem struct {

	// Fitting location for the item. Entries placed in 'Invalid' will be discarded. If this leaves the fitting with nothing, it will cause an error.
	Flag string `json:"flag"`

	// quantity integer
	Quantity int32 `json:"quantity"`

	// type_id integer
	TypeId int32 `json:"type_id"`
}
