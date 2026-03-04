package hl7v27

// OPR_O38_RESPONSE - Group struct
type OPR_O38_RESPONSE struct {
	Order []OPR_O38_RESPONSE_ORDER `hl7:"GROUP"`
}

// OPR_O38_RESPONSE_ORDER - Group struct
type OPR_O38_RESPONSE_ORDER struct {
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Specimen []OPR_O38_RESPONSE_ORDER_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// OPR_O38_RESPONSE_ORDER_SPECIMEN - Group struct
type OPR_O38_RESPONSE_ORDER_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	Specimen_observation []OPR_O38_RESPONSE_ORDER_SPECIMEN_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
	SpecimenContainerDetail []SAC `hl7:"TAG=SAC;ATR=optional"`
	Observation_request []OPR_O38_RESPONSE_ORDER_SPECIMEN_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
	Timing []OPR_O38_RESPONSE_ORDER_SPECIMEN_TIMING `hl7:"GROUP;ATR=optional"`
}

// OPR_O38_RESPONSE_ORDER_SPECIMEN_SPECIMEN_OBSERVATION - Group struct
type OPR_O38_RESPONSE_ORDER_SPECIMEN_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OPR_O38_RESPONSE_ORDER_SPECIMEN_OBSERVATION_REQUEST - Group struct
type OPR_O38_RESPONSE_ORDER_SPECIMEN_OBSERVATION_REQUEST struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OPR_O38_RESPONSE_ORDER_SPECIMEN_TIMING - Group struct
type OPR_O38_RESPONSE_ORDER_SPECIMEN_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OPR_O38 - Population/Location-Based Laboratory Order Acknowledgment Message
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/OPR_O38
type OPR_O38 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response OPR_O38_RESPONSE `hl7:"GROUP;ATR=optional"`
}

