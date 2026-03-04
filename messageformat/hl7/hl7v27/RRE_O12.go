package hl7v27

// RRE_O12_RESPONSE - Group struct
type RRE_O12_RESPONSE struct {
	Patient RRE_O12_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRE_O12_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRE_O12_RESPONSE_PATIENT - Group struct
type RRE_O12_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRE_O12_RESPONSE_ORDER - Group struct
type RRE_O12_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RRE_O12_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Encoding RRE_O12_RESPONSE_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// RRE_O12_RESPONSE_ORDER_TIMING - Group struct
type RRE_O12_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RRE_O12_RESPONSE_ORDER_ENCODING - Group struct
type RRE_O12_RESPONSE_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Timing_encoded []RRE_O12_RESPONSE_ORDER_ENCODING_TIMING_ENCODED `hl7:"GROUP"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RRE_O12_RESPONSE_ORDER_ENCODING_TIMING_ENCODED - Group struct
type RRE_O12_RESPONSE_ORDER_ENCODING_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RRE_O12 - Pharmacy/treatment encoded order acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/RRE_O12
type RRE_O12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRE_O12_RESPONSE `hl7:"GROUP;ATR=optional"`
}

