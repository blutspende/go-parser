package hl7v26

// RSP_K21_QUERY_RESPONSE - Group struct
type RSP_K21_QUERY_RESPONSE struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	AccessRestrictions []ARV `hl7:"TAG=ARV;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	QueryResponseInstance QRI `hl7:"TAG=QRI"`
}

// RSP_K21 - Get Person Demographics Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/RSP_K21
type RSP_K21 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	QueryResponse RSP_K21_QUERY_RESPONSE `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

