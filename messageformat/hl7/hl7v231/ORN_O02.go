package hl7v231

// ORN_O02_RESPONSE - Group struct
type ORN_O02_RESPONSE struct {
	Patient ORN_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORN_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORN_O02_RESPONSE_PATIENT - Group struct
type ORN_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORN_O02_RESPONSE_ORDER - Group struct
type ORN_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORN_O02 - Non-stock requisition - general order acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORN_O02
type ORN_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORN_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

