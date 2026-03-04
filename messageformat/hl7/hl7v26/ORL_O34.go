package hl7v26

// ORL_O34_RESPONSE - Group struct
type ORL_O34_RESPONSE struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	Specimen []ORL_O34_RESPONSE_SPECIMEN `hl7:"GROUP"`
}

// ORL_O34_RESPONSE_SPECIMEN - Group struct
type ORL_O34_RESPONSE_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	SpecimenContainerDetail []SAC `hl7:"TAG=SAC;ATR=optional"`
	Order []ORL_O34_RESPONSE_SPECIMEN_ORDER `hl7:"GROUP;ATR=optional"`
}

// ORL_O34_RESPONSE_SPECIMEN_ORDER - Group struct
type ORL_O34_RESPONSE_SPECIMEN_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []ORL_O34_RESPONSE_SPECIMEN_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationRequest ORL_O34_RESPONSE_SPECIMEN_ORDER_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
}

// ORL_O34_RESPONSE_SPECIMEN_ORDER_TIMING - Group struct
type ORL_O34_RESPONSE_SPECIMEN_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORL_O34_RESPONSE_SPECIMEN_ORDER_OBSERVATION_REQUEST - Group struct
type ORL_O34_RESPONSE_SPECIMEN_ORDER_OBSERVATION_REQUEST struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// ORL_O34 - Multiple Laboratory Orders Related to Single Specimen Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ORL_O34
type ORL_O34 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORL_O34_RESPONSE `hl7:"GROUP;ATR=optional"`
}

