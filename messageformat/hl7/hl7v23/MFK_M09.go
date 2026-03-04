package hl7v23

// MFK_M09 - Test/observation (categorical) master file acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFK_M09
type MFK_M09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgementSegment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

