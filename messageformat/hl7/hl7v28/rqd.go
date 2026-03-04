package hl7v28

import "time"

// RQD - Requisition Detail
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/RQD
type RQD struct {
	RequisitionLineNumber int `hl7:"POS=2;ATR=sequence"`
	ItemCodeInternal CWE `hl7:"POS=3"`
	ItemCodeExternal CWE `hl7:"POS=4"`
	HospitalItemCode CWE `hl7:"POS=5"`
	RequisitionQuantity *int `hl7:"POS=6"`
	RequisitionUnitOfMeasure CWE `hl7:"POS=7"`
	CostCenterAccountNumber CX `hl7:"POS=8"`
	ItemNaturalAccountCode CWE `hl7:"POS=9"`
	DeliverToID CWE `hl7:"POS=10"`
	DateNeeded time.Time `hl7:"POS=11;ATR=date"`
}

