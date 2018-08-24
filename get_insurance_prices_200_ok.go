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
type GetInsurancePrices200Ok struct {

	// A list of a available insurance levels for this ship type
	Levels []GetInsurancePricesLevel `json:"levels"`

	// type_id integer
	TypeId int32 `json:"type_id"`
}
