package hl7v26

// ORU_R32_VISIT - Group struct
type ORU_R32_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// ORU_R32_TIMING_QTY - Group struct
type ORU_R32_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORU_R32_OBSERVATION - Group struct
type ORU_R32_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORU_R32 - Unsolicited Pre-Ordered Point-Of-Care Observation 
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ORU_R32
type ORU_R32 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	Visit ORU_R32_VISIT `hl7:"GROUP;ATR=optional"`
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	TimingQty []ORU_R32_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	Observation []ORU_R32_OBSERVATION `hl7:"GROUP"`
}

