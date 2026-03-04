package hl7v231

// ADT_A17 - Swap patients
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A17
type ADT_A17 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment1 PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic1 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisitSegment1 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment1 PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment1 []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResultSegment1 []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientIdentificationSegment2 PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic2 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisitSegment2 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment2 PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment2 []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResultSegment2 []OBX `hl7:"TAG=OBX;ATR=optional"`
}

