package hl7v28

// ADD - Addendum
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/ADD
type ADD struct {
	AddendumContinuationPointer string `hl7:"POS=2"`
}

