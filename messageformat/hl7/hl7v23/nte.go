package hl7v23

// NTE - Notes and comments segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/NTE
type NTE struct {
	SetID           int    `hl7:"POS=2;ATR=sequence" json:"commentID,omitempty"`
	SourceOfComment string `hl7:"POS=3" json:"sourceOfComment,omitempty"`
	Comment         string `hl7:"POS=4" json:"comment,omitempty"`
}
