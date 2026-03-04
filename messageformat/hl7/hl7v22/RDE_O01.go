package hl7v22

// RDE_O01_PATIENT - Group struct
type RDE_O01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// RDE_O01_ORDER - Group struct
type RDE_O01_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail RDE_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	PharmacyEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RDE_O01_ORDER_OBSERVATION `hl7:"GROUP"`
}

// RDE_O01_ORDER_ORDER_DETAIL - Group struct
type RDE_O01_ORDER_ORDER_DETAIL struct {
	PharmacyPrescriptionOrder RXO `hl7:"TAG=RXO"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	Component RDE_O01_ORDER_ORDER_DETAIL_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDE_O01_ORDER_ORDER_DETAIL_COMPONENT - Group struct
type RDE_O01_ORDER_ORDER_DETAIL_COMPONENT struct {
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O01_ORDER_OBSERVATION - Group struct
type RDE_O01_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O01 - Pharmacy/treatment encoded order message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RDE_O01
type RDE_O01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDE_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDE_O01_ORDER `hl7:"GROUP"`
}

