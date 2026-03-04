package hl7v251

// ORL_O22_RESPONSE - Group struct
type ORL_O22_RESPONSE struct {
	Patient ORL_O22_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT - Group struct
type ORL_O22_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	Order []ORL_O22_RESPONSE_PATIENT_ORDER `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT_ORDER - Group struct
type ORL_O22_RESPONSE_PATIENT_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []ORL_O22_RESPONSE_PATIENT_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationRequest ORL_O22_RESPONSE_PATIENT_ORDER_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT_ORDER_TIMING - Group struct
type ORL_O22_RESPONSE_PATIENT_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT_ORDER_OBSERVATION_REQUEST - Group struct
type ORL_O22_RESPONSE_PATIENT_ORDER_OBSERVATION_REQUEST struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	Specimen []ORL_O22_RESPONSE_PATIENT_ORDER_OBSERVATION_REQUEST_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT_ORDER_OBSERVATION_REQUEST_SPECIMEN - Group struct
type ORL_O22_RESPONSE_PATIENT_ORDER_OBSERVATION_REQUEST_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	SpecimenContainerDetail []SAC `hl7:"TAG=SAC;ATR=optional"`
}

// ORL_O22 - General Laboratory Order Response
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/ORL_O22
type ORL_O22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORL_O22_RESPONSE `hl7:"GROUP;ATR=optional"`
}

