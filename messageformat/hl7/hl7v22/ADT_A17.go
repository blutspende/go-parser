package hl7v22

// ADT_A17 - Swap Patients
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADT_A17
type ADT_A17 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification1 PID `hl7:"TAG=PID"`
	PatientVisit1 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation1 PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResult1 []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientIdentification2 PID `hl7:"TAG=PID"`
	PatientVisit2 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation2 PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResult2 []OBX `hl7:"TAG=OBX;ATR=optional"`
}

