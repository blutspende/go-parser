package hl7v24

// RRE_O12_RESPONSE - Group struct
type RRE_O12_RESPONSE struct {
	Patient RRE_O12_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRE_O12_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRE_O12_RESPONSE_PATIENT - Group struct
type RRE_O12_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRE_O12_RESPONSE_ORDER - Group struct
type RRE_O12_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Encoding RRE_O12_RESPONSE_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
}

// RRE_O12_RESPONSE_ORDER_ENCODING - Group struct
type RRE_O12_RESPONSE_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRE_O12 - Pharmacy/treatment encoded order acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RRE_O12
type RRE_O12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRE_O12_RESPONSE `hl7:"GROUP;ATR=optional"`
}

