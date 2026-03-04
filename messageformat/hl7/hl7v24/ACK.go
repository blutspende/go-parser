package hl7v24

// ACK - General acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ACK
type ACK struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
}

