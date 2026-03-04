package hl7v23

// QCK_Q02 - Deferred query response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/QCK_Q02
type QCK_Q02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
}

