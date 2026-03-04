package hl7v26

// MFN_M14_MF_SITE_DEFINED - Group struct
type MFN_M14_MF_SITE_DEFINED struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
}

// MFN_M14 - Site Defined Master File Notification
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/MFN_M14
type MFN_M14 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfSiteDefined []MFN_M14_MF_SITE_DEFINED `hl7:"GROUP"`
}

