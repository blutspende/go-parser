package hl7v231

import "time"

// ACC - Accident segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/ACC
type ACC struct {
	AccidentDateTime time.Time `hl7:"POS=2"`
	AccidentCode CE `hl7:"POS=3"`
	AccidentLocation string `hl7:"POS=4"`
	AutoAccidentState CE `hl7:"POS=5"`
	AccidentJobRelatedIndicator string `hl7:"POS=6"`
	AccidentDeathIndicator string `hl7:"POS=7"`
}

