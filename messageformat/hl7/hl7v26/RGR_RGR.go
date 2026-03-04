package hl7v26

// RGR_RGR_DEFINITION - Group struct
type RGR_RGR_DEFINITION struct {
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RGR_RGR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RGR_RGR_DEFINITION_ORDER `hl7:"GROUP"`
}

// RGR_RGR_DEFINITION_PATIENT - Group struct
type RGR_RGR_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGR_RGR_DEFINITION_ORDER - Group struct
type RGR_RGR_DEFINITION_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Encoding RGR_RGR_DEFINITION_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentGive []RXG `hl7:"TAG=RXG"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGR_RGR_DEFINITION_ORDER_ENCODING - Group struct
type RGR_RGR_DEFINITION_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGR_RGR - Pharmacy/Treatment Dose Information Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/RGR_RGR
type RGR_RGR struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	Definition []RGR_RGR_DEFINITION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

