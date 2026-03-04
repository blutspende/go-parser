package hl7v231

// DOC_T12_RESULT - Group struct
type DOC_T12_RESULT struct {
	EventTypeSegment EVN `hl7:"TAG=EVN;ATR=optional"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	DocumentNotificationSegment TXA `hl7:"TAG=TXA"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// DOC_T12 - Document response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/DOC_T12
type DOC_T12 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	Result []DOC_T12_RESULT `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

