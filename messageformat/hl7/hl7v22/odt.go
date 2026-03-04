package hl7v22

// ODT - Diet Tray Instruction
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/ODT
type ODT struct {
	TrayType CE `hl7:"POS=2"`
	ServicePeriod []CE `hl7:"POS=3"`
	TextInstruction []string `hl7:"POS=4"`
}

