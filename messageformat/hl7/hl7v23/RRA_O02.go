package hl7v23

// RRA_O02_RESPONSE - Group struct
type RRA_O02_RESPONSE struct {
	Patient RRA_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRA_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRA_O02_RESPONSE_PATIENT - Group struct
type RRA_O02_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRA_O02_RESPONSE_ORDER - Group struct
type RRA_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Administration RRA_O02_RESPONSE_ORDER_ADMINISTRATION `hl7:"GROUP;ATR=optional"`
}

// RRA_O02_RESPONSE_ORDER_ADMINISTRATION - Group struct
type RRA_O02_RESPONSE_ORDER_ADMINISTRATION struct {
	PharmacyAdministrationSegment []RXA `hl7:"TAG=RXA"`
	PharmacyRouteSegment RXR `hl7:"TAG=RXR"`
}

// RRA_O02 - Pharmacy/treatment administration acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RRA_O02
type RRA_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRA_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

