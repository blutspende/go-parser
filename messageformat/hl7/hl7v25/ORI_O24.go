package hl7v25

// ORI_O24_RESPONSE - Group struct
type ORI_O24_RESPONSE struct {
	Patient ORI_O24_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORI_O24_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORI_O24_RESPONSE_PATIENT - Group struct
type ORI_O24_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORI_O24_RESPONSE_ORDER - Group struct
type ORI_O24_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []ORI_O24_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ImagingProcedureControlSegment []IPC `hl7:"TAG=IPC"`
}

// ORI_O24_RESPONSE_ORDER_TIMING - Group struct
type ORI_O24_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORI_O24 - Imaging Order Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ORI_O24
type ORI_O24 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORI_O24_RESPONSE `hl7:"GROUP;ATR=optional"`
}

