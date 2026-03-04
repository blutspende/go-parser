package hl7v271

import "time"

// NTE - Notes And Comments
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/NTE
type NTE struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SourceOfComment string `hl7:"POS=3"`
	Comment []string `hl7:"POS=4"`
	CommentType CWE `hl7:"POS=5"`
	EnteredBy XCN `hl7:"POS=6"`
	EnteredDateTime time.Time `hl7:"POS=7"`
	EffectiveStartDate time.Time `hl7:"POS=8"`
	ExpirationDate time.Time `hl7:"POS=9"`
}

