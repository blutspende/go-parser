package hl7v28

// ORA_R33 - Observation Report Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ORA_R33
type ORA_R33 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
}

