package hl7v22

// ADT_A36 - pat ID and account number
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADT_A36
type ADT_A36 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

