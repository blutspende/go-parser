package hl7v251

// ODT - Diet Tray Instructions
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/ODT
type ODT struct {
	TrayType CE `hl7:"POS=2"`
	ServicePeriod []CE `hl7:"POS=3"`
	TextInstruction string `hl7:"POS=4"`
}

