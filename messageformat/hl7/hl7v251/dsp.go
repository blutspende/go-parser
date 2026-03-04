package hl7v251

// DSP - Display Data
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/DSP
type DSP struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	DisplayLevel int `hl7:"POS=3;ATR=sequence"`
	DataLine string `hl7:"POS=4"`
	LogicalBreakPoint string `hl7:"POS=5"`
	ResultID string `hl7:"POS=6"`
}

