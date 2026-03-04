package hl7v23

// DOC_T12_RESULT - Group struct
type DOC_T12_RESULT struct {
	EventType EVN `hl7:"TAG=EVN;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	DocumentNotificationSegment TXA `hl7:"TAG=TXA"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// DOC_T12 - Document response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/DOC_T12
type DOC_T12 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	Result []DOC_T12_RESULT `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

