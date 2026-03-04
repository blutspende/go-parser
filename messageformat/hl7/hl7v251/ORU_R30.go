package hl7v251

// ORU_R30_VISIT - Group struct
type ORU_R30_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// ORU_R30_TIMING_QTY - Group struct
type ORU_R30_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORU_R30_OBSERVATION - Group struct
type ORU_R30_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORU_R30 - Unsolicited Point-Of-Care Observation Message Without Existing Order – Place An Order
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/ORU_R30
type ORU_R30 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	Visit ORU_R30_VISIT `hl7:"GROUP;ATR=optional"`
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	TimingQty []ORU_R30_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	Observation []ORU_R30_OBSERVATION `hl7:"GROUP"`
}

