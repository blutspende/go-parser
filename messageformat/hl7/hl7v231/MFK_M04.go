package hl7v231

// MFK_M04 - Master files charge description acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFK_M04
type MFK_M04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgmentSegment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

