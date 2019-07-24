/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// recipient object
type PostCharactersCharacterIdMailRecipient struct {

	// recipient_id integer
	RecipientId int32 `json:"recipient_id"`

	// recipient_type string
	RecipientType string `json:"recipient_type"`
}
