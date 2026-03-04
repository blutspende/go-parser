package hl7v25

// ORF_W02_RESPONSE - Group struct
type ORF_W02_RESPONSE struct {
	Patient ORF_W02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORF_W02_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORF_W02_RESPONSE_PATIENT - Group struct
type ORF_W02_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_W02_RESPONSE_ORDER - Group struct
type ORF_W02_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Observation []ORF_W02_RESPONSE_ORDER_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORF_W02_RESPONSE_ORDER_OBSERVATION - Group struct
type ORF_W02_RESPONSE_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_W02 - Waveform result, response to query
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ORF_W02
type ORF_W02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Response []ORF_W02_RESPONSE `hl7:"GROUP"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

