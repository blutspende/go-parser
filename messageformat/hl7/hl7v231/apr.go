package hl7v231

// APR - Appointment preferences segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/APR
type APR struct {
	TimeSelectionCriteria []SCV `hl7:"POS=2"`
	ResourceSelectionCriteria []SCV `hl7:"POS=3"`
	LocationSelectionCriteria []SCV `hl7:"POS=4"`
	SlotSpacingCriteria *int `hl7:"POS=5"`
	FillerOverrideCriteria []SCV `hl7:"POS=6"`
}

