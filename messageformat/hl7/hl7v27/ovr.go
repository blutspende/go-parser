package hl7v27

// OVR - Override Segment
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/OVR
type OVR struct {
	BusinessRuleOverrideType CWE `hl7:"POS=2"`
	BusinessRuleOverrideCode CWE `hl7:"POS=3"`
	OverrideComments string `hl7:"POS=4"`
	OverrideEnteredBy XCN `hl7:"POS=5"`
	OverrideAuthorizedBy XCN `hl7:"POS=6"`
}

