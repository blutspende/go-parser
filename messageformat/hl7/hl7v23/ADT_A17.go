package hl7v23

// ADT_A17 - Swap patients
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A17
type ADT_A17 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification1 PID `hl7:"TAG=PID"`
	PatientDemographic1 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit1 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation1 PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment1 []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationSegment1 []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientIdentification2 PID `hl7:"TAG=PID"`
	PatientDemographic2 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit2 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation2 PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment2 []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationSegment2 []OBX `hl7:"TAG=OBX;ATR=optional"`
}

