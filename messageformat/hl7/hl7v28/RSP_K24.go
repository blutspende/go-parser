package hl7v28

// RSP_K24_QUERY_RESPONSE - Group struct
type RSP_K24_QUERY_RESPONSE struct {
	PatientIdentification PID `hl7:"TAG=PID"`
}

// RSP_K24 - Allocate identifiers response
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RSP_K24
type RSP_K24 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Query_response RSP_K24_QUERY_RESPONSE `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

