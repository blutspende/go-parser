package hl7v231

// ADD - Addendum segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/ADD
type ADD struct {
	AddendumContinuationPointer string `hl7:"POS=2"`
}

