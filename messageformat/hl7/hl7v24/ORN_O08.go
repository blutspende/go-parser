package hl7v24

// ORN_O08_RESPONSE - Group struct
type ORN_O08_RESPONSE struct {
	Patient ORN_O08_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORN_O08_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORN_O08_RESPONSE_PATIENT - Group struct
type ORN_O08_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORN_O08_RESPONSE_ORDER - Group struct
type ORN_O08_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	RequisitionDetail RQD `hl7:"TAG=RQD"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORN_O08 - Non-stock requisition acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORN_O08
type ORN_O08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORN_O08_RESPONSE `hl7:"GROUP;ATR=optional"`
}

