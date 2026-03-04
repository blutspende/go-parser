package hl7v22

// RRD_O02_PATIENT - Group struct
type RRD_O02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Order []RRD_O02_PATIENT_ORDER `hl7:"GROUP"`
}

// RRD_O02_PATIENT_ORDER - Group struct
type RRD_O02_PATIENT_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail RRD_O02_PATIENT_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// RRD_O02_PATIENT_ORDER_ORDER_DETAIL - Group struct
type RRD_O02_PATIENT_ORDER_ORDER_DETAIL struct {
	PharmacyDispense RXD `hl7:"TAG=RXD"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRD_O02 - Pharmacy/treatment dispense acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RRD_O02
type RRD_O02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RRD_O02_PATIENT `hl7:"GROUP;ATR=optional"`
}

