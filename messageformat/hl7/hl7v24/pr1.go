package hl7v24

import "time"

// PR1 - Procedures
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PR1
type PR1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ProcedureCodingMethod string `hl7:"POS=3"`
	ProcedureCode CE `hl7:"POS=4"`
	ProcedureDescription string `hl7:"POS=5"`
	ProcedureDateTime time.Time `hl7:"POS=6"`
	ProcedureFunctionalType string `hl7:"POS=7"`
	ProcedureMinutes *int `hl7:"POS=8"`
	Anesthesiologist []XCN `hl7:"POS=9"`
	AnesthesiaCode string `hl7:"POS=10"`
	AnesthesiaMinutes *int `hl7:"POS=11"`
	Surgeon []XCN `hl7:"POS=12"`
	ProcedurePractitioner []XCN `hl7:"POS=13"`
	ConsentCode CE `hl7:"POS=14"`
	ProcedurePriority string `hl7:"POS=15"`
	AssociatedDiagnosisCode CE `hl7:"POS=16"`
	ProcedureCodeModifier []CE `hl7:"POS=17"`
	ProcedureDrgType string `hl7:"POS=18"`
	TissueTypeCode []CE `hl7:"POS=19"`
}

