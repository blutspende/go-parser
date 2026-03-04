package hl7v23

// ADT_A13_PROCEDURE - Group struct
type ADT_A13_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// ADT_A13_INSURANCE - Group struct
type ADT_A13_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// ADT_A13 - Cancel discharge/end visit
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADT_A13
type ADT_A13 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Procedure []ADT_A13_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []ADT_A13_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Ub82Data UB1 `hl7:"TAG=UB1;ATR=optional"`
	Ub92Data UB2 `hl7:"TAG=UB2;ATR=optional"`
}

