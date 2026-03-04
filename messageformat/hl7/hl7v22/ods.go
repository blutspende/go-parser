package hl7v22

// ODS - Dietary Orders, Supplements, And Preferences
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/ODS
type ODS struct {
	Type string `hl7:"POS=2"`
	ServicePeriod []CE `hl7:"POS=3"`
	DietSupplementOrPreferenceCode []CE `hl7:"POS=4"`
	TextInstruction []string `hl7:"POS=5"`
}

