package hl7v24

// PMU_B03 - Delete personnel record
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/PMU_B03
type PMU_B03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	StaffIdentification STF `hl7:"TAG=STF"`
}

