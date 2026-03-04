package hl7v23

// ADT_A44_PATIENT - Group struct
type ADT_A44_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

// ADT_A44 - Move account information - patient account number
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A44
type ADT_A44 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A44_PATIENT `hl7:"GROUP"`
}

