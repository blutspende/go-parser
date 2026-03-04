package hl7v24

// RGS - Resource Group
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/RGS
type RGS struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	ResourceGroupID CE `hl7:"POS=4"`
}

