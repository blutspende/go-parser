package hl7v24

// RGR_RGR_Definition - Group struct
type RGR_RGR_Definition struct {
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RGR_RGR_Definition_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RGR_RGR_Definition_ORDER `hl7:"GROUP"`
}

// RGR_RGR_Definition_PATIENT - Group struct
type RGR_RGR_Definition_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGR_RGR_Definition_ORDER - Group struct
type RGR_RGR_Definition_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Encoding RGR_RGR_Definition_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentGive []RXG `hl7:"TAG=RXG"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGR_RGR_Definition_ORDER_ENCODING - Group struct
type RGR_RGR_Definition_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGR_RGR - Pharmacy/treatment dose information response
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RGR_RGR
type RGR_RGR struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []RGR_RGR_Definition `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

