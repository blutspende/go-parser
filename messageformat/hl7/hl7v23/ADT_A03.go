package hl7v23

// ADT_A03_PROCEDURE - Group struct
type ADT_A03_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// ADT_A03 - Discharge/end visit
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A03
type ADT_A03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment []DB1 `hl7:"TAG=DB1;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Procedure []ADT_A03_PROCEDURE `hl7:"GROUP;ATR=optional"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
}

