package hl7v231

// ORS_O02_RESPONSE - Group struct
type ORS_O02_RESPONSE struct {
	Patient ORS_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORS_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORS_O02_RESPONSE_PATIENT - Group struct
type ORS_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORS_O02_RESPONSE_ORDER - Group struct
type ORS_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORS_O02 - Stock requisition - order acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORS_O02
type ORS_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORS_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

