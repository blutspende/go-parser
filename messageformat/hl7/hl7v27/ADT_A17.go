package hl7v27

// ADT_A17 - Swap patients
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/ADT_A17
type ADT_A17 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification1 PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic1 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit1 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation1 PV2 `hl7:"TAG=PV2;ATR=optional"`
	Disability1 []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResult1 []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientIdentification2 PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic2 PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit2 PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation2 PV2 `hl7:"TAG=PV2;ATR=optional"`
	Disability2 []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResult2 []OBX `hl7:"TAG=OBX;ATR=optional"`
}

