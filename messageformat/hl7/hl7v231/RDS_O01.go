package hl7v231

// RDS_O01_PATIENT - Group struct
type RDS_O01_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit RDS_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RDS_O01_PATIENT_PATIENT_VISIT - Group struct
type RDS_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RDS_O01_ORDER - Group struct
type RDS_O01_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail RDS_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RDS_O01_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentDispenseSegment RXD `hl7:"TAG=RXD"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RDS_O01_ORDER_OBSERVATION `hl7:"GROUP"`
}

// RDS_O01_ORDER_ORDER_DETAIL - Group struct
type RDS_O01_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	Component RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT - Group struct
type RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT struct {
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDS_O01_ORDER_ENCODING - Group struct
type RDS_O01_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDS_O01_ORDER_OBSERVATION - Group struct
type RDS_O01_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDS_O01 - Pharmacy/treatment dispense information
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RDS_O01
type RDS_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDS_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDS_O01_ORDER `hl7:"GROUP"`
}

