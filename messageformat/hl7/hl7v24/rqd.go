package hl7v24

import "time"

// RQD - Requisition Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/RQD
type RQD struct {
	RequisitionLineNumber int `hl7:"POS=2;ATR=sequence"`
	ItemCodeInternal CE `hl7:"POS=3"`
	ItemCodeExternal CE `hl7:"POS=4"`
	HospitalItemCode CE `hl7:"POS=5"`
	RequisitionQuantity *int `hl7:"POS=6"`
	RequisitionUnitOfMeasure CE `hl7:"POS=7"`
	DeptCostCenter string `hl7:"POS=8"`
	ItemNaturalAccountCode string `hl7:"POS=9"`
	DeliverToID CE `hl7:"POS=10"`
	DateNeeded time.Time `hl7:"POS=11;ATR=date"`
}

