package hl7v251

// QCK_Q02 - Deferred Query Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QCK_Q02
type QCK_Q02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
}

