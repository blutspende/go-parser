package hl7v26

// ADT_A34 - Merge Patient Information - Patient ID Only
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ADT_A34
type ADT_A34 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	AccessRestrictions []ARV `hl7:"TAG=ARV;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

