package hl7v24

// RSP_K23 - Response - Get corresponding identifiers
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RSP_K23
type RSP_K23 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

