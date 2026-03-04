package hl7v22

// RRG_O02_PATIENT - Group struct
type RRG_O02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Order []RRG_O02_PATIENT_ORDER `hl7:"GROUP"`
}

// RRG_O02_PATIENT_ORDER - Group struct
type RRG_O02_PATIENT_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail RRG_O02_PATIENT_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// RRG_O02_PATIENT_ORDER_ORDER_DETAIL - Group struct
type RRG_O02_PATIENT_ORDER_ORDER_DETAIL struct {
	PharmacyGive RXG `hl7:"TAG=RXG"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRG_O02 - Pharmacy/treatment give acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RRG_O02
type RRG_O02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RRG_O02_PATIENT `hl7:"GROUP;ATR=optional"`
}

