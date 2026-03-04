package hl7v23

// ADT_A43_PATIENT - Group struct
type ADT_A43_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

// ADT_A43 - Move patient information - internal ID
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A43
type ADT_A43 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A43_PATIENT `hl7:"GROUP"`
}

