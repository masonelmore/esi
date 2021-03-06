/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// destination object
type GetUniverseStargatesStargateIdDestination struct {

	// The stargate this stargate connects to
	StargateId int32 `json:"stargate_id"`

	// The solar system this stargate connects to
	SystemId int32 `json:"system_id"`
}
