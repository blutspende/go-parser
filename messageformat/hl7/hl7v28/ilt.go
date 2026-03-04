package hl7v28

import "time"

// ILT - Material Lot
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/ILT
type ILT struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	InventoryLotNumber string `hl7:"POS=3"`
	InventoryExpirationDate time.Time `hl7:"POS=4"`
	InventoryReceivedDate time.Time `hl7:"POS=5"`
	InventoryReceivedQuantity *int `hl7:"POS=6"`
	InventoryReceivedQuantityUnit CWE `hl7:"POS=7"`
	InventoryReceivedItemCost MO `hl7:"POS=8"`
	InventoryOnHandDate time.Time `hl7:"POS=9"`
	InventoryOnHandQuantity *int `hl7:"POS=10"`
	InventoryOnHandQuantityUnit CWE `hl7:"POS=11"`
}

