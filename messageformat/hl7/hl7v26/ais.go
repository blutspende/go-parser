package hl7v26

import "time"

// AIS - Appointment Information
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/AIS
type AIS struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	UniversalServiceIdentifier CWE `hl7:"POS=4"`
	StartDateTime time.Time `hl7:"POS=5"`
	StartDateTimeOffset *int `hl7:"POS=6"`
	StartDateTimeOffsetUnits CNE `hl7:"POS=7"`
	Duration *int `hl7:"POS=8"`
	DurationUnits CNE `hl7:"POS=9"`
	AllowSubstitutionCode string `hl7:"POS=10"`
	FillerStatusCode CWE `hl7:"POS=11"`
	PlacerSupplementalServiceInformation []CWE `hl7:"POS=12"`
	FillerSupplementalServiceInformation []CWE `hl7:"POS=13"`
}

