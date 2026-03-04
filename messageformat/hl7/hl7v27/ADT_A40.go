package hl7v27

// ADT_A40_PATIENT - Group struct
type ADT_A40_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A40 - Merge patient - patient identifier list
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/ADT_A40
type ADT_A40 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A40_PATIENT `hl7:"GROUP"`
}

