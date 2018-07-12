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
type GetCharactersCharacterIdMailLists200Ok struct {

	// Mailing list ID
	MailingListId int32 `json:"mailing_list_id"`

	// name string
	Name string `json:"name"`
}
