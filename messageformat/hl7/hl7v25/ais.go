package hl7v25

import "time"

// AIS - Appointment Information
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/AIS
type AIS struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	UniversalServiceIdentifier CE `hl7:"POS=4"`
	StartDateTime time.Time `hl7:"POS=5"`
	StartDateTimeOffset *int `hl7:"POS=6"`
	StartDateTimeOffsetUnits CE `hl7:"POS=7"`
	Duration *int `hl7:"POS=8"`
	DurationUnits CE `hl7:"POS=9"`
	AllowSubstitutionCode string `hl7:"POS=10"`
	FillerStatusCode CE `hl7:"POS=11"`
	PlacerSupplementalServiceInformation []CE `hl7:"POS=12"`
	FillerSupplementalServiceInformation []CE `hl7:"POS=13"`
}

