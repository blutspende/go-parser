package hl7v231

// VXR_V03_PATIENT_VISIT - Group struct
type VXR_V03_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// VXR_V03_INSURANCE - Group struct
type VXR_V03_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// VXR_V03_ORDER - Group struct
type VXR_V03_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	PharmacyTreatmentAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
	Observation []VXR_V03_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// VXR_V03_ORDER_OBSERVATION - Group struct
type VXR_V03_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// VXR_V03 - Vaccination record response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/VXR_V03
type VXR_V03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisit VXR_V03_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []VXR_V03_INSURANCE `hl7:"GROUP;ATR=optional"`
	Order []VXR_V03_ORDER `hl7:"GROUP;ATR=optional"`
}

