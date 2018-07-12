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
type GetCharactersCharacterIdBookmarksFolders200Ok struct {

	// folder_id integer
	FolderId int32 `json:"folder_id"`

	// name string
	Name string `json:"name"`
}
