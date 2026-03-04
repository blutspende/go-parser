package hl7v231

// RAR_RAR_DEFINITION - Group struct
type RAR_RAR_DEFINITION struct {
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RAR_RAR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RAR_RAR_DEFINITION_ORDER `hl7:"GROUP"`
}

// RAR_RAR_DEFINITION_PATIENT - Group struct
type RAR_RAR_DEFINITION_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAR_RAR_DEFINITION_ORDER - Group struct
type RAR_RAR_DEFINITION_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Encoding RAR_RAR_DEFINITION_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentAdministrationSegment []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR"`
}

// RAR_RAR_DEFINITION_ORDER_ENCODING - Group struct
type RAR_RAR_DEFINITION_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrderSegment []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RAR_RAR - Pharmacy/treatment administration information response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RAR_RAR
type RAR_RAR struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []RAR_RAR_DEFINITION `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

