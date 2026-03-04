package hl7v22

// DSR_Q01 - Display Response Message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/DSR_Q01
type DSR_Q01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileAcknowledgement MFA `hl7:"TAG=MFA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

