/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetUniverseSystemJumps200Ok struct {

	// ship_jumps integer
	ShipJumps int32 `json:"ship_jumps"`

	// system_id integer
	SystemId int32 `json:"system_id"`
}
