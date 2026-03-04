package hl7v22

// ADT_A18 - Merge Patient Information
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADT_A18
type ADT_A18 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	MergePatientInformation MRG `hl7:"TAG=MRG;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
}

