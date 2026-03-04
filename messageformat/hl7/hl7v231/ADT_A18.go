package hl7v231

// ADT_A18 - Merge patient information
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A18
type ADT_A18 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG;ATR=optional"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
}

