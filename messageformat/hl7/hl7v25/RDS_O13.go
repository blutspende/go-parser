package hl7v25

// RDS_O13_PATIENT - Group struct
type RDS_O13_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit RDS_O13_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RDS_O13_PATIENT_PATIENT_VISIT - Group struct
type RDS_O13_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RDS_O13_ORDER - Group struct
type RDS_O13_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RDS_O13_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	OrderDetail RDS_O13_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RDS_O13_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentDispense RXD `hl7:"TAG=RXD"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RDS_O13_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
}

// RDS_O13_ORDER_TIMING - Group struct
type RDS_O13_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RDS_O13_ORDER_ORDER_DETAIL - Group struct
type RDS_O13_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RDS_O13_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RDS_O13_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RDS_O13_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndComments []NTE `hl7:"TAG=NTE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Component []RDS_O13_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDS_O13_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT - Group struct
type RDS_O13_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENT struct {
	PharmacyTreatmentComponentOrder RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDS_O13_ORDER_ENCODING - Group struct
type RDS_O13_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	TimingEncoded []RDS_O13_ORDER_ENCODING_TIMING_ENCODED `hl7:"GROUP"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RDS_O13_ORDER_ENCODING_TIMING_ENCODED - Group struct
type RDS_O13_ORDER_ENCODING_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RDS_O13_ORDER_OBSERVATION - Group struct
type RDS_O13_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDS_O13 - Pharmacy/Treatment Dispense
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/RDS_O13
type RDS_O13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDS_O13_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDS_O13_ORDER `hl7:"GROUP"`
}

