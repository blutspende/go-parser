package hl7v24

// ORF_R04_RESPONSE - Group struct
type ORF_R04_RESPONSE struct {
	Patient ORF_R04_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORF_R04_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORF_R04_RESPONSE_PATIENT - Group struct
type ORF_R04_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_R04_RESPONSE_ORDER - Group struct
type ORF_R04_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Observation []ORF_R04_RESPONSE_ORDER_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORF_R04_RESPONSE_ORDER_OBSERVATION - Group struct
type ORF_R04_RESPONSE_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_R04 - Response to query; transmission of requested observation
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORF_R04
type ORF_R04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Response []ORF_R04_RESPONSE `hl7:"GROUP"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

