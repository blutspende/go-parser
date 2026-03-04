package hl7v231

import "time"

// AIG - Appointment information - general resource segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/AIG
type AIG struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	ResourceID CE `hl7:"POS=4"`
	ResourceType CE `hl7:"POS=5"`
	ResourceGroup []CE `hl7:"POS=6"`
	ResourceQuantity *int `hl7:"POS=7"`
	ResourceQuantityUnits CE `hl7:"POS=8"`
	StartDateTime time.Time `hl7:"POS=9"`
	StartDateTimeOffset *int `hl7:"POS=10"`
	StartDateTimeOffsetUnits CE `hl7:"POS=11"`
	Duration *int `hl7:"POS=12"`
	DurationUnits CE `hl7:"POS=13"`
	AllowSubstitutionCode string `hl7:"POS=14"`
	FillerStatusCode CE `hl7:"POS=15"`
}

