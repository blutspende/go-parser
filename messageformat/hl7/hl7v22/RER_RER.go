package hl7v22

// RER_RER_DEFINITION - Group struct
type RER_RER_DEFINITION struct {
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RER_RER_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RER_RER_DEFINITION_ORDER `hl7:"GROUP"`
}

// RER_RER_DEFINITION_PATIENT - Group struct
type RER_RER_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RER_RER_DEFINITION_ORDER - Group struct
type RER_RER_DEFINITION_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	PharmacyEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RER_RER - Pharmacy Encoded Order Information Response
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/RER_RER
type RER_RER struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []RER_RER_DEFINITION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

