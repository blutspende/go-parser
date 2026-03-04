package hl7v231

// ADT_A35 - Merge patient information - account number only
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A35
type ADT_A35 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG"`
}

