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
type GetCorporationsCorporationIdAssets200Ok struct {

	// is_blueprint_copy boolean
	IsBlueprintCopy bool `json:"is_blueprint_copy,omitempty"`

	// is_singleton boolean
	IsSingleton bool `json:"is_singleton"`

	// item_id integer
	ItemId int64 `json:"item_id"`

	// location_flag string
	LocationFlag string `json:"location_flag"`

	// location_id integer
	LocationId int64 `json:"location_id"`

	// location_type string
	LocationType string `json:"location_type"`

	// quantity integer
	Quantity int32 `json:"quantity"`

	// type_id integer
	TypeId int32 `json:"type_id"`
}
