package hl7v22

import "time"

// PR1 - Procedures
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/PR1
type PR1 struct {
	SetID                 int               `hl7:"POS=2;ATR=sequence"`
	ProcedureCodingMethod []string          `hl7:"POS=3"`
	ProcedureCode         []string          `hl7:"POS=4"`
	ProcedureDescription  []string          `hl7:"POS=5"`
	ProcedureDateTime     time.Time         `hl7:"POS=6"`
	ProcedureType         string            `hl7:"POS=7"`
	ProcedureMinutes      *int              `hl7:"POS=8"`
	Anesthesiologist      CN_PERSON         `hl7:"POS=9"`
	AnesthesiaCode        string            `hl7:"POS=10"`
	AnesthesiaMinutes     *int              `hl7:"POS=11"`
	Surgeon               CN_PERSON         `hl7:"POS=12"`
	ProcedurePractitioner []CM_PRACTITIONER `hl7:"POS=13"`
	ConsentCode           string            `hl7:"POS=14"`
	ProcedurePriority     *int              `hl7:"POS=15"`
}
