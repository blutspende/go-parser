package hl7v24

import "time"

// AL1 - Patient allergy information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/AL1
type AL1 struct {
	SetID CE `hl7:"POS=2"`
	AllergenTypeCode CE `hl7:"POS=3"`
	AllergenCodeMnemonicDescription CE `hl7:"POS=4"`
	AllergySeverityCode CE `hl7:"POS=5"`
	AllergyReactionCode []string `hl7:"POS=6"`
	IdentificationDate time.Time `hl7:"POS=7;ATR=date"`
}

