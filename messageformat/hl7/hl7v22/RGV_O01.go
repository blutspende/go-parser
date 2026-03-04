package hl7v22

// RGV_O01_PATIENT - Group struct
type RGV_O01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// RGV_O01_ORDER - Group struct
type RGV_O01_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail RGV_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RGV_O01_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	Give RGV_O01_ORDER_GIVE `hl7:"GROUP"`
}

// RGV_O01_ORDER_ORDER_DETAIL - Group struct
type RGV_O01_ORDER_ORDER_DETAIL struct {
	PharmacyPrescriptionOrder RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndComments []NTE `hl7:"TAG=NTE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	Component RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT - Group struct
type RGV_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT struct {
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGV_O01_ORDER_ENCODING - Group struct
type RGV_O01_ORDER_ENCODING struct {
	PharmacyEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGV_O01_ORDER_GIVE - Group struct
type RGV_O01_ORDER_GIVE struct {
	PharmacyGive RXG `hl7:"TAG=RXG"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation RGV_O01_ORDER_GIVE_OBSERVATION `hl7:"GROUP"`
}

// RGV_O01_ORDER_GIVE_OBSERVATION - Group struct
type RGV_O01_ORDER_GIVE_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGV_O01 - Pharmacy/treatment give message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RGV_O01
type RGV_O01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RGV_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RGV_O01_ORDER `hl7:"GROUP"`
}

