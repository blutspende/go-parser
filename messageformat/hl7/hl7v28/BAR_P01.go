package hl7v28

// BAR_P01_VISIT - Group struct
type BAR_P01_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	Disability []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Procedure []BAR_P01_VISIT_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Insurance []BAR_P01_VISIT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Ub82 UB1 `hl7:"TAG=UB1;ATR=optional"`
	UniformBillingData UB2 `hl7:"TAG=UB2;ATR=optional"`
}

// BAR_P01_VISIT_PROCEDURE - Group struct
type BAR_P01_VISIT_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// BAR_P01_VISIT_INSURANCE - Group struct
type BAR_P01_VISIT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification []IN3 `hl7:"TAG=IN3;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// BAR_P01 - Add patient accounts
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/BAR_P01
type BAR_P01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	Visit []BAR_P01_VISIT `hl7:"GROUP"`
}

