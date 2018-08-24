/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// Coordinates of the victim in Cartesian space relative to the Sun 
type GetKillmailsKillmailIdKillmailHashPosition struct {

	// x number
	X float64 `json:"x"`

	// y number
	Y float64 `json:"y"`

	// z number
	Z float64 `json:"z"`
}
