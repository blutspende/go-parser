package hl7v25

// RSP_Z90_QUERY_RESPONSE - Group struct
type RSP_Z90_QUERY_RESPONSE struct {
	Patient RSP_Z90_QUERY_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	CommonOrder []RSP_Z90_QUERY_RESPONSE_COMMON_ORDER `hl7:"GROUP"`
	Specimen []RSP_Z90_QUERY_RESPONSE_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// RSP_Z90_QUERY_RESPONSE_PATIENT - Group struct
type RSP_Z90_QUERY_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit RSP_Z90_QUERY_RESPONSE_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RSP_Z90_QUERY_RESPONSE_PATIENT_VISIT - Group struct
type RSP_Z90_QUERY_RESPONSE_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RSP_Z90_QUERY_RESPONSE_COMMON_ORDER - Group struct
type RSP_Z90_QUERY_RESPONSE_COMMON_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RSP_Z90_QUERY_RESPONSE_COMMON_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Observation []RSP_Z90_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION `hl7:"GROUP"`
}

// RSP_Z90_QUERY_RESPONSE_COMMON_ORDER_TIMING - Group struct
type RSP_Z90_QUERY_RESPONSE_COMMON_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RSP_Z90_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION - Group struct
type RSP_Z90_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_Z90_QUERY_RESPONSE_SPECIMEN - Group struct
type RSP_Z90_QUERY_RESPONSE_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// RSP_Z90 - Lab Results History Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/RSP_Z90
type RSP_Z90 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	QueryResponse []RSP_Z90_QUERY_RESPONSE `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC"`
}

