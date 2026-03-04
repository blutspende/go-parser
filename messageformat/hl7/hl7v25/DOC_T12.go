package hl7v25

// DOC_T12_RESULT - Group struct
type DOC_T12_RESULT struct {
	EventType EVN `hl7:"TAG=EVN;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	TranscriptionDocumentHeader TXA `hl7:"TAG=TXA"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// DOC_T12 - Document response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/DOC_T12
type DOC_T12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	Result []DOC_T12_RESULT `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

