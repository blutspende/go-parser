package hl7v23

// ADT_A39_PATIENT - Group struct
type ADT_A39_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A39 - Merge person - external ID
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A39
type ADT_A39 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A39_PATIENT `hl7:"GROUP"`
}

