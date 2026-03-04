package hl7v28

// DMI - Drg Master File Information
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/DMI
type DMI struct {
	DiagnosticRelatedGroup CNE `hl7:"POS=2"`
	MajorDiagnosticCategory CNE `hl7:"POS=3"`
	LowerAndUpperTrimPoints NR `hl7:"POS=4"`
	AverageLengthOfStay *int `hl7:"POS=5"`
	RelativeWeight *int `hl7:"POS=6"`
}

