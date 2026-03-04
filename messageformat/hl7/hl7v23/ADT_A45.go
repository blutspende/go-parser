package hl7v23

// ADT_A45_MERGE_INFO - Group struct
type ADT_A45_MERGE_INFO struct {
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
}

// ADT_A45 - Move visit information - visit number
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A45
type ADT_A45 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergeInfo []ADT_A45_MERGE_INFO `hl7:"GROUP"`
}

