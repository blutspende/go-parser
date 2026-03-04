package hl7v22

// ADT_A35 - account number only
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADT_A35
type ADT_A35 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

