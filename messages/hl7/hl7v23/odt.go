package hl7v23

// ODT - Diet tray instructions segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ODT
type ODT struct {
	TrayType        CE     `hl7:"1" json:"trayType,omitempty"`
	ServicePeriod   []CE   `hl7:"2" json:"servicePeriod,omitempty"`
	TextInstruction string `hl7:"3" json:"textInstruction,omitempty"`
}
