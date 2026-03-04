package hl7v28

import "time"

// ACC - Accident
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/ACC
type ACC struct {
	AccidentDateTime time.Time `hl7:"POS=2"`
	AccidentCode CWE `hl7:"POS=3"`
	AccidentLocation string `hl7:"POS=4"`
	AutoAccidentState CWE `hl7:"POS=5"`
	AccidentJobRelatedIndicator string `hl7:"POS=6"`
	AccidentDeathIndicator string `hl7:"POS=7"`
	EnteredBy XCN `hl7:"POS=8"`
	AccidentDescription string `hl7:"POS=9"`
	BroughtInBy string `hl7:"POS=10"`
	PoliceNotifiedIndicator string `hl7:"POS=11"`
	AccidentAddress XAD `hl7:"POS=12"`
	DegreeOfPatientLiability *int `hl7:"POS=13"`
	AccidentIdentifier []EI `hl7:"POS=14"`
}

