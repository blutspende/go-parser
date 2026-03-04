package hl7v26

// RGS - Resource Group
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/RGS
type RGS struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SegmentActionCode string `hl7:"POS=3"`
	ResourceGroupID CWE `hl7:"POS=4"`
}

