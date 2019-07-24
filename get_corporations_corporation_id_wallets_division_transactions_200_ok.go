/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"time"
)

// wallet transaction
type GetCorporationsCorporationIdWalletsDivisionTransactions200Ok struct {

	// client_id integer
	ClientId int32 `json:"client_id"`

	// Date and time of transaction
	Date time.Time `json:"date"`

	// is_buy boolean
	IsBuy bool `json:"is_buy"`

	// -1 if there is no corresponding wallet journal entry
	JournalRefId int64 `json:"journal_ref_id"`

	// location_id integer
	LocationId int64 `json:"location_id"`

	// quantity integer
	Quantity int32 `json:"quantity"`

	// Unique transaction ID
	TransactionId int64 `json:"transaction_id"`

	// type_id integer
	TypeId int32 `json:"type_id"`

	// Amount paid per unit
	UnitPrice float64 `json:"unit_price"`
}
