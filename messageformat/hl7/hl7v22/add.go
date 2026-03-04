package hl7v22

// ADD - Addendum
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/ADD
type ADD struct {
	AddendumContinuationPointer string `hl7:"POS=2"`
}

