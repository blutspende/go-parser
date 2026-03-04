package hl7v22

// DSC - Continuation Pointer
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/DSC
type DSC struct {
	ContinuationPointer string `hl7:"POS=2"`
}

