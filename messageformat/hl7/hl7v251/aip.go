package hl7v251

import "time"

// AIP - Appointment Information - Personnel Resource
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/AIP
type AIP struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	PersonnelResourceID []XCN `hl7:"POS=4"`
	ResourceType CE `hl7:"POS=5"`
	ResourceGroup CE `hl7:"POS=6"`
	StartDateTime time.Time `hl7:"POS=7"`
	StartDateTimeOffset *int `hl7:"POS=8"`
	StartDateTimeOffsetUnits CE `hl7:"POS=9"`
	Duration *int `hl7:"POS=10"`
	DurationUnits CE `hl7:"POS=11"`
	AllowSubstitutionCode string `hl7:"POS=12"`
	FillerStatusCode CE `hl7:"POS=13"`
}

