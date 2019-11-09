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
type GetUniverseStarsStarIdOk struct {

	// Age of star in years
	Age int64 `json:"age"`

	// luminosity number
	Luminosity float32 `json:"luminosity"`

	// name string
	Name string `json:"name"`

	// radius integer
	Radius int64 `json:"radius"`

	// solar_system_id integer
	SolarSystemId int32 `json:"solar_system_id"`

	// spectral_class string
	SpectralClass string `json:"spectral_class"`

	// temperature integer
	Temperature int32 `json:"temperature"`

	// type_id integer
	TypeId int32 `json:"type_id"`
}
