package hl7v28

// MFN_M13 - Master file notification - general
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/MFN_M13
type MFN_M13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileEntry []MFE `hl7:"TAG=MFE"`
}

