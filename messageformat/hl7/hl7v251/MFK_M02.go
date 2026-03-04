package hl7v251

// MFK_M02_MF_STAFF - Group struct
type MFK_M02_MF_STAFF struct {
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

// MFK_M02 - Master file acknowledgment - Staff/practitioner master file
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/MFK_M02
type MFK_M02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
	MfStaff []MFK_M02_MF_STAFF `hl7:"GROUP"`
}

