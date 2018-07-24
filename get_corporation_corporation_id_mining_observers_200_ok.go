/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetCorporationCorporationIdMiningObservers200Ok struct {

	// last_updated string
	LastUpdated string `json:"last_updated"`

	// The entity that was observing the asteroid field when it was mined. 
	ObserverId int64 `json:"observer_id"`

	// The category of the observing entity
	ObserverType string `json:"observer_type"`
}
