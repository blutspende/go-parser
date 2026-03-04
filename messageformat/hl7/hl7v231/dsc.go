package hl7v231

// DSC - Continuation pointer segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/DSC
type DSC struct {
	ContinuationPointer string `hl7:"POS=2"`
}

