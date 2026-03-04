package hl7v23

// ORF_W02_QUERY_RESPONSE - Group struct
type ORF_W02_QUERY_RESPONSE struct {
	Patient ORF_W02_QUERY_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORF_W02_QUERY_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORF_W02_QUERY_RESPONSE_PATIENT - Group struct
type ORF_W02_QUERY_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_W02_QUERY_RESPONSE_ORDER - Group struct
type ORF_W02_QUERY_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequestSegment OBR `hl7:"TAG=OBR"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []ORF_W02_QUERY_RESPONSE_ORDER_OBSERVATION `hl7:"GROUP"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORF_W02_QUERY_RESPONSE_ORDER_OBSERVATION - Group struct
type ORF_W02_QUERY_RESPONSE_ORDER_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_W02 - Waveform result, response to query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ORF_W02
type ORF_W02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	QueryResponse []ORF_W02_QUERY_RESPONSE `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

