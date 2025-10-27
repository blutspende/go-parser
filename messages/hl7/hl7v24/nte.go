package hl7v24

// NTE - Notes and comments segment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NTE
type NTE struct {
	CommentID       int      `hl7:"POS=2" json:"CommentID,omitempty"`
	SourceOfComment string   `hl7:"POS=3" json:"SourceOfComment,omitempty"`
	Comment         []string `hl7:"POS=4" json:"Comment,omitempty"`
	CommentType     string   `hl7:"POS=5" json:"CommentType,omitempty"`
}
