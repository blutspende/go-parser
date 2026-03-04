package hl7v251

// APR - Appointment Preferences
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/APR
type APR struct {
	TimeSelectionCriteria []SCV `hl7:"POS=2"`
	ResourceSelectionCriteria []SCV `hl7:"POS=3"`
	LocationSelectionCriteria []SCV `hl7:"POS=4"`
	SlotSpacingCriteria *int `hl7:"POS=5"`
	FillerOverrideCriteria []SCV `hl7:"POS=6"`
}

