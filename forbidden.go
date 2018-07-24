/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// Forbidden model
type Forbidden struct {

	// Forbidden message
	Error_ string `json:"error"`

	// status code received from SSO
	SsoStatus int32 `json:"sso_status,omitempty"`
}
