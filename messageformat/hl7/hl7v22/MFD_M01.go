package hl7v22

// MFD_M01 - Master files change not otherwise specified (for backwards comp.)
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/MFD_M01
type MFD_M01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgement []MFA `hl7:"TAG=MFA;ATR=optional"`
}

