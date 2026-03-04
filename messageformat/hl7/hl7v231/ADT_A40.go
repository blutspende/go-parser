package hl7v231

// ADT_A40_PATIENT - Group struct
type ADT_A40_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A40 - Merge patient - patient identifier list
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A40
type ADT_A40 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	Patient []ADT_A40_PATIENT `hl7:"GROUP"`
}

