package hl7v22

// RDR_RDR_DEFINITION - Group struct
type RDR_RDR_DEFINITION struct {
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RDR_RDR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDR_RDR_DEFINITION_ORDER `hl7:"GROUP"`
}

// RDR_RDR_DEFINITION_PATIENT - Group struct
type RDR_RDR_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDR_RDR_DEFINITION_ORDER - Group struct
type RDR_RDR_DEFINITION_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Encoding RDR_RDR_DEFINITION_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	Dispense []RDR_RDR_DEFINITION_ORDER_DISPENSE `hl7:"GROUP"`
}

// RDR_RDR_DEFINITION_ORDER_ENCODING - Group struct
type RDR_RDR_DEFINITION_ORDER_ENCODING struct {
	PharmacyEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyRoute RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDR_RDR_DEFINITION_ORDER_DISPENSE - Group struct
type RDR_RDR_DEFINITION_ORDER_DISPENSE struct {
	PharmacyDispense RXD `hl7:"TAG=RXD"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
}

// RDR_RDR - Pharmacy Dispense Information Response
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RDR_RDR
type RDR_RDR struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []RDR_RDR_DEFINITION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

