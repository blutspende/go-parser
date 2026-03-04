package hl7v22

// ROR_ROR_DEFINITION - Group struct
type ROR_ROR_DEFINITION struct {
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient ROR_ROR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ROR_ROR_DEFINITION_ORDER `hl7:"GROUP"`
}

// ROR_ROR_DEFINITION_PATIENT - Group struct
type ROR_ROR_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ROR_ROR_DEFINITION_ORDER - Group struct
type ROR_ROR_DEFINITION_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	PharmacyPrescriptionOrder RXO `hl7:"TAG=RXO"`
	PharmacyRoute []RXR `hl7:"TAG=RXR"`
	PharmacyComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// ROR_ROR - Pharmacy Prescription Order Response
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ROR_ROR
type ROR_ROR struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	Definition []ROR_ROR_DEFINITION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

