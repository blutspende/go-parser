package hl7v25

// ADD - Addendum
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/ADD
type ADD struct {
	AddendumContinuationPointer string `hl7:"POS=2"`
}

