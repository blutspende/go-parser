package hl7v23

// RDE_O01_PATIENT - Group struct
type RDE_O01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit RDE_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []RDE_O01_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// RDE_O01_PATIENT_PATIENT_VISIT - Group struct
type RDE_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RDE_O01_PATIENT_INSURANCE - Group struct
type RDE_O01_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RDE_O01_ORDER - Group struct
type RDE_O01_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail RDE_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	PharmacyEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RDE_O01_ORDER_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// RDE_O01_ORDER_ORDER_DETAIL - Group struct
type RDE_O01_ORDER_ORDER_DETAIL struct {
	PharmacyPrescriptionOrderSegment RXO `hl7:"TAG=RXO"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	Component RDE_O01_ORDER_ORDER_DETAIL_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDE_O01_ORDER_ORDER_DETAIL_COMPONENT - Group struct
type RDE_O01_ORDER_ORDER_DETAIL_COMPONENT struct {
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O01_ORDER_OBSERVATION - Group struct
type RDE_O01_ORDER_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O01 - Pharmacy/treatment encoded order message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RDE_O01
type RDE_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDE_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDE_O01_ORDER `hl7:"GROUP"`
}

