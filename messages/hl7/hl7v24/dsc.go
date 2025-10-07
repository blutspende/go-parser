package hl7v24

// HL7 v2.4 - DSC - Continuation Pointer
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/DSC
type DSC struct {
	ContinuationPointer string `hl7:"1" json:"ContinuationPointer,omitempty"`
	ContinuationStyle   string `hl7:"2" json:"ContinuationStyle,omitempty"`
}
