package hl7v24

// NTE - Notes and Comments
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NTE
type NTE struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SourceOfComment string `hl7:"POS=3"`
	Comment []string `hl7:"POS=4"`
	CommentType CE `hl7:"POS=5"`
}

