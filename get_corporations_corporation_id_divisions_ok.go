/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetCorporationsCorporationIdDivisionsOk struct {

	// hangar array
	Hangar []GetCorporationsCorporationIdDivisionsHangarHangar `json:"hangar,omitempty"`

	// wallet array
	Wallet []GetCorporationsCorporationIdDivisionsWalletWallet `json:"wallet,omitempty"`
}
