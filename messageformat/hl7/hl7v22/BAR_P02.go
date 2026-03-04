package hl7v22

// BAR_P02_PATIENT - Group struct
type BAR_P02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// BAR_P02 - Purge Patient Accounts
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/BAR_P02
type BAR_P02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	Patient []BAR_P02_PATIENT `hl7:"GROUP"`
}

