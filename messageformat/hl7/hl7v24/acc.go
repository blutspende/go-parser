package hl7v24

import "time"

// ACC - Accident
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ACC
type ACC struct {
	AccidentDateTime time.Time `hl7:"POS=2"`
	AccidentCode CE `hl7:"POS=3"`
	AccidentLocation string `hl7:"POS=4"`
	AutoAccidentState CE `hl7:"POS=5"`
	AccidentJobRelatedIndicator string `hl7:"POS=6"`
	AccidentDeathIndicator string `hl7:"POS=7"`
	EnteredBy XCN `hl7:"POS=8"`
	AccidentDescription string `hl7:"POS=9"`
	BroughtInBy string `hl7:"POS=10"`
	PoliceNotifiedIndicator string `hl7:"POS=11"`
}

