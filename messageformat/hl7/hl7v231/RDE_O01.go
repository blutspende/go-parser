package hl7v231

// RDE_O01_PATIENT - Group struct
type RDE_O01_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit RDE_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []RDE_O01_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	GuarantorSegment GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// RDE_O01_PATIENT_PATIENT_VISIT - Group struct
type RDE_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RDE_O01_PATIENT_INSURANCE - Group struct
type RDE_O01_PATIENT_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RDE_O01_ORDER - Group struct
type RDE_O01_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail RDE_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RDE_O01_ORDER_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// RDE_O01_ORDER_ORDER_DETAIL - Group struct
type RDE_O01_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	Component RDE_O01_ORDER_ORDER_DETAIL_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDE_O01_ORDER_ORDER_DETAIL_COMPONENT - Group struct
type RDE_O01_ORDER_ORDER_DETAIL_COMPONENT struct {
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O01_ORDER_OBSERVATION - Group struct
type RDE_O01_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O01 - Pharmacy/treatment encoded order message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RDE_O01
type RDE_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDE_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDE_O01_ORDER `hl7:"GROUP"`
}

