package hl7v231

// ADT_A43_PATIENT - Group struct
type ADT_A43_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG"`
}

// ADT_A43 - Move patient information - patient identifier list
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A43
type ADT_A43 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	Patient []ADT_A43_PATIENT `hl7:"GROUP"`
}

