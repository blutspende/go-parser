package hl7v231

// ADT_A29 - Delete person information
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A29
type ADT_A29 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
}

