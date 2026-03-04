package hl7v231

// BAR_P06_PATIENT - Group struct
type BAR_P06_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// BAR_P06 - End account
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/BAR_P06
type BAR_P06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	Patient []BAR_P06_PATIENT `hl7:"GROUP"`
}

