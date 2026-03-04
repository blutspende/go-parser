package hl7v27

// CCF_I22 - Collaborative Care Fetch / Collaborative Care Information
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/CCF_I22
type CCF_I22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
}

