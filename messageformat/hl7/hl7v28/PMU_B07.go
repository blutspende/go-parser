package hl7v28

// PMU_B07_CERTIFICATE - Group struct
type PMU_B07_CERTIFICATE struct {
	CertificateDetail CER `hl7:"TAG=CER"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// PMU_B07 - Grant Certificate/Permission
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/PMU_B07
type PMU_B07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	StaffIdentification STF `hl7:"TAG=STF"`
	PractitionerDetail PRA `hl7:"TAG=PRA;ATR=optional"`
	Certificate []PMU_B07_CERTIFICATE `hl7:"GROUP;ATR=optional"`
}

