package hl7v26

// ADD - Addendum
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/ADD
type ADD struct {
	AddendumContinuationPointer string `hl7:"POS=2"`
}

