package hl7v26

// RRG_O16_RESPONSE - Group struct
type RRG_O16_RESPONSE struct {
	Patient RRG_O16_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRG_O16_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRG_O16_RESPONSE_PATIENT - Group struct
type RRG_O16_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRG_O16_RESPONSE_ORDER - Group struct
type RRG_O16_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RRG_O16_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Give RRG_O16_RESPONSE_ORDER_GIVE `hl7:"GROUP;ATR=optional"`
}

// RRG_O16_RESPONSE_ORDER_TIMING - Group struct
type RRG_O16_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RRG_O16_RESPONSE_ORDER_GIVE - Group struct
type RRG_O16_RESPONSE_ORDER_GIVE struct {
	PharmacyTreatmentGive RXG `hl7:"TAG=RXG"`
	TimingGive []RRG_O16_RESPONSE_ORDER_GIVE_TIMING_GIVE `hl7:"GROUP"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRG_O16_RESPONSE_ORDER_GIVE_TIMING_GIVE - Group struct
type RRG_O16_RESPONSE_ORDER_GIVE_TIMING_GIVE struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RRG_O16 - Pharmacy/Treatment Give Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/RRG_O16
type RRG_O16 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRG_O16_RESPONSE `hl7:"GROUP;ATR=optional"`
}

