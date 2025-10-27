package hl7v23

// ODS - Dietary orders, supplements, and preferences
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ODS
type ODS struct {
	Type                           string   `hl7:"POS=2" json:"type,omitempty"`
	ServicePeriod                  []CE     `hl7:"POS=3" json:"servicePeriod,omitempty"`
	DietSupplementOrPreferenceCode []CE     `hl7:"POS=4" json:"dietSupplementOrPreferenceCode,omitempty"`
	TextInstruction                []string `hl7:"POS=5" json:"textInstruction,omitempty"`
}
