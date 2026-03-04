package hl7v23

// ADT_A47 - Change internal ID
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A47
type ADT_A47 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

