package hl7v22

// ADT_A21 - A Patient Goes On A "Leave Of Absence"
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADT_A21
type ADT_A21 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

