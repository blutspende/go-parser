package hl7v25

// ADT_A08_PROCEDURE - Group struct
type ADT_A08_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// ADT_A08_INSURANCE - Group struct
type ADT_A08_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification []IN3 `hl7:"TAG=IN3;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// ADT_A08 - Update Patient Information
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ADT_A08
type ADT_A08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	Role1 []ROL `hl7:"TAG=ROL;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	Role2 []ROL `hl7:"TAG=ROL;ATR=optional"`
	Disability []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Procedure []ADT_A08_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []ADT_A08_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Ub82Data UB1 `hl7:"TAG=UB1;ATR=optional"`
	Ub92Data UB2 `hl7:"TAG=UB2;ATR=optional"`
	PatientDeathAndAutopsy PDA `hl7:"TAG=PDA;ATR=optional"`
}

