package hl7v27

// MFN_M14_MF_SITE_DEFINED - Group struct
type MFN_M14_MF_SITE_DEFINED struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	AnyHl7Segment Hxx `hl7:"TAG=Hxx"`
}

// MFN_M14 - Master file notification - site defined
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/MFN_M14
type MFN_M14 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_site_defined []MFN_M14_MF_SITE_DEFINED `hl7:"GROUP"`
}

