package hl7v231

// RRE_O02_RESPONSE - Group struct
type RRE_O02_RESPONSE struct {
	Patient RRE_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRE_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRE_O02_RESPONSE_PATIENT - Group struct
type RRE_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRE_O02_RESPONSE_ORDER - Group struct
type RRE_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Encoding RRE_O02_RESPONSE_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
}

// RRE_O02_RESPONSE_ORDER_ENCODING - Group struct
type RRE_O02_RESPONSE_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRE_O02 - Pharmacy/treatment encoded order acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RRE_O02
type RRE_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRE_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

