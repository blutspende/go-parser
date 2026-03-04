package hl7v24

// ORS_O06_Response - Group struct
type ORS_O06_Response struct {
	Patient ORS_O06_Response_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORS_O06_Response_ORDER `hl7:"GROUP"`
}

// ORS_O06_Response_PATIENT - Group struct
type ORS_O06_Response_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORS_O06_Response_ORDER - Group struct
type ORS_O06_Response_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	RequisitionDetail RQD `hl7:"TAG=RQD"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORS_O06 - Stock requisition acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORS_O06
type ORS_O06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORS_O06_Response `hl7:"GROUP;ATR=optional"`
}

