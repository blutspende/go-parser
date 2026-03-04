package hl7v271

// RSP_K21_QUERY_RESPONSE - Group struct
type RSP_K21_QUERY_RESPONSE struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	QueryResponseInstance QRI `hl7:"TAG=QRI"`
}

// RSP_K21 - Get person demographics response
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/RSP_K21
type RSP_K21 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Query_response RSP_K21_QUERY_RESPONSE `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

