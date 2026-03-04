package hl7v26

// MFN_M03_MF_TEST - Group struct
type MFN_M03_MF_TEST struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
}

// MFN_M03 - Test/Observation Master File Notification
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/MFN_M03
type MFN_M03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfTest []MFN_M03_MF_TEST `hl7:"GROUP"`
}

