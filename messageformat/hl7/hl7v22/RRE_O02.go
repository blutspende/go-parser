package hl7v22

// RRE_O02_PATIENT - Group struct
type RRE_O02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Order []RRE_O02_PATIENT_ORDER `hl7:"GROUP"`
}

// RRE_O02_PATIENT_ORDER - Group struct
type RRE_O02_PATIENT_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail RRE_O02_PATIENT_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// RRE_O02_PATIENT_ORDER_ORDER_DETAIL - Group struct
type RRE_O02_PATIENT_ORDER_ORDER_DETAIL struct {
	PharmacyEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRE_O02 - Pharmacy/treatment encoded order acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RRE_O02
type RRE_O02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RRE_O02_PATIENT `hl7:"GROUP;ATR=optional"`
}

