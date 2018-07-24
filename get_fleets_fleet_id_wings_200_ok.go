/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetFleetsFleetIdWings200Ok struct {

	// id integer
	Id int64 `json:"id"`

	// name string
	Name string `json:"name"`

	// squads array
	Squads []GetFleetsFleetIdWingsSquad `json:"squads"`
}
