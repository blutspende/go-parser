package hl7v23

// NTE - Notes and comments segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/NTE
type NTE struct {
	CommentID       *int   `hl7:"POS=2" json:"commentID,omitempty"`
	SourceOfcomment string `hl7:"POS=3" json:"sourceOfcomment,omitempty"`
	Comment         string `hl7:"POS=4" json:"comment,omitempty"`
}
