package hl7v231

// QCK_Q02 - Deferred query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/QCK_Q02
type QCK_Q02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
}

