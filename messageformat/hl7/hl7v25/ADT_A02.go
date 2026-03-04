package hl7v25

// ADT_A02 - Transfer a Patient
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ADT_A02
type ADT_A02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	Role1 []ROL `hl7:"TAG=ROL;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	Role2 []ROL `hl7:"TAG=ROL;ATR=optional"`
	Disability []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientDeathAndAutopsy PDA `hl7:"TAG=PDA;ATR=optional"`
}

