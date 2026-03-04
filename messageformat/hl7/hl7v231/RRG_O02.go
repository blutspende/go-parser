package hl7v231

// RRG_O02_RESPONSE - Group struct
type RRG_O02_RESPONSE struct {
	Patient RRG_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRG_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRG_O02_RESPONSE_PATIENT - Group struct
type RRG_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRG_O02_RESPONSE_ORDER - Group struct
type RRG_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Give RRG_O02_RESPONSE_ORDER_GIVE `hl7:"GROUP;ATR=optional"`
}

// RRG_O02_RESPONSE_ORDER_GIVE - Group struct
type RRG_O02_RESPONSE_ORDER_GIVE struct {
	PharmacyTreatmentGiveSegment RXG `hl7:"TAG=RXG"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRG_O02 - Pharmacy/treatment give acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RRG_O02
type RRG_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRG_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

