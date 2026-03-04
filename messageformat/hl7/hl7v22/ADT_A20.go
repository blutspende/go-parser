package hl7v22

// ADT_A20 - Bed Status Update
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADT_A20
type ADT_A20 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	BedStatusUpdate NPU `hl7:"TAG=NPU"`
}

