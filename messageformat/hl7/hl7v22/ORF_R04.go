package hl7v22

// ORF_R04_QUERY_RESPONSE - Group struct
type ORF_R04_QUERY_RESPONSE struct {
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Order []ORF_R04_QUERY_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORF_R04_QUERY_RESPONSE_ORDER - Group struct
type ORF_R04_QUERY_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []ORF_R04_QUERY_RESPONSE_ORDER_OBSERVATION `hl7:"GROUP"`
}

// ORF_R04_QUERY_RESPONSE_ORDER_OBSERVATION - Group struct
type ORF_R04_QUERY_RESPONSE_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORF_R04 - Response to query; transmission of requested observation
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ORF_R04
type ORF_R04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Query_response []ORF_R04_QUERY_RESPONSE `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

