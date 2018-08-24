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
type GetMarketsGroupsMarketGroupIdOk struct {

	// description string
	Description string `json:"description"`

	// market_group_id integer
	MarketGroupId int32 `json:"market_group_id"`

	// name string
	Name string `json:"name"`

	// parent_group_id integer
	ParentGroupId int32 `json:"parent_group_id,omitempty"`

	// types array
	Types []int32 `json:"types"`
}
