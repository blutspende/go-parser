package hl7v231

// BAR_P02_PATIENT - Group struct
type BAR_P02_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
	DisabilitySegment []DB1 `hl7:"TAG=DB1;ATR=optional"`
}

// BAR_P02 - Purge patient accounts
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/BAR_P02
type BAR_P02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	Patient []BAR_P02_PATIENT `hl7:"GROUP"`
}

