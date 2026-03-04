package hl7v251

// RAR_RAR_DEFINITION - Group struct
type RAR_RAR_DEFINITION struct {
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Patient RAR_RAR_DEFINITION_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RAR_RAR_DEFINITION_ORDER `hl7:"GROUP"`
}

// RAR_RAR_DEFINITION_PATIENT - Group struct
type RAR_RAR_DEFINITION_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAR_RAR_DEFINITION_ORDER - Group struct
type RAR_RAR_DEFINITION_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Encoding RAR_RAR_DEFINITION_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
}

// RAR_RAR_DEFINITION_ORDER_ENCODING - Group struct
type RAR_RAR_DEFINITION_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RAR_RAR - Pharmacy/Treatment Administration Information Response
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/RAR_RAR
type RAR_RAR struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	Definition []RAR_RAR_DEFINITION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

