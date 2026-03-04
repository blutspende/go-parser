package hl7v22

// RDS_O01_PATIENT - Group struct
type RDS_O01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// RDS_O01_ORDER - Group struct
type RDS_O01_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail RDS_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RDS_O01_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyDispense RXD `hl7:"TAG=RXD"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyEncodedOrder []RXE `hl7:"TAG=RXE;ATR=optional"`
	Observation []RDS_O01_ORDER_OBSERVATION `hl7:"GROUP"`
}

// RDS_O01_ORDER_ORDER_DETAIL - Group struct
type RDS_O01_ORDER_ORDER_DETAIL struct {
	PharmacyPrescriptionOrder RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndComments []NTE `hl7:"TAG=NTE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	Component RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT - Group struct
type RDS_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT struct {
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDS_O01_ORDER_ENCODING - Group struct
type RDS_O01_ORDER_ENCODING struct {
	PharmacyEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDS_O01_ORDER_OBSERVATION - Group struct
type RDS_O01_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDS_O01 - Pharmacy/treatment dispense information
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RDS_O01
type RDS_O01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDS_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDS_O01_ORDER `hl7:"GROUP"`
}

