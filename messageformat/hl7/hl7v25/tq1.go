package hl7v25

import "time"

// TQ1 - Timing/Quantity
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/TQ1
type TQ1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	Quantity CQ `hl7:"POS=3"`
	RepeatPattern []RPT `hl7:"POS=4"`
	ExplicitTime []string `hl7:"POS=5"` // TODO - handle TM type parsing
	RelativeTimeAndUnits []CQ `hl7:"POS=6"`
	ServiceDuration CQ `hl7:"POS=7"`
	StartDateTime time.Time `hl7:"POS=8"`
	EndDateTime time.Time `hl7:"POS=9"`
	Priority []CWE `hl7:"POS=10"`
	ConditionText string `hl7:"POS=11"`
	TextInstruction string `hl7:"POS=12"`
	Conjunction string `hl7:"POS=13"`
	OccurrenceDuration CQ `hl7:"POS=14"`
	TotalOccurrences *int `hl7:"POS=15"`
}

