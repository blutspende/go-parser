package hl7v231

// MFK_M03 - Master file - Test/Observation acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFK_M03
type MFK_M03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgmentSegment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

