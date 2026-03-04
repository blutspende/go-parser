package hl7v251

// RAS_O17_PATIENT - Group struct
type RAS_O17_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit RAS_O17_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RAS_O17_PATIENT_PATIENT_VISIT - Group struct
type RAS_O17_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RAS_O17_ORDER - Group struct
type RAS_O17_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RAS_O17_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	OrderDetail RAS_O17_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RAS_O17_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	Administration []RAS_O17_ORDER_ADMINISTRATION `hl7:"GROUP"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// RAS_O17_ORDER_TIMING - Group struct
type RAS_O17_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RAS_O17_ORDER_ORDER_DETAIL - Group struct
type RAS_O17_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RAS_O17_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RAS_O17_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RAS_O17_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndComments []NTE `hl7:"TAG=NTE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Components []RAS_O17_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS `hl7:"GROUP;ATR=optional"`
}

// RAS_O17_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS - Group struct
type RAS_O17_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS struct {
	PharmacyTreatmentComponentOrder RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O17_ORDER_ENCODING - Group struct
type RAS_O17_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	TimingEncoded []RAS_O17_ORDER_ENCODING_TIMING_ENCODED `hl7:"GROUP"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RAS_O17_ORDER_ENCODING_TIMING_ENCODED - Group struct
type RAS_O17_ORDER_ENCODING_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RAS_O17_ORDER_ADMINISTRATION - Group struct
type RAS_O17_ORDER_ADMINISTRATION struct {
	PharmacyTreatmentAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
	Observation []RAS_O17_ORDER_ADMINISTRATION_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// RAS_O17_ORDER_ADMINISTRATION_OBSERVATION - Group struct
type RAS_O17_ORDER_ADMINISTRATION_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RAS_O17 - Pharmacy/Treatment Administration
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/RAS_O17
type RAS_O17 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RAS_O17_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RAS_O17_ORDER `hl7:"GROUP"`
}

