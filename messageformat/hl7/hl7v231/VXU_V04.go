package hl7v231

// VXU_V04_PATIENT - Group struct
type VXU_V04_PATIENT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// VXU_V04_INSURANCE - Group struct
type VXU_V04_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// VXU_V04_ORDER - Group struct
type VXU_V04_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	PharmacyTreatmentAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
	Observation []VXU_V04_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// VXU_V04_ORDER_OBSERVATION - Group struct
type VXU_V04_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// VXU_V04 - Unsolicited vaccination record update
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/VXU_V04
type VXU_V04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Patient VXU_V04_PATIENT `hl7:"GROUP;ATR=optional"`
	Insurance []VXU_V04_INSURANCE `hl7:"GROUP;ATR=optional"`
	Order []VXU_V04_ORDER `hl7:"GROUP;ATR=optional"`
}

