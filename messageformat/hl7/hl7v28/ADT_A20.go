package hl7v28

// ADT_A20 - Bed status update
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ADT_A20
type ADT_A20 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	BedStatusUpdate NPU `hl7:"TAG=NPU"`
}

