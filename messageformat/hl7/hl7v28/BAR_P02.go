package hl7v28

// BAR_P02_PATIENT - Group struct
type BAR_P02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	Disability []DB1 `hl7:"TAG=DB1;ATR=optional"`
}

// BAR_P02 - Purge patient accounts
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/BAR_P02
type BAR_P02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []BAR_P02_PATIENT `hl7:"GROUP"`
}

