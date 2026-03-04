package hl7v28

// ADT_A49_PATIENT - Group struct
type ADT_A49_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

// ADT_A49 - Change patient account number
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ADT_A49
type ADT_A49 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A49_PATIENT `hl7:"GROUP"`
}

