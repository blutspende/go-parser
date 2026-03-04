package hl7v25

// RER_RER_DEFINITION - Group struct
type RER_RER_DEFINITION struct {
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
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
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RER_RER - Pharmacy/Treatment Encoded Order Information Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/RER_RER
type RER_RER struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	Definition []RER_RER_DEFINITION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

