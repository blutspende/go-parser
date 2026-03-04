package hl7v26

import "time"

// ABS - Abstract
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/ABS
type ABS struct {
	DischargeCareProvider XCN `hl7:"POS=2"`
	TransferMedicalServiceCode CWE `hl7:"POS=3"`
	SeverityOfIllnessCode CWE `hl7:"POS=4"`
	DateTimeOfAttestation time.Time `hl7:"POS=5"`
	AttestedBy XCN `hl7:"POS=6"`
	TriageCode CWE `hl7:"POS=7"`
	AbstractCompletionDateTime time.Time `hl7:"POS=8"`
	AbstractedBy XCN `hl7:"POS=9"`
	CaseCategoryCode CWE `hl7:"POS=10"`
	CaesareanSectionIndicator string `hl7:"POS=11"`
	GestationCategoryCode CWE `hl7:"POS=12"`
	GestationPeriodWeeks *int `hl7:"POS=13"`
	NewbornCode CWE `hl7:"POS=14"`
	StillbornIndicator string `hl7:"POS=15"`
}

