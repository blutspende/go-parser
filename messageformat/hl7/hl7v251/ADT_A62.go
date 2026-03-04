package hl7v251

// ADT_A62 - Cancel Change Consulting Doctor
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/ADT_A62
type ADT_A62 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

