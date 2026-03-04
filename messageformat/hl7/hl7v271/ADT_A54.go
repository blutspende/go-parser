package hl7v271

// ADT_A54 - Change attending doctor
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/ADT_A54
type ADT_A54 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	Role1 []ROL `hl7:"TAG=ROL;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	Role2 []ROL `hl7:"TAG=ROL;ATR=optional"`
}

