package hl7v28

// RSP_K31_RESPONSE - Group struct
type RSP_K31_RESPONSE struct {
	Patient RSP_K31_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RSP_K31_RESPONSE_ORDER `hl7:"GROUP"`
}

// RSP_K31_RESPONSE_PATIENT - Group struct
type RSP_K31_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	Additional_demographics RSP_K31_RESPONSE_PATIENT_ADDITIONAL_DEMOGRAPHICS `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Patient_visit RSP_K31_RESPONSE_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RSP_K31_RESPONSE_PATIENT_ADDITIONAL_DEMOGRAPHICS - Group struct
type RSP_K31_RESPONSE_PATIENT_ADDITIONAL_DEMOGRAPHICS struct {
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// RSP_K31_RESPONSE_PATIENT_PATIENT_VISIT - Group struct
type RSP_K31_RESPONSE_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// RSP_K31_RESPONSE_ORDER - Group struct
type RSP_K31_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RSP_K31_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Order_detail RSP_K31_RESPONSE_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RSP_K31_RESPONSE_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentDispense RXD `hl7:"TAG=RXD"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RSP_K31_RESPONSE_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// RSP_K31_RESPONSE_ORDER_TIMING - Group struct
type RSP_K31_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RSP_K31_RESPONSE_ORDER_ORDER_DETAIL - Group struct
type RSP_K31_RESPONSE_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Components []RSP_K31_RESPONSE_ORDER_ORDER_DETAIL_COMPONENTS `hl7:"GROUP;ATR=optional"`
}

// RSP_K31_RESPONSE_ORDER_ORDER_DETAIL_COMPONENTS - Group struct
type RSP_K31_RESPONSE_ORDER_ORDER_DETAIL_COMPONENTS struct {
	PharmacyTreatmentComponentOrder RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_K31_RESPONSE_ORDER_ENCODING - Group struct
type RSP_K31_RESPONSE_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Timing_encoded []RSP_K31_RESPONSE_ORDER_ENCODING_TIMING_ENCODED `hl7:"GROUP"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RSP_K31_RESPONSE_ORDER_ENCODING_TIMING_ENCODED - Group struct
type RSP_K31_RESPONSE_ORDER_ENCODING_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RSP_K31_RESPONSE_ORDER_OBSERVATION - Group struct
type RSP_K31_RESPONSE_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_K31 - Dispense History Response
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RSP_K31
type RSP_K31 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	Response []RSP_K31_RESPONSE `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

