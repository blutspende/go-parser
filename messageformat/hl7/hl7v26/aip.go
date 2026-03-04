package hl7v26

import "time"

// AIP - Appointment Information - Personnel Resource
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/AIP
type AIP struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	PersonnelResourceID []XCN `hl7:"POS=4"`
	ResourceType CWE `hl7:"POS=5"`
	ResourceGroup CWE `hl7:"POS=6"`
	StartDateTime time.Time `hl7:"POS=7"`
	StartDateTimeOffset *int `hl7:"POS=8"`
	StartDateTimeOffsetUnits CNE `hl7:"POS=9"`
	Duration *int `hl7:"POS=10"`
	DurationUnits CNE `hl7:"POS=11"`
	AllowSubstitutionCode string `hl7:"POS=12"`
	FillerStatusCode CWE `hl7:"POS=13"`
}

