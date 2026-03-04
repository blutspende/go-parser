package hl7v25

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
	Timing []ORG_O20_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Specimen []ORG_O20_RESPONSE_ORDER_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// ORG_O20_RESPONSE_ORDER_TIMING - Group struct
type ORG_O20_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORG_O20_RESPONSE_ORDER_SPECIMEN - Group struct
type ORG_O20_RESPONSE_ORDER_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	SpecimenContainerDetail []SAC `hl7:"TAG=SAC;ATR=optional"`
}

// ORG_O20 - General Clinical Order Acknowledgement 
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ORG_O20
type ORG_O20 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORG_O20_RESPONSE `hl7:"GROUP;ATR=optional"`
}

