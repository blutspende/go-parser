package hl7v24

// MDM_T04 - Document status change notification and content
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MDM_T04
type MDM_T04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	TranscriptionDocumentHeader TXA `hl7:"TAG=TXA"`
	ObservationResult []OBX `hl7:"TAG=OBX"`
}

