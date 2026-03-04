package hl7v22

// ACK - General acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ACK
type ACK struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
}

