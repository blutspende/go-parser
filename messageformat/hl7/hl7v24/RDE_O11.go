package hl7v24

// RDE_O11_PATIENT - Group struct
type RDE_O11_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit RDE_O11_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []RDE_O11_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// RDE_O11_PATIENT_PATIENT_VISIT - Group struct
type RDE_O11_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RDE_O11_PATIENT_INSURANCE - Group struct
type RDE_O11_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RDE_O11_ORDER - Group struct
type RDE_O11_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetail RDE_O11_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RDE_O11_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// RDE_O11_ORDER_ORDER_DETAIL - Group struct
type RDE_O11_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Component RDE_O11_ORDER_ORDER_DETAIL_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDE_O11_ORDER_ORDER_DETAIL_COMPONENT - Group struct
type RDE_O11_ORDER_ORDER_DETAIL_COMPONENT struct {
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O11_ORDER_OBSERVATION - Group struct
type RDE_O11_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O11 - Pharmacy/treatment encoded order
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RDE_O11
type RDE_O11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDE_O11_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDE_O11_ORDER `hl7:"GROUP"`
}

