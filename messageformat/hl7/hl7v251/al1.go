package hl7v251

import "time"

// AL1 - Patient Allergy Information
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/AL1
type AL1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	AllergenTypeCode CE `hl7:"POS=3"`
	AllergenCodeMnemonicDescription CE `hl7:"POS=4"`
	AllergySeverityCode CE `hl7:"POS=5"`
	AllergyReactionCode []string `hl7:"POS=6"`
	IdentificationDate time.Time `hl7:"POS=7;ATR=date"`
}

