package hl7v25

// PMU_B03 - Delete personnel record
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/PMU_B03
type PMU_B03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	StaffIdentification STF `hl7:"TAG=STF"`
}

