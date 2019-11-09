/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"time"
)

// 200 ok object
type GetFleetsFleetIdMembers200Ok struct {

	// character_id integer
	CharacterId int32 `json:"character_id"`

	// join_time string
	JoinTime time.Time `json:"join_time"`

	// Member’s role in fleet
	Role string `json:"role"`

	// Localized role names
	RoleName string `json:"role_name"`

	// ship_type_id integer
	ShipTypeId int32 `json:"ship_type_id"`

	// Solar system the member is located in
	SolarSystemId int32 `json:"solar_system_id"`

	// ID of the squad the member is in. If not applicable, will be set to -1
	SquadId int64 `json:"squad_id"`

	// Station in which the member is docked in, if applicable
	StationId int64 `json:"station_id,omitempty"`

	// Whether the member take fleet warps
	TakesFleetWarp bool `json:"takes_fleet_warp"`

	// ID of the wing the member is in. If not applicable, will be set to -1
	WingId int64 `json:"wing_id"`
}
