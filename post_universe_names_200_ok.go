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
type PostUniverseNames200Ok struct {

	// category string
	Category string `json:"category"`

	// id integer
	Id int32 `json:"id"`

	// name string
	Name string `json:"name"`
}
