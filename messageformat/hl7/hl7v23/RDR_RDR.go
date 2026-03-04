package hl7v23

// RDR_RDR_DEFINITION - Group struct
type RDR_RDR_DEFINITION struct {
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RDR_RDR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDR_RDR_DEFINITION_ORDER `hl7:"GROUP"`
}

// RDR_RDR_DEFINITION_PATIENT - Group struct
type RDR_RDR_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDR_RDR_DEFINITION_ORDER - Group struct
type RDR_RDR_DEFINITION_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Encoding RDR_RDR_DEFINITION_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	Dispense []RDR_RDR_DEFINITION_ORDER_DISPENSE `hl7:"GROUP"`
}

// RDR_RDR_DEFINITION_ORDER_ENCODING - Group struct
type RDR_RDR_DEFINITION_ORDER_ENCODING struct {
	PharmacyEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDR_RDR_DEFINITION_ORDER_DISPENSE - Group struct
type RDR_RDR_DEFINITION_ORDER_DISPENSE struct {
	PharmacyDispenseSegment RXD `hl7:"TAG=RXD"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDR_RDR - Pharmacy/treatment dispense information response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RDR_RDR
type RDR_RDR struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []RDR_RDR_DEFINITION `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

