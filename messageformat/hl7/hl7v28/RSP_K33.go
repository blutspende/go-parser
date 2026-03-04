package hl7v28

// RSP_K33_DONOR - Group struct
type RSP_K33_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
}

// RSP_K33 - Get Donor Record Candidates Response
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RSP_K33
type RSP_K33 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Donor []RSP_K33_DONOR `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

