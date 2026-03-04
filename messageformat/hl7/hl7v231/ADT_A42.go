package hl7v231

// ADT_A42_PATIENT - Group struct
type ADT_A42_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A42 - Merge visit - visit number
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A42
type ADT_A42 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	Patient []ADT_A42_PATIENT `hl7:"GROUP"`
}

