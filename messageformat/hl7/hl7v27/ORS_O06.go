package hl7v27

// ORS_O06_RESPONSE - Group struct
type ORS_O06_RESPONSE struct {
	Patient ORS_O06_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORS_O06_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORS_O06_RESPONSE_PATIENT - Group struct
type ORS_O06_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORS_O06_RESPONSE_ORDER - Group struct
type ORS_O06_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []ORS_O06_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORS_O06_RESPONSE_ORDER_TIMING - Group struct
type ORS_O06_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORS_O06 - Stock requisition acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/ORS_O06
type ORS_O06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORS_O06_RESPONSE `hl7:"GROUP;ATR=optional"`
}

