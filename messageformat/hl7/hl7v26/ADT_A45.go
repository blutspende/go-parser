package hl7v26

// ADT_A45_MERGE_INFO - Group struct
type ADT_A45_MERGE_INFO struct {
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
}

// ADT_A45 - Move Visit Information - Visit Number
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ADT_A45
type ADT_A45 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergeInfo []ADT_A45_MERGE_INFO `hl7:"GROUP"`
}

