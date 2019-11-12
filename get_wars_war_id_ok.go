/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"time"
)

// 200 ok object
type GetWarsWarIdOk struct {

	Aggressor *GetWarsWarIdAggressor `json:"aggressor"`

	// allied corporations or alliances, each object contains either corporation_id or alliance_id
	Allies []GetWarsWarIdAlly `json:"allies,omitempty"`

	// Time that the war was declared
	Declared time.Time `json:"declared"`

	Defender *GetWarsWarIdDefender `json:"defender"`

	// Time the war ended and shooting was no longer allowed
	Finished time.Time `json:"finished,omitempty"`

	// ID of the specified war
	Id int32 `json:"id"`

	// Was the war declared mutual by both parties
	Mutual bool `json:"mutual"`

	// Is the war currently open for allies or not
	OpenForAllies bool `json:"open_for_allies"`

	// Time the war was retracted but both sides could still shoot each other
	Retracted time.Time `json:"retracted,omitempty"`

	// Time when the war started and both sides could shoot each other
	Started time.Time `json:"started,omitempty"`
}
