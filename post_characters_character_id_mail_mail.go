/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// mail object
type PostCharactersCharacterIdMailMail struct {

	// approved_cost integer
	ApprovedCost int64 `json:"approved_cost,omitempty"`

	// body string
	Body string `json:"body"`

	// recipients array
	Recipients []PostCharactersCharacterIdMailRecipient `json:"recipients"`

	// subject string
	Subject string `json:"subject"`
}
