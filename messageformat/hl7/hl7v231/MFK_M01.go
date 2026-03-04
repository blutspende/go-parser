package hl7v231

// MFK_M01_MF - Group struct
type MFK_M01_MF struct {
	MasterFileAcknowledgmentSegment MFA `hl7:"TAG=MFA;ATR=optional"`
}

// MFK_M01 - Master file not otherwise specified acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFK_M01
type MFK_M01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	MasterFileIdentificationSegment1 MFI `hl7:"TAG=MFI"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	MasterFileAcknowledgmentSegment []MFA `hl7:"TAG=MFA;ATR=optional"`
	MasterFileIdentificationSegment2 MFI `hl7:"TAG=MFI"`
	Mf []MFK_M01_MF `hl7:"GROUP"`
}

