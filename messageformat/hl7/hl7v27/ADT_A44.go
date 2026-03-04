package hl7v27

// ADT_A44_PATIENT - Group struct
type ADT_A44_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

// ADT_A44 - Move account information - patient account number
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/ADT_A44
type ADT_A44 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []ADT_A44_PATIENT `hl7:"GROUP"`
}

