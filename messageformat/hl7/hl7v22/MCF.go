package hl7v22

// MCF - Delayed acknowledgement 
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/MCF
type MCF struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileAcknowledgement MFA `hl7:"TAG=MFA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
}

