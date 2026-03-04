package hl7v23

// MFD_MFA - Master file delayed application acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFD_MFA
type MFD_MFA struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgementSegment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

