package hl7v28

// ODS - Dietary Orders, Supplements, And Preferences
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/ODS
type ODS struct {
	Type string `hl7:"POS=2"`
	ServicePeriod []CWE `hl7:"POS=3"`
	DietSupplementOrPreferenceCode []CWE `hl7:"POS=4"`
	TextInstruction []string `hl7:"POS=5"`
}

