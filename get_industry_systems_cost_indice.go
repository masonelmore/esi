/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// cost_indice object
type GetIndustrySystemsCostIndice struct {

	// activity string
	Activity string `json:"activity"`

	// cost_index number
	CostIndex float32 `json:"cost_index"`
}
