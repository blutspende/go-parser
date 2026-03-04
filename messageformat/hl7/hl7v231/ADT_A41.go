package hl7v231

// ADT_A41_PATIENT - Group struct
type ADT_A41_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A41 - Merge account - patient account number
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A41
type ADT_A41 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	Patient []ADT_A41_PATIENT `hl7:"GROUP"`
}

