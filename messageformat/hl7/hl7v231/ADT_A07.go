package hl7v231

// ADT_A07_PROCEDURE - Group struct
type ADT_A07_PROCEDURE struct {
	ProceduresSegment PR1 `hl7:"TAG=PR1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// ADT_A07_INSURANCE - Group struct
type ADT_A07_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment []IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// ADT_A07 - Change an inpatient to an outpatient
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ADT_A07
type ADT_A07 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformationSegment MRG `hl7:"TAG=MRG;ATR=optional"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroupSegment DRG `hl7:"TAG=DRG;ATR=optional"`
	Procedure []ADT_A07_PROCEDURE `hl7:"GROUP;ATR=optional"`
	GuarantorSegment []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []ADT_A07_INSURANCE `hl7:"GROUP;ATR=optional"`
	AccidentSegment ACC `hl7:"TAG=ACC;ATR=optional"`
	Ub82DataSegment UB1 `hl7:"TAG=UB1;ATR=optional"`
	Ub92DataSegment UB2 `hl7:"TAG=UB2;ATR=optional"`
}

