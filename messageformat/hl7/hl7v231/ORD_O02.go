package hl7v231

// ORD_O02_RESPONSE - Group struct
type ORD_O02_RESPONSE struct {
	Patient ORD_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderDiet []ORD_O02_RESPONSE_ORDER_DIET `hl7:"GROUP"`
	OrderTray []ORD_O02_RESPONSE_ORDER_TRAY `hl7:"GROUP;ATR=optional"`
}

// ORD_O02_RESPONSE_PATIENT - Group struct
type ORD_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORD_O02_RESPONSE_ORDER_DIET - Group struct
type ORD_O02_RESPONSE_ORDER_DIET struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	DietaryOrdersSupplementsAndPreferencesSegment []ODS `hl7:"TAG=ODS;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORD_O02_RESPONSE_ORDER_TRAY - Group struct
type ORD_O02_RESPONSE_ORDER_TRAY struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	DietTrayInstructionsSegment []ODT `hl7:"TAG=ODT;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORD_O02 - Dietary order acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORD_O02
type ORD_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORD_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

