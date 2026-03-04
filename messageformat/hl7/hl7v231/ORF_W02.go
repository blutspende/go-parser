package hl7v231

// ORF_W02_QUERY_RESPONSE - Group struct
type ORF_W02_QUERY_RESPONSE struct {
	Patient ORF_W02_QUERY_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORF_W02_QUERY_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORF_W02_QUERY_RESPONSE_PATIENT - Group struct
type ORF_W02_QUERY_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_W02_QUERY_RESPONSE_ORDER - Group struct
type ORF_W02_QUERY_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []ORF_W02_QUERY_RESPONSE_ORDER_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORF_W02_QUERY_RESPONSE_ORDER_OBSERVATION - Group struct
type ORF_W02_QUERY_RESPONSE_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_W02 - Waveform result, response to query
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORF_W02
type ORF_W02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	QueryResponse []ORF_W02_QUERY_RESPONSE `hl7:"GROUP"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

