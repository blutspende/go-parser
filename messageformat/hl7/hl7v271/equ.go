package hl7v271

import "time"

// EQU - Equipment Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/EQU
type EQU struct {
	EquipmentInstanceIdentifier EI `hl7:"POS=2"`
	EventDateTime time.Time `hl7:"POS=3"`
	EquipmentState CWE `hl7:"POS=4"`
	LocalRemoteControlState CWE `hl7:"POS=5"`
	AlertLevel CWE `hl7:"POS=6"`
}

