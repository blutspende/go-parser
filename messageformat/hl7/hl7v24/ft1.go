package hl7v24

import "time"

// FT1 - Financial Transaction
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/FT1
type FT1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	TransactionID string `hl7:"POS=3"`
	TransactionBatchID string `hl7:"POS=4"`
	TransactionDate time.Time `hl7:"POS=5"`
	TransactionPostingDate time.Time `hl7:"POS=6"`
	TransactionType string `hl7:"POS=7"`
	TransactionCode CE `hl7:"POS=8"`
	TransactionDescription string `hl7:"POS=9"`
	TransactionDescriptionAlt string `hl7:"POS=10"`
	TransactionQuantity *int `hl7:"POS=11"`
	TransactionAmountExtended CP `hl7:"POS=12"`
	TransactionAmountUnit CP `hl7:"POS=13"`
	DepartmentCode CE `hl7:"POS=14"`
	InsurancePlanID CE `hl7:"POS=15"`
	InsuranceAmount CP `hl7:"POS=16"`
	AssignedPatientLocation PL `hl7:"POS=17"`
	FeeSchedule string `hl7:"POS=18"`
	PatientType string `hl7:"POS=19"`
	DiagnosisCodeFt1 []CE `hl7:"POS=20"`
	PerformedByCode []XCN `hl7:"POS=21"`
	OrderedByCode []XCN `hl7:"POS=22"`
	UnitCost CP `hl7:"POS=23"`
	FillerOrderNumber EI `hl7:"POS=24"`
	EnteredByCode []XCN `hl7:"POS=25"`
	ProcedureCode CE `hl7:"POS=26"`
	ProcedureCodeModifier []CE `hl7:"POS=27"`
}

