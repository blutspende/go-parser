package hl7v28

// MFN_M02_MF_STAFF - Group struct
type MFN_M02_MF_STAFF struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	StaffIdentification STF `hl7:"TAG=STF"`
	PractitionerDetail []PRA `hl7:"TAG=PRA;ATR=optional"`
	PractitionerOrganizationUnit []ORG `hl7:"TAG=ORG;ATR=optional"`
	ProfessionalAffiliation []AFF `hl7:"TAG=AFF;ATR=optional"`
	LanguageDetail []LAN `hl7:"TAG=LAN;ATR=optional"`
	EducationalDetail []EDU `hl7:"TAG=EDU;ATR=optional"`
	CertificateDetail []CER `hl7:"TAG=CER;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// MFN_M02 - Master file - staff practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/MFN_M02
type MFN_M02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_staff []MFN_M02_MF_STAFF `hl7:"GROUP"`
}

