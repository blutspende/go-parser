package hl7v271

// BAR_P06_PATIENT - Group struct
type BAR_P06_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// BAR_P06 - End account
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/BAR_P06
type BAR_P06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []BAR_P06_PATIENT `hl7:"GROUP"`
}

