package hl7v231

// ODS - Dietary orders, supplements, and preferences segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/ODS
type ODS struct {
	Type string `hl7:"POS=2"`
	ServicePeriod []CE `hl7:"POS=3"`
	DietSupplementOrPreferenceCode []CE `hl7:"POS=4"`
	TextInstruction []string `hl7:"POS=5"`
}

