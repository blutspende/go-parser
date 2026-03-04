package hl7v26

import "time"

// PR1 - Procedures
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/PR1
type PR1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ProcedureCode CNE `hl7:"POS=4"`
	ProcedureDateTime time.Time `hl7:"POS=6"`
	ProcedureFunctionalType string `hl7:"POS=7"`
	ProcedureMinutes *int `hl7:"POS=8"`
	AnesthesiaCode string `hl7:"POS=10"`
	AnesthesiaMinutes *int `hl7:"POS=11"`
	ConsentCode CWE `hl7:"POS=14"`
	ProcedurePriority string `hl7:"POS=15"`
	AssociatedDiagnosisCode CWE `hl7:"POS=16"`
	ProcedureCodeModifier []CNE `hl7:"POS=17"`
	ProcedureDrgType string `hl7:"POS=18"`
	TissueTypeCode []CWE `hl7:"POS=19"`
	ProcedureIdentifier EI `hl7:"POS=20"`
	ProcedureActionCode string `hl7:"POS=21"`
	DrgProcedureDeterminationStatus string `hl7:"POS=22"`
	DrgProcedureRelevance string `hl7:"POS=23"`
}

