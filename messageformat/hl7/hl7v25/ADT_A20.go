package hl7v25

// ADT_A20 - Bed Status Update
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ADT_A20
type ADT_A20 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	BedStatusUpdate NPU `hl7:"TAG=NPU"`
}

