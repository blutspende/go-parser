package hl7v23

// ADD - Addendum
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ADD
type ADD struct {
	AddendumContinuationPointer string `hl7:"POS=2"`
}

