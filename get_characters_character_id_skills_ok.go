/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetCharactersCharacterIdSkillsOk struct {

	// skills array
	Skills []GetCharactersCharacterIdSkillsSkill `json:"skills"`

	// total_sp integer
	TotalSp int64 `json:"total_sp"`

	// Skill points available to be assigned
	UnallocatedSp int32 `json:"unallocated_sp,omitempty"`
}
