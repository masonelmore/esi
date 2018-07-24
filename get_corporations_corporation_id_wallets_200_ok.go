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
type GetCorporationsCorporationIdWallets200Ok struct {

	// balance number
	Balance float64 `json:"balance"`

	// division integer
	Division int32 `json:"division"`
}
