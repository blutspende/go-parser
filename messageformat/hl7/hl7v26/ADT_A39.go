package hl7v26

// ADT_A39_PATIENT - Group struct
type ADT_A39_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A39 - Merge Person - Patient ID
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ADT_A39
type ADT_A39 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A39_PATIENT `hl7:"GROUP"`
}

