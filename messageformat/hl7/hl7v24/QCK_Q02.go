package hl7v24

// QCK_Q02 - Deferred query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QCK_Q02
type QCK_Q02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
}

