package hl7v27

// AL1 - Patient Allergy Information
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/AL1
type AL1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	AllergenTypeCode CWE `hl7:"POS=3"`
	AllergenCodeMnemonicDescription CWE `hl7:"POS=4"`
	AllergySeverityCode CWE `hl7:"POS=5"`
	AllergyReactionCode []string `hl7:"POS=6"`
}

