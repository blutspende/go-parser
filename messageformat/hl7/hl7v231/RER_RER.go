package hl7v231

// RER_RER_DEFINITION - Group struct
type RER_RER_DEFINITION struct {
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RER_RER_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RER_RER_DEFINITION_ORDER `hl7:"GROUP"`
}

// RER_RER_DEFINITION_PATIENT - Group struct
type RER_RER_DEFINITION_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RER_RER_DEFINITION_ORDER - Group struct
type RER_RER_DEFINITION_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RER_RER - Pharmacy/treatment encoded order information response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RER_RER
type RER_RER struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []RER_RER_DEFINITION `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

