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
type GetAlliancesAllianceIdContacts200Ok struct {

	// contact_id integer
	ContactId int32 `json:"contact_id"`

	// contact_type string
	ContactType string `json:"contact_type"`

	// label_ids array
	LabelIds []int64 `json:"label_ids,omitempty"`

	// Standing of the contact
	Standing float32 `json:"standing"`
}
