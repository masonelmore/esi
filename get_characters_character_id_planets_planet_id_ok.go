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
type GetCharactersCharacterIdPlanetsPlanetIdOk struct {

	// links array
	Links []GetCharactersCharacterIdPlanetsPlanetIdLink `json:"links"`

	// pins array
	Pins []GetCharactersCharacterIdPlanetsPlanetIdPin `json:"pins"`

	// routes array
	Routes []GetCharactersCharacterIdPlanetsPlanetIdRoute `json:"routes"`
}
