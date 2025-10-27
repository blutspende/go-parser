package hl7v24

import "time"

// HL7 v2.4 - EQU - Equipment Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/EQU
type EQU struct {
	EquipmentIDentifier EI        `hl7:"POS=2" json:"EquipmentIDentifier,omitempty"`
	EventDateTime       time.Time `hl7:"POS=3" json:"EventDateTime,omitempty"`
	EquipmentState      CE        `hl7:"POS=4" json:"EquipmentState,omitempty"`
	LocalControlState   CE        `hl7:"POS=5" json:"LocalControlState,omitempty"`
	AlertLevel          CE        `hl7:"POS=6" json:"AlertLevel,omitempty"`
}
