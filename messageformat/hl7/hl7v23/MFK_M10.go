package hl7v23

// MFK_M10 - Test/observation batteries master file acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFK_M10
type MFK_M10 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgementSegment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

