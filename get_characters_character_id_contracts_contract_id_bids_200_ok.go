/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"time"
)

// 200 ok object
type GetCharactersCharacterIdContractsContractIdBids200Ok struct {

	// The amount bid, in ISK
	Amount float32 `json:"amount"`

	// Unique ID for the bid
	BidId int32 `json:"bid_id"`

	// Character ID of the bidder
	BidderId int32 `json:"bidder_id"`

	// Datetime when the bid was placed
	DateBid time.Time `json:"date_bid"`
}
