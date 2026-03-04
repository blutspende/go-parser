package hl7v251

import "time"

// EQU - Equipment Detail
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/EQU
type EQU struct {
	EquipmentInstanceIdentifier EI `hl7:"POS=2"`
	EventDateTime time.Time `hl7:"POS=3"`
	EquipmentState CE `hl7:"POS=4"`
	LocalRemoteControlState CE `hl7:"POS=5"`
	AlertLevel CE `hl7:"POS=6"`
}

