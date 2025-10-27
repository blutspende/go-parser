package hl7v23

type OrderObservation struct {
	CommonOrder                 ORC           `hl7:"TAG=ORC;ATR=optional" json:"commonOrder,omitempty"`
	ObservationRequest          OBR           `hl7:"TAG=OBR" json:"observationRequest,omitempty"`
	NotesAndComments            []NTE         `hl7:"TAG=NTE;ATR=optional" json:"notesAndComments,omitempty"`
	Observation                 []Observation `hl7:"GROUP" json:"observation,omitempty"`
	ClinicalTrialIdentification []CTI         `hl7:"TAG=CTI;ATR=optional" json:"clinicalTrialIdentification,omitempty"`
}

type PatientResult struct {
	Patient          Patient            `hl7:"GROUP;ATR=optional" json:"patient,omitempty"`
	OrderObservation []OrderObservation `hl7:"GROUP" json:"orderObservation,omitempty"`
}

// ORU_R01 - Unsolicited transmission of an observation message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ORU_R01
type ORU_R01 struct {
	MSH                 MSH             `hl7:"TAG=MSH" json:"msh,omitempty"`
	PatientResult       []PatientResult `hl7:"GROUP" json:"patientResult,omitempty"`
	ContinuationPointer DSC             `hl7:"TAG=DSC;ATR= optional" json:"continuationPointer,omitempty"`
}
