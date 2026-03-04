package hl7v26

// MFN_M17_MF_DRG - Group struct
type MFN_M17_MF_DRG struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	DrgMasterFileInformation DMI `hl7:"TAG=DMI"`
}

// MFN_M17 - DRG Master File
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/MFN_M17
type MFN_M17 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfDrg []MFN_M17_MF_DRG `hl7:"GROUP"`
}

