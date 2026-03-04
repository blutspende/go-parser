package hl7v231

// RRO_O02_RESPONSE - Group struct
type RRO_O02_RESPONSE struct {
	Patient RRO_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRO_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRO_O02_RESPONSE_PATIENT - Group struct
type RRO_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRO_O02_RESPONSE_ORDER - Group struct
type RRO_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail RRO_O02_RESPONSE_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// RRO_O02_RESPONSE_ORDER_ORDER_DETAIL - Group struct
type RRO_O02_RESPONSE_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO"`
	NotesAndCommentsSegment1 []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
	NotesAndCommentsSegment2 []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRO_O02 - Pharmacy/treatment Order acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RRO_O02
type RRO_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRO_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

