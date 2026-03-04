package hl7v27

// ADT_A43_PATIENT - Group struct
type ADT_A43_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

// ADT_A43 - Move patient information - patient identifier list
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/ADT_A43
type ADT_A43 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A43_PATIENT `hl7:"GROUP"`
}

