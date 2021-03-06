/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

// 200 ok object
type GetWarsWarIdKillmails200Ok struct {

	// A hash of this killmail
	KillmailHash string `json:"killmail_hash"`

	// ID of this killmail
	KillmailId int32 `json:"killmail_id"`
}
