/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// inventory object
type GetCharactersCharacterIdStatsInventory struct {

	// abandon_loot_quantity integer
	AbandonLootQuantity int64 `json:"abandon_loot_quantity,omitempty"`

	// trash_item_quantity integer
	TrashItemQuantity int64 `json:"trash_item_quantity,omitempty"`
}
