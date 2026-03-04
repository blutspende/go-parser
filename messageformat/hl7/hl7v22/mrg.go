package hl7v22

// MRG - Merge Patient Information
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/MRG
type MRG struct {
	PriorPatientIDInternal CM_PAT_ID `hl7:"POS=2"`
	PriorAlternatePatientID CM_PAT_ID `hl7:"POS=3"`
	PriorPatientAccountNumber CK `hl7:"POS=4"`
	PriorPatientIDExternal CK `hl7:"POS=5"`
}

