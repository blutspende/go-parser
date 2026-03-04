package hl7v26

// ODT - Diet Tray Instructions
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/ODT
type ODT struct {
	TrayType CWE `hl7:"POS=2"`
	ServicePeriod []CWE `hl7:"POS=3"`
	TextInstruction string `hl7:"POS=4"`
}

