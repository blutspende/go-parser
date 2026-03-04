package hl7v231

// ADT_A50 - Change visit number
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A50
type ADT_A50 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
}

