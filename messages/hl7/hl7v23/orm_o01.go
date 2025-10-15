// Standard implementation for the lis2a2 protocol according to standard in every detail,
// will work for most with some tinkering....
package hl7v23

type Observation struct {
	Observation                 OBX   `hl7:"OBX" json:"observation,omitempty"`
	ObservationNotesAndComments []NTE `hl7:"NTE,optional" json:"observationNotesAndComments,omitempty"`
}

type OrderDetailSegment struct {
	ObservationRequestSegment OBR `hl7:"OBR,optional" json:"observationRequestSegment,omitempty"`
	RequisitionDetail         RQD `hl7:"RQD,optional" json:"requisitionDetail,omitempty"`
	RequisitionDetail1        RQ1 `hl7:"RQ1,optional" json:"requisitionDetail1,omitempty"`
	PharmacyPrescription      RXO `hl7:"RQ1,optional" json:"pharmacyPrescription,omitempty"`
	DietaryOrders             ODS `hl7:"ODS,optional" json:"dietaryOrders,omitempty"`
	DietTrayInstructions      ODT `hl7:"ODT,optional" json:"dietTrayInstructions,omitempty"`
}

type OrderDetail struct {
	OrderDetailSegment OrderDetailSegment `json:"orderDetailSegment,omitempty"`
	NotesAndComments   []NTE              `hl7:"NTE,optional" json:"notesAndComments,omitempty"`
	Diagnosis          []DG1              `hl7:"DG1,optional" json:"diagnosis,omitempty"`
	Observation        []Observation      `hl7:"Observation,optional" json:"observation,omitempty"`
}

type Order struct {
	CommonOrderSegment          ORC         `hl7:"ORC" json:"commonOrderSegment,omitempty"`
	OrderDetail                 OrderDetail `json:"orderDetail,omitempty"`
	ClinicalTrialIdentification []CTI       `hl7:"CTI,optional" json:"clinicalTrialIdentification,omitempty"`
	Billing                     BLG         `hl7:"BLG,optional" json:"billing,omitempty"`
}

type Insurance struct {
	Insurance             IN1 `hl7:"IN1" json:"insurance,omitempty"`
	AdditionalInformation IN2 `hl7:"IN2,optional" json:"additionalInformation,omitempty"`
	Certification         IN3 `hl7:"IN3,optional" json:"certification,omitempty"`
}

type PatientVisit struct {
	PatientVisit          PV1 `hl7:"PV1" json:"patientVisit,omitempty"`
	AdditionalInformation PV2 `hl7:"PV2,optional" json:"additionalInformation,omitempty"`
}

type Patient struct {
	PatientIdentification     PID          `hl7:"PID,optional" json:"patientIdentification,omitempty"`
	PatientDemographics       PD1          `hl7:"PD1,optional" json:"patientDemographics,omitempty"`
	NotesAndComments          []NTE        `hl7:"NTE,optional" json:"notesAndComments,omitempty"`
	PatientVisit              PatientVisit `json:"patientVisit,omitempty"`
	Insurance                 []Insurance  `json:"insurance,omitempty"`
	Guarantor                 GT1          `hl7:"GT1,optional" json:"guarantor,omitempty"`
	PatientAllergyInformation []AL1        `hl7:"AL1,optional" json:"patientAllergyInformation,omitempty"`
}

// ORM_O01 - Order message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ORM_O01
type ORM_O01 struct {
	MSH              MSH     `hl7:"MSH" json:"msh,omitempty"`
	NotesAndComments []NTE   `hl7:"NTE,optional" json:"notesAndComments,omitempty"`
	Patient          Patient `json:"patient,omitempty"`
	Order            []Order `json:"order,omitempty"`
}
