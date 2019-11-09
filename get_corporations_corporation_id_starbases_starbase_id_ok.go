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
type GetCorporationsCorporationIdStarbasesStarbaseIdOk struct {

	// allow_alliance_members boolean
	AllowAllianceMembers bool `json:"allow_alliance_members"`

	// allow_corporation_members boolean
	AllowCorporationMembers bool `json:"allow_corporation_members"`

	// Who can anchor starbase (POS) and its structures
	Anchor string `json:"anchor"`

	// attack_if_at_war boolean
	AttackIfAtWar bool `json:"attack_if_at_war"`

	// attack_if_other_security_status_dropping boolean
	AttackIfOtherSecurityStatusDropping bool `json:"attack_if_other_security_status_dropping"`

	// Starbase (POS) will attack if target's security standing is lower than this value
	AttackSecurityStatusThreshold float32 `json:"attack_security_status_threshold,omitempty"`

	// Starbase (POS) will attack if target's standing is lower than this value
	AttackStandingThreshold float32 `json:"attack_standing_threshold,omitempty"`

	// Who can take fuel blocks out of the starbase (POS)'s fuel bay
	FuelBayTake string `json:"fuel_bay_take"`

	// Who can view the starbase (POS)'s fule bay. Characters either need to have required role or belong to the starbase (POS) owner's corporation or alliance, as described by the enum, all other access settings follows the same scheme
	FuelBayView string `json:"fuel_bay_view"`

	// Fuel blocks and other things that will be consumed when operating a starbase (POS)
	Fuels []GetCorporationsCorporationIdStarbasesStarbaseIdFuel `json:"fuels,omitempty"`

	// Who can offline starbase (POS) and its structures
	Offline string `json:"offline"`

	// Who can online starbase (POS) and its structures
	Online string `json:"online"`

	// Who can unanchor starbase (POS) and its structures
	Unanchor string `json:"unanchor"`

	// True if the starbase (POS) is using alliance standings, otherwise using corporation's
	UseAllianceStandings bool `json:"use_alliance_standings"`
}
