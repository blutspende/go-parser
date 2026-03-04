package hl7v24

// ORD_O04_RESPONSE - Group struct
type ORD_O04_RESPONSE struct {
	Patient ORD_O04_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderDiet []ORD_O04_RESPONSE_ORDER_DIET `hl7:"GROUP"`
	OrderTray []ORD_O04_RESPONSE_ORDER_TRAY `hl7:"GROUP;ATR=optional"`
}

// ORD_O04_RESPONSE_PATIENT - Group struct
type ORD_O04_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORD_O04_RESPONSE_ORDER_DIET - Group struct
type ORD_O04_RESPONSE_ORDER_DIET struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	DietaryOrdersSupplementsAndPreferences []ODS `hl7:"TAG=ODS;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORD_O04_RESPONSE_ORDER_TRAY - Group struct
type ORD_O04_RESPONSE_ORDER_TRAY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	DietTrayInstructions []ODT `hl7:"TAG=ODT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORD_O04 - Dietary order acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORD_O04
type ORD_O04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORD_O04_RESPONSE `hl7:"GROUP;ATR=optional"`
}

