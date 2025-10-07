package hl7v23

// NTE - Notes and comments segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/NTE
type NTE struct {
	CommentID       int    `hl7:"1" json:"commentID,omitempty"`
	SourceOfcomment string `hl7:"2" json:"sourceOfcomment,omitempty"`
	Comment         string `hl7:"3" json:"comment,omitempty"`
}
