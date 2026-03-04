package hl7v28

// ORA_R41 - Observation Report Alert Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ORA_R41
type ORA_R41 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

