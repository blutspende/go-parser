// Standard implementation for the lis2a2 protocol according to standard in every detail,
// will work for most with some tinkering....
package hl7v23

type Observation struct {
	Observation                 OBX   `hl7:"TAG=OBX" json:"observation,omitempty"`
	ObservationNotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional" json:"observationNotesAndComments,omitempty"`
}

type OrderDetailSegment struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional" json:"observationRequestSegment,omitempty"`
	RequisitionDetail         RQD `hl7:"TAG=RQD;ATR=optional" json:"requisitionDetail,omitempty"`
	RequisitionDetail1        RQ1 `hl7:"TAG=RQ1;ATR=optional" json:"requisitionDetail1,omitempty"`
	PharmacyPrescription      RXO `hl7:"TAG=RQ1;ATR=optional" json:"pharmacyPrescription,omitempty"`
	DietaryOrders             ODS `hl7:"TAG=ODS;ATR=optional" json:"dietaryOrders,omitempty"`
	DietTrayInstructions      ODT `hl7:"TAG=ODT;ATR=optional" json:"dietTrayInstructions,omitempty"`
}

type OrderDetail struct {
	OrderDetailSegment OrderDetailSegment `hl7:"GROUP" json:"orderDetailSegment,omitempty"`
	NotesAndComments   []NTE              `hl7:"TAG=NTE;ATR=optional" json:"notesAndComments,omitempty"`
	Diagnosis          []DG1              `hl7:"TAG=DG1;ATR=optional" json:"diagnosis,omitempty"`
	Observation        []Observation      `hl7:"GROUP;ATR=optional" json:"observation,omitempty"`
}

type Order struct {
	CommonOrderSegment          ORC         `hl7:"TAG=ORC" json:"commonOrderSegment,omitempty"`
	OrderDetail                 OrderDetail `hl7:"GROUP" json:"orderDetail,omitempty"`
	ClinicalTrialIdentification []CTI       `hl7:"TAG=CTI;ATR=optional" json:"clinicalTrialIdentification,omitempty"`
	Billing                     BLG         `hl7:"TAG=BLG;ATR=optional" json:"billing,omitempty"`
}

type Insurance struct {
	Insurance             IN1 `hl7:"TAG=IN1" json:"insurance,omitempty"`
	AdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional" json:"additionalInformation,omitempty"`
	Certification         IN3 `hl7:"TAG=IN3;ATR=optional" json:"certification,omitempty"`
}

type PatientVisit struct {
	PatientVisit          PV1 `hl7:"TAG=PV1" json:"patientVisit,omitempty"`
	AdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional" json:"additionalInformation,omitempty"`
}

type Patient struct {
	PatientIdentification     PID          `hl7:"TAG=PID;ATR=optional" json:"patientIdentification,omitempty"`
	PatientDemographics       PD1          `hl7:"TAG=PD1;ATR=optional" json:"patientDemographics,omitempty"`
	NotesAndComments          []NTE        `hl7:"TAG=NTE;ATR=optional" json:"notesAndComments,omitempty"`
	PatientVisit              PatientVisit `hl7:"GROUP;ATR=optional" json:"patientVisit,omitempty"`
	Insurance                 []Insurance  `hl7:"GROUP" json:"insurance,omitempty"`
	Guarantor                 GT1          `hl7:"TAG=GT1;ATR=optional" json:"guarantor,omitempty"`
	PatientAllergyInformation []AL1        `hl7:"TAG=AL1;ATR=optional" json:"patientAllergyInformation,omitempty"`
}

// ORM_O01 - Order message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ORM_O01
type ORM_O01 struct {
	MSH              MSH     `hl7:"TAG=MSH" json:"msh,omitempty"`
	NotesAndComments []NTE   `hl7:"TAG=NTE;ATR=optional" json:"notesAndComments,omitempty"`
	Patient          Patient `hl7:"GROUP" json:"patient,omitempty"`
	Order            []Order `hl7:"GROUP" json:"order,omitempty"`
}
