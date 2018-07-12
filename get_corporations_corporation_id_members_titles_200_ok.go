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
type GetCorporationsCorporationIdMembersTitles200Ok struct {

	// character_id integer
	CharacterId int32 `json:"character_id"`

	// A list of title_id
	Titles []int32 `json:"titles"`
}
