package hl7v271

// MFK_M06 - Clinical Study with Phases and Schedules Master File Application Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/MFK_M06
type MFK_M06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MasterFileAcknowledgment []MFA `hl7:"TAG=MFA;ATR=optional"`
}

