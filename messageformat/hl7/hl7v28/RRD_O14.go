package hl7v28

// RRD_O14_RESPONSE - Group struct
type RRD_O14_RESPONSE struct {
	Patient RRD_O14_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRD_O14_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRD_O14_RESPONSE_PATIENT - Group struct
type RRD_O14_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRD_O14_RESPONSE_ORDER - Group struct
type RRD_O14_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing []RRD_O14_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Dispense RRD_O14_RESPONSE_ORDER_DISPENSE `hl7:"GROUP;ATR=optional"`
}

// RRD_O14_RESPONSE_ORDER_TIMING - Group struct
type RRD_O14_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RRD_O14_RESPONSE_ORDER_DISPENSE - Group struct
type RRD_O14_RESPONSE_ORDER_DISPENSE struct {
	PharmacyTreatmentDispense RXD `hl7:"TAG=RXD"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRD_O14 - Pharmacy/treatment dispense acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RRD_O14
type RRD_O14 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRD_O14_RESPONSE `hl7:"GROUP;ATR=optional"`
}

