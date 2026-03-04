package hl7v24

// ADT_A36 - Merge patient information - patient id and account number
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ADT_A36
type ADT_A36 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

