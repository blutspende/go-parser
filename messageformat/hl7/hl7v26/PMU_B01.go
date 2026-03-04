package hl7v26

// PMU_B01 - Add Personnel Record
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/PMU_B01
type PMU_B01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	StaffIdentification STF `hl7:"TAG=STF"`
	PractitionerDetail []PRA `hl7:"TAG=PRA;ATR=optional"`
	PractitionerOrganizationUnit []ORG `hl7:"TAG=ORG;ATR=optional"`
	ProfessionalAffiliation []AFF `hl7:"TAG=AFF;ATR=optional"`
	LanguageDetail []LAN `hl7:"TAG=LAN;ATR=optional"`
	EducationalDetail []EDU `hl7:"TAG=EDU;ATR=optional"`
	CertificateDetail []CER `hl7:"TAG=CER;ATR=optional"`
}

