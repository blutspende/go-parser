package hl7v231

// ACK - General acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ACK
type ACK struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
}

