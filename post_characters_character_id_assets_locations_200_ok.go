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
type PostCharactersCharacterIdAssetsLocations200Ok struct {

	// item_id integer
	ItemId int64 `json:"item_id"`

	Position *PostCharactersCharacterIdAssetsLocationsPosition `json:"position"`
}
