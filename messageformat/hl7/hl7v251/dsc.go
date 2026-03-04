package hl7v251

// DSC - Continuation Pointer
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/DSC
type DSC struct {
	ContinuationPointer string `hl7:"POS=2"`
	ContinuationStyle string `hl7:"POS=3"`
}

