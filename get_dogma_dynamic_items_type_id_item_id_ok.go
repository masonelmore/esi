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
type GetDogmaDynamicItemsTypeIdItemIdOk struct {

	// The ID of the character who created the item
	CreatedBy int32 `json:"created_by"`

	// dogma_attributes array
	DogmaAttributes []GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute `json:"dogma_attributes"`

	// dogma_effects array
	DogmaEffects []GetDogmaDynamicItemsTypeIdItemIdDogmaEffect `json:"dogma_effects"`

	// The type ID of the mutator used to generate the dynamic item.
	MutatorTypeId int32 `json:"mutator_type_id"`

	// The type ID of the source item the mutator was applied to create the dynamic item.
	SourceTypeId int32 `json:"source_type_id"`
}
