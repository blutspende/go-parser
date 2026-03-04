package hl7v24

// ORG_O20_RESPONSE - Group struct
type ORG_O20_RESPONSE struct {
	Patient ORG_O20_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORG_O20_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORG_O20_RESPONSE_PATIENT - Group struct
type ORG_O20_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORG_O20_RESPONSE_ORDER - Group struct
type ORG_O20_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORG_O20 - General clinical order response
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORG_O20
type ORG_O20 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORG_O20_RESPONSE `hl7:"GROUP;ATR=optional"`
}

