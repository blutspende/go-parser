package hl7v24

// RSP_K22_QUERY_RESPONSE - Group struct
type RSP_K22_QUERY_RESPONSE struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	QueryResponseInstance QRI `hl7:"TAG=QRI;ATR=optional"`
}

// RSP_K22 - Response - Find candidates
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RSP_K22
type RSP_K22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	QueryResponse []RSP_K22_QUERY_RESPONSE `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

