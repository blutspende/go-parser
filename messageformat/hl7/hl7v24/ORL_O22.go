package hl7v24

// ORL_O22_RESPONSE - Group struct
type ORL_O22_RESPONSE struct {
	Patient ORL_O22_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT - Group struct
type ORL_O22_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	GeneralOrder []ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER `hl7:"GROUP"`
}

// ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER - Group struct
type ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER struct {
	Container ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_CONTAINER `hl7:"GROUP;ATR=optional"`
	Order []ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_ORDER `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_CONTAINER - Group struct
type ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_CONTAINER struct {
	SpecimenAndContainerDetail SAC `hl7:"TAG=SAC"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_ORDER - Group struct
type ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_ORDER_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_ORDER_OBSERVATION_REQUEST - Group struct
type ORL_O22_RESPONSE_PATIENT_GENERAL_ORDER_ORDER_OBSERVATION_REQUEST struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	SpecimenAndContainerDetail []SAC `hl7:"TAG=SAC;ATR=optional"`
}

// ORL_O22 - General laboratory order response message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORL_O22
type ORL_O22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORL_O22_RESPONSE `hl7:"GROUP;ATR=optional"`
}

