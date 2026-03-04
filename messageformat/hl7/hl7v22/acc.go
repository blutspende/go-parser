package hl7v22

import "time"

// ACC - Accident
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/ACC
type ACC struct {
	AccidentDateTime time.Time `hl7:"POS=2"`
	AccidentCode string `hl7:"POS=3"`
	AccidentLocation string `hl7:"POS=4"`
}

