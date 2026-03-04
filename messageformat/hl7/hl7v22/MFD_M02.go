package hl7v22

// MFD_M02 - Staff Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/MFD_M02
type MFD_M02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgement []MFA `hl7:"TAG=MFA;ATR=optional"`
}

