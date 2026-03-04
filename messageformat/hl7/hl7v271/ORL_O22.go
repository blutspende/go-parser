package hl7v271

// ORL_O22_RESPONSE - Group struct
type ORL_O22_RESPONSE struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Order []ORL_O22_RESPONSE_ORDER `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_ORDER - Group struct
type ORL_O22_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing []ORL_O22_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Observation_request ORL_O22_RESPONSE_ORDER_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_ORDER_TIMING - Group struct
type ORL_O22_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORL_O22_RESPONSE_ORDER_OBSERVATION_REQUEST - Group struct
type ORL_O22_RESPONSE_ORDER_OBSERVATION_REQUEST struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Specimen []ORL_O22_RESPONSE_ORDER_OBSERVATION_REQUEST_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// ORL_O22_RESPONSE_ORDER_OBSERVATION_REQUEST_SPECIMEN - Group struct
type ORL_O22_RESPONSE_ORDER_OBSERVATION_REQUEST_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	SpecimenContainerDetail []SAC `hl7:"TAG=SAC;ATR=optional"`
}

// ORL_O22 - General laboratory order response message to any OML
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/ORL_O22
type ORL_O22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORL_O22_RESPONSE `hl7:"GROUP;ATR=optional"`
}

