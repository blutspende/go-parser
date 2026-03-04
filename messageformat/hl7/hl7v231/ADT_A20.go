package hl7v231

// ADT_A20 - Bed status update
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A20
type ADT_A20 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	BedStatusUpdateSegment NPU `hl7:"TAG=NPU"`
}

