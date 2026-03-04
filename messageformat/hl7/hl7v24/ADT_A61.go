package hl7v24

// ADT_A61 - Change consulting doctor
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ADT_A61
type ADT_A61 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

