package hl7v231

import "time"

// AL1 - Patient allergy information segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/AL1
type AL1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	AllergyType string `hl7:"POS=3"`
	AllergyCodeMnemonicDescription CE `hl7:"POS=4"`
	AllergySeverity string `hl7:"POS=5"`
	AllergyReaction []string `hl7:"POS=6"`
	IdentificationDate time.Time `hl7:"POS=7;ATR=date"`
}

