package hl7v231

// RGS - Resource group segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/RGS
type RGS struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	ResourceGroupID CE `hl7:"POS=4"`
}

