package hl7v23

type OrderObservation struct {
	CommonOrder                 ORC           `hl7:"ORC,optional" json:"commonOrder,omitempty"`
	ObservationRequest          OBR           `hl7:"OBR" json:"observationRequest,omitempty"`
	NotesAndComments            []NTE         `hl7:"NTE,optional" json:"notesAndComments,omitempty"`
	Observation                 []Observation `hl7:"Observation" json:"observation,omitempty"`
	ClinicalTrialIdentification []CTI         `hl7:"CTI,optional" json:"clinicalTrialIdentification,omitempty"`
}

type PatientResult struct {
	Patient          Patient            `hl7:"Patient,optional" json:"patient,omitempty"`
	OrderObservation []OrderObservation `hl7:"OrderObservation" json:"orderObservation,omitempty"`
}

// ORU_R01 - Unsolicited transmission of an observation message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ORU_R01
type ORU_R01 struct {
	MSH                 MSH             `hl7:"MSH" json:"msh,omitempty"`
	PatientResult       []PatientResult `json:"patientResult,omitempty"`
	ContinuationPointer DSC             `hl7:"DSC, optional" json:"continuationPointer,omitempty"`
}
