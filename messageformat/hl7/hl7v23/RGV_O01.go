package hl7v23

// RGV_O01_PATIENT - Group struct
type RGV_O01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit RGV_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RGV_O01_PATIENT_PATIENT_VISIT - Group struct
type RGV_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RGV_O01_ORDER - Group struct
type RGV_O01_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail RGV_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RGV_O01_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	Give []RGV_O01_ORDER_GIVE `hl7:"GROUP"`
}

// RGV_O01_ORDER_ORDER_DETAIL - Group struct
type RGV_O01_ORDER_ORDER_DETAIL struct {
	PharmacyPrescriptionOrderSegment RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	Components RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS `hl7:"GROUP;ATR=optional"`
}

// RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS - Group struct
type RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS struct {
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGV_O01_ORDER_ENCODING - Group struct
type RGV_O01_ORDER_ENCODING struct {
	PharmacyEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGV_O01_ORDER_GIVE - Group struct
type RGV_O01_ORDER_GIVE struct {
	PharmacyGiveSegment RXG `hl7:"TAG=RXG"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RGV_O01_ORDER_GIVE_OBSERVATION `hl7:"GROUP"`
}

// RGV_O01_ORDER_GIVE_OBSERVATION - Group struct
type RGV_O01_ORDER_GIVE_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGV_O01 - Pharmacy/treatment give message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RGV_O01
type RGV_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RGV_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RGV_O01_ORDER `hl7:"GROUP"`
}

