package hl7v23

// RAS_O02_PATIENT - Group struct
type RAS_O02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit RAS_O02_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RAS_O02_PATIENT_PATIENT_VISIT - Group struct
type RAS_O02_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RAS_O02_ORDER - Group struct
type RAS_O02_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail RAS_O02_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RAS_O02_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyAdministrationSegment []RXA `hl7:"TAG=RXA"`
	PharmacyRouteSegment RXR `hl7:"TAG=RXR"`
	Observation []RAS_O02_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// RAS_O02_ORDER_ORDER_DETAIL - Group struct
type RAS_O02_ORDER_ORDER_DETAIL struct {
	PharmacyPrescriptionOrderSegment RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	Components RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS `hl7:"GROUP;ATR=optional"`
}

// RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS - Group struct
type RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS struct {
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O02_ORDER_ENCODING - Group struct
type RAS_O02_ORDER_ENCODING struct {
	PharmacyEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RAS_O02_ORDER_OBSERVATION - Group struct
type RAS_O02_ORDER_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O02 - Pharmacy/treatment administration message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RAS_O02
type RAS_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RAS_O02_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RAS_O02_ORDER `hl7:"GROUP"`
}

