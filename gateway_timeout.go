/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// Gateway timeout model
type GatewayTimeout struct {

	// Gateway timeout message
	Error_ string `json:"error"`

	// number of seconds the request was given
	Timeout int32 `json:"timeout,omitempty"`
}
