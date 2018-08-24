/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetCharactersCharacterIdLoyaltyPoints200Ok struct {

	// corporation_id integer
	CorporationId int32 `json:"corporation_id"`

	// loyalty_points integer
	LoyaltyPoints int32 `json:"loyalty_points"`
}
