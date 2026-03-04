package hl7v23

// RGR_RGR_DEFINITION - Group struct
type RGR_RGR_DEFINITION struct {
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RGR_RGR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RGR_RGR_DEFINITION_ORDER `hl7:"GROUP"`
}

// RGR_RGR_DEFINITION_PATIENT - Group struct
type RGR_RGR_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGR_RGR_DEFINITION_ORDER - Group struct
type RGR_RGR_DEFINITION_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Encoding RGR_RGR_DEFINITION_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyGiveSegment []RXG `hl7:"TAG=RXG"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGR_RGR_DEFINITION_ORDER_ENCODING - Group struct
type RGR_RGR_DEFINITION_ORDER_ENCODING struct {
	PharmacyEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGR_RGR - Pharmacy/treatment dose information response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RGR_RGR
type RGR_RGR struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []RGR_RGR_DEFINITION `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

