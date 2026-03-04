package hl7v25

// PMU_B07_CERTIFICATE - Group struct
type PMU_B07_CERTIFICATE struct {
	CertificateDetail CER `hl7:"TAG=CER"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// PMU_B07 - Add personnel record
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/PMU_B07
type PMU_B07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	StaffIdentification STF `hl7:"TAG=STF"`
	PractitionerDetail PRA `hl7:"TAG=PRA;ATR=optional"`
	Certificate []PMU_B07_CERTIFICATE `hl7:"GROUP;ATR=optional"`
}

