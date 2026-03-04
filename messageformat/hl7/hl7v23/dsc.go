package hl7v23

// DSC - Continuation pointer segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/DSC
type DSC struct {
	ContinuationPointer string `hl7:"POS=2"`
}

