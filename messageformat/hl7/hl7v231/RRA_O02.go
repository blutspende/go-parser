package hl7v231

// RRA_O02_RESPONSE - Group struct
type RRA_O02_RESPONSE struct {
	Patient RRA_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRA_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRA_O02_RESPONSE_PATIENT - Group struct
type RRA_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRA_O02_RESPONSE_ORDER - Group struct
type RRA_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Administration RRA_O02_RESPONSE_ORDER_ADMINISTRATION `hl7:"GROUP;ATR=optional"`
}

// RRA_O02_RESPONSE_ORDER_ADMINISTRATION - Group struct
type RRA_O02_RESPONSE_ORDER_ADMINISTRATION struct {
	PharmacyTreatmentAdministrationSegment []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR"`
}

// RRA_O02 - Pharmacy/treatment administration acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RRA_O02
type RRA_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRA_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

