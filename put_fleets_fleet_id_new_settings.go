/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// new_settings object
type PutFleetsFleetIdNewSettings struct {

	// Should free-move be enabled in the fleet
	IsFreeMove bool `json:"is_free_move,omitempty"`

	// New fleet MOTD in CCP flavoured HTML
	Motd string `json:"motd,omitempty"`
}
