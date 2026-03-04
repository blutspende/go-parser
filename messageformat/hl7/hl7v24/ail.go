package hl7v24

import "time"

// AIL - Appointment Information - Location Resource
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/AIL
type AIL struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	LocationResourceID PL `hl7:"POS=4"`
	LocationTypeAil CE `hl7:"POS=5"`
	LocationGroup CE `hl7:"POS=6"`
	StartDateTime time.Time `hl7:"POS=7"`
	StartDateTimeOffset *int `hl7:"POS=8"`
	StartDateTimeOffsetUnits CE `hl7:"POS=9"`
	Duration *int `hl7:"POS=10"`
	DurationUnits CE `hl7:"POS=11"`
	AllowSubstitutionCode string `hl7:"POS=12"`
	FillerStatusCode CE `hl7:"POS=13"`
}

