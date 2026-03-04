package hl7v25

// RSP_Z86_QUERY_RESPONSE - Group struct
type RSP_Z86_QUERY_RESPONSE struct {
	Patient RSP_Z86_QUERY_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	CommonOrder []RSP_Z86_QUERY_RESPONSE_COMMON_ORDER `hl7:"GROUP"`
}

// RSP_Z86_QUERY_RESPONSE_PATIENT - Group struct
type RSP_Z86_QUERY_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	OrderDetail RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	EncodedOrder RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER `hl7:"GROUP;ATR=optional"`
	Dispense RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_DISPENSE `hl7:"GROUP;ATR=optional"`
	Give RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_GIVE `hl7:"GROUP;ATR=optional"`
	Administration RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ADMINISTRATION `hl7:"GROUP;ATR=optional"`
	Observation []RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION `hl7:"GROUP"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_TIMING - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	TimingEncoded []RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER_TIMING_ENCODED `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER_TIMING_ENCODED - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ENCODED_ORDER_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_DISPENSE - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_DISPENSE struct {
	PharmacyTreatmentDispense RXD `hl7:"TAG=RXD"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_GIVE - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_GIVE struct {
	PharmacyTreatmentGive RXG `hl7:"TAG=RXG"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ADMINISTRATION - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_ADMINISTRATION struct {
	PharmacyTreatmentAdministration RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION - Group struct
type RSP_Z86_QUERY_RESPONSE_COMMON_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_Z86 - Pharmacy Information Comprehensive Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/RSP_Z86
type RSP_Z86 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	QueryResponse []RSP_Z86_QUERY_RESPONSE `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

