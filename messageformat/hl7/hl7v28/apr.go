package hl7v28

// APR - Appointment Preferences
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/APR
type APR struct {
	TimeSelectionCriteria []SCV `hl7:"POS=2"`
	ResourceSelectionCriteria []SCV `hl7:"POS=3"`
	LocationSelectionCriteria []SCV `hl7:"POS=4"`
	SlotSpacingCriteria *int `hl7:"POS=5"`
	FillerOverrideCriteria []SCV `hl7:"POS=6"`
}

