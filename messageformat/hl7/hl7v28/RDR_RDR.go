package hl7v28

// RDR_RDR_DEFINITION - Group struct
type RDR_RDR_DEFINITION struct {
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
	Timing []RDR_RDR_DEFINITION_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Encoding RDR_RDR_DEFINITION_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	Dispense []RDR_RDR_DEFINITION_ORDER_DISPENSE `hl7:"GROUP"`
}

// RDR_RDR_DEFINITION_ORDER_TIMING - Group struct
type RDR_RDR_DEFINITION_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RDR_RDR_DEFINITION_ORDER_ENCODING - Group struct
type RDR_RDR_DEFINITION_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	Timing_encoded []RDR_RDR_DEFINITION_ORDER_ENCODING_TIMING_ENCODED `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDR_RDR_DEFINITION_ORDER_ENCODING_TIMING_ENCODED - Group struct
type RDR_RDR_DEFINITION_ORDER_ENCODING_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RDR_RDR_DEFINITION_ORDER_DISPENSE - Group struct
type RDR_RDR_DEFINITION_ORDER_DISPENSE struct {
	PharmacyTreatmentDispense RXD `hl7:"TAG=RXD"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDR_RDR - Pharmacy/treatment Dispense Information
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RDR_RDR
type RDR_RDR struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Definition []RDR_RDR_DEFINITION `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

