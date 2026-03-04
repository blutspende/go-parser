package hl7v23

// RRD_O02_RESPONSE - Group struct
type RRD_O02_RESPONSE struct {
	Patient RRD_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRD_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRD_O02_RESPONSE_PATIENT - Group struct
type RRD_O02_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRD_O02_RESPONSE_ORDER - Group struct
type RRD_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Dispense RRD_O02_RESPONSE_ORDER_DISPENSE `hl7:"GROUP;ATR=optional"`
}

// RRD_O02_RESPONSE_ORDER_DISPENSE - Group struct
type RRD_O02_RESPONSE_ORDER_DISPENSE struct {
	PharmacyDispenseSegment RXD `hl7:"TAG=RXD"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRD_O02 - Pharmacy/treatment dispense acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RRD_O02
type RRD_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRD_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

