/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetCharactersCharacterIdTitles200Ok struct {

	// name string
	Name string `json:"name,omitempty"`

	// title_id integer
	TitleId int32 `json:"title_id,omitempty"`
}
