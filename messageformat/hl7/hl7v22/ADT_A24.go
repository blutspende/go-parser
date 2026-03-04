package hl7v22

// ADT_A24 - Link Patient Information
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADT_A24
type ADT_A24 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification1 PID `hl7:"TAG=PID"`
	PatientVisit1 PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientIdentification2 PID `hl7:"TAG=PID"`
	PatientVisit2 PV1 `hl7:"TAG=PV1;ATR=optional"`
}

