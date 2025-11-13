package hl7v23

// ODT - Diet tray instructions segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ODT
type ODT struct {
	TrayType        CE     `hl7:"POS=2" json:"trayType,omitempty"`
	ServicePeriod   []CE   `hl7:"POS=3" json:"servicePeriod,omitempty"`
	TextInstruction string `hl7:"POS=4" json:"textInstruction,omitempty"`
}
