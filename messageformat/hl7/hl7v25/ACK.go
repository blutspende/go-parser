package hl7v25

// ACK - General acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ACK
type ACK struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
}

