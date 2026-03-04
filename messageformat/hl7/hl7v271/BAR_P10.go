package hl7v271

// BAR_P10_PROCEDURE - Group struct
type BAR_P10_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	GroupingReimbursementProcedureLineItem GP2 `hl7:"TAG=GP2;ATR=optional"`
}

// BAR_P10 - Transmit Ambulatory Payment  Classification(APC)
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/BAR_P10
type BAR_P10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	GroupingReimbursementVisit GP1 `hl7:"TAG=GP1"`
	Procedure []BAR_P10_PROCEDURE `hl7:"GROUP;ATR=optional"`
}

