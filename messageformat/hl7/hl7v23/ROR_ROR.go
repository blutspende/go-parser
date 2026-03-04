package hl7v23

// ROR_ROR_DEFINITION - Group struct
type ROR_ROR_DEFINITION struct {
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient ROR_ROR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ROR_ROR_DEFINITION_ORDER `hl7:"GROUP"`
}

// ROR_ROR_DEFINITION_PATIENT - Group struct
type ROR_ROR_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ROR_ROR_DEFINITION_ORDER - Group struct
type ROR_ROR_DEFINITION_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	PharmacyPrescriptionOrderSegment RXO `hl7:"TAG=RXO"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// ROR_ROR - Pharmacy/treatment order response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ROR_ROR
type ROR_ROR struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []ROR_ROR_DEFINITION `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

