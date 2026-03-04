package hl7v25

// MFN_M13 - Master File Notification - General
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/MFN_M13
type MFN_M13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileEntry []MFE `hl7:"TAG=MFE"`
}

