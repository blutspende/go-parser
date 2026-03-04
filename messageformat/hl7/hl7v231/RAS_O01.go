package hl7v231

// RAS_O01_PATIENT - Group struct
type RAS_O01_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit RAS_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RAS_O01_PATIENT_PATIENT_VISIT - Group struct
type RAS_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RAS_O01_ORDER - Group struct
type RAS_O01_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail RAS_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RAS_O01_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentAdministrationSegment []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR"`
	Observation []RAS_O01_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// RAS_O01_ORDER_ORDER_DETAIL - Group struct
type RAS_O01_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RAS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RAS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RAS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	Components RAS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS `hl7:"GROUP;ATR=optional"`
}

// RAS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS - Group struct
type RAS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS struct {
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O01_ORDER_ENCODING - Group struct
type RAS_O01_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RAS_O01_ORDER_OBSERVATION - Group struct
type RAS_O01_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O01 - Pharmacy/treatment administration message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RAS_O01
type RAS_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RAS_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RAS_O01_ORDER `hl7:"GROUP"`
}

