package hl7v26

// MFN_M01_MF - Group struct
type MFN_M01_MF struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
}

// MFN_M01 - Master File Notification
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/MFN_M01
type MFN_M01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf []MFN_M01_MF `hl7:"GROUP"`
}

