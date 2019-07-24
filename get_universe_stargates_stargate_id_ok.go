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
type GetUniverseStargatesStargateIdOk struct {

	Destination *GetUniverseStargatesStargateIdDestination `json:"destination"`

	// name string
	Name string `json:"name"`

	Position *GetUniverseStargatesStargateIdPosition `json:"position"`

	// stargate_id integer
	StargateId int32 `json:"stargate_id"`

	// The solar system this stargate is in
	SystemId int32 `json:"system_id"`

	// type_id integer
	TypeId int32 `json:"type_id"`
}
