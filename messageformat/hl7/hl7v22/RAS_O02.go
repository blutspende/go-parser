package hl7v22

// RAS_O02_PATIENT - Group struct
type RAS_O02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// RAS_O02_ORDER - Group struct
type RAS_O02_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetail RAS_O02_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RAS_O02_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyRoute RXR `hl7:"TAG=RXR"`
	Observation []RAS_O02_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// RAS_O02_ORDER_ORDER_DETAIL - Group struct
type RAS_O02_ORDER_ORDER_DETAIL struct {
	PharmacyPrescriptionOrder RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndComments []NTE `hl7:"TAG=NTE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	Components RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS `hl7:"GROUP;ATR=optional"`
}

// RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS - Group struct
type RAS_O02_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS struct {
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O02_ORDER_ENCODING - Group struct
type RAS_O02_ORDER_ENCODING struct {
	PharmacyEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RAS_O02_ORDER_OBSERVATION - Group struct
type RAS_O02_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O02 - Pharmacy/treatment administration message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RAS_O02
type RAS_O02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RAS_O02_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RAS_O02_ORDER `hl7:"GROUP"`
}

