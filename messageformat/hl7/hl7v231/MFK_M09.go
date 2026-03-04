package hl7v231

// MFK_M09 - Test/observation (categorical) master file acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFK_M09
type MFK_M09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgmentSegment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

