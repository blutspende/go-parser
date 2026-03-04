package hl7v271

// ADT_A42_PATIENT - Group struct
type ADT_A42_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A42 - Merge visit - visit number
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/ADT_A42
type ADT_A42 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A42_PATIENT `hl7:"GROUP"`
}

