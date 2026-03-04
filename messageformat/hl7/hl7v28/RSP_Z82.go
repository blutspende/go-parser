package hl7v28

// RSP_Z82_QUERY_RESPONSE - Group struct
type RSP_Z82_QUERY_RESPONSE struct {
	Patient RSP_Z82_QUERY_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Common_order []RSP_Z82_QUERY_RESPONSE_COMMON_ORDER `hl7:"GROUP"`
}

// RSP_Z82_QUERY_RESPONSE_PATIENT - Group struct
type RSP_Z82_QUERY_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit RSP_Z82_QUERY_RESPONSE_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RSP_Z82_QUERY_RESPONSE_PATIENT_VISIT - Group struct
type RSP_Z82_QUERY_RESPONSE_PATIENT_VISIT struct {
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RSP_Z82_QUERY_RESPONSE_COMMON_ORDER - Group struct
type RSP_Z82_QUERY_RESPONSE_COMMON_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Order_detail RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoded_order RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentDispense RXD `hl7:"TAG=RXD"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION `hl7:"GROUP"`
}

// RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_TIMING - Group struct
type RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL - Group struct
type RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Treatment RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL_TREATMENT `hl7:"GROUP;ATR=optional"`
}

// RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL_TREATMENT - Group struct
type RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL_TREATMENT struct {
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER - Group struct
type RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	Timing_encoded []RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER_TIMING_ENCODED `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER_TIMING_ENCODED - Group struct
type RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION - Group struct
type RSP_Z82_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_Z82 - Dispense History (Response)
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RSP_Z82
type RSP_Z82 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	Query_response []RSP_Z82_QUERY_RESPONSE `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

