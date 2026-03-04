package hl7v251

import "time"

// IIM - Inventory Item Master
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/IIM
type IIM struct {
	PrimaryKeyValueIim CWE `hl7:"POS=2"`
	ServiceItemCode CWE `hl7:"POS=3"`
	InventoryLotNumber string `hl7:"POS=4"`
	InventoryExpirationDate time.Time `hl7:"POS=5"`
	InventoryManufacturerName CWE `hl7:"POS=6"`
	InventoryLocation CWE `hl7:"POS=7"`
	InventoryReceivedDate time.Time `hl7:"POS=8"`
	InventoryReceivedQuantity *int `hl7:"POS=9"`
	InventoryReceivedQuantityUnit CWE `hl7:"POS=10"`
	InventoryReceivedItemCost MO `hl7:"POS=11"`
	InventoryOnHandDate time.Time `hl7:"POS=12"`
	InventoryOnHandQuantity *int `hl7:"POS=13"`
	InventoryOnHandQuantityUnit CWE `hl7:"POS=14"`
	ProcedureCode CE `hl7:"POS=15"`
	ProcedureCodeModifier []CE `hl7:"POS=16"`
}

