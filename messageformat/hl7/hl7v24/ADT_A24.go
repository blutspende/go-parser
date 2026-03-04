package hl7v24

// ADT_A24 - Link patient information
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ADT_A24
type ADT_A24 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification1 PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic1 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit1 PV1 `hl7:"TAG=PV1;ATR=optional"`
	Disability1 []DB1 `hl7:"TAG=DB1;ATR=optional"`
	PatientIdentification2 PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic2 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit2 PV1 `hl7:"TAG=PV1;ATR=optional"`
	Disability2 []DB1 `hl7:"TAG=DB1;ATR=optional"`
}

