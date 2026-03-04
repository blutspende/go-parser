package hl7v28

// IAR - Allergy Reaction
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/IAR
type IAR struct {
	AllergyReactionCode CWE `hl7:"POS=2"`
	AllergySeverityCode CWE `hl7:"POS=3"`
	SensitivityToCausativeAgentCode CWE `hl7:"POS=4"`
	Management string `hl7:"POS=5"`
}

