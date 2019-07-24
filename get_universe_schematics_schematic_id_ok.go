/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetUniverseSchematicsSchematicIdOk struct {

	// Time in seconds to process a run
	CycleTime int32 `json:"cycle_time"`

	// schematic_name string
	SchematicName string `json:"schematic_name"`
}
