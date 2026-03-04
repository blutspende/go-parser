package hl7v271

import "time"

// AIG - Appointment Information - General Resource
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/AIG
type AIG struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	ResourceID CWE `hl7:"POS=4"`
	ResourceType CWE `hl7:"POS=5"`
	ResourceGroup []CWE `hl7:"POS=6"`
	ResourceQuantity *int `hl7:"POS=7"`
	ResourceQuantityUnits CNE `hl7:"POS=8"`
	StartDateTime time.Time `hl7:"POS=9"`
	StartDateTimeOffset *int `hl7:"POS=10"`
	StartDateTimeOffsetUnits CNE `hl7:"POS=11"`
	Duration *int `hl7:"POS=12"`
	DurationUnits CNE `hl7:"POS=13"`
	AllowSubstitutionCode CWE `hl7:"POS=14"`
	FillerStatusCode CWE `hl7:"POS=15"`
}

