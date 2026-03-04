package hl7v251

// RGV_O15_PATIENT - Group struct
type RGV_O15_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit RGV_O15_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
}

// RGV_O15_PATIENT_PATIENT_VISIT - Group struct
type RGV_O15_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// RGV_O15_ORDER - Group struct
type RGV_O15_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []RGV_O15_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	OrderDetail RGV_O15_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Encoding RGV_O15_ORDER_ENCODING `hl7:"GROUP;ATR=optional"`
	Give []RGV_O15_ORDER_GIVE `hl7:"GROUP"`
}

// RGV_O15_ORDER_TIMING - Group struct
type RGV_O15_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RGV_O15_ORDER_ORDER_DETAIL - Group struct
type RGV_O15_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	OrderDetailSupplement RGV_O15_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT `hl7:"GROUP;ATR=optional"`
}

// RGV_O15_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT - Group struct
type RGV_O15_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT struct {
	NotesAndComments []NTE `hl7:"TAG=NTE"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Components []RGV_O15_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS `hl7:"GROUP;ATR=optional"`
}

// RGV_O15_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS - Group struct
type RGV_O15_ORDER_ORDER_DETAIL_ORDER_DETAIL_SUPPLEMENT_COMPONENTS struct {
	PharmacyTreatmentComponentOrder RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGV_O15_ORDER_ENCODING - Group struct
type RGV_O15_ORDER_ENCODING struct {
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	TimingEncoded []RGV_O15_ORDER_ENCODING_TIMING_ENCODED `hl7:"GROUP"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
}

// RGV_O15_ORDER_ENCODING_TIMING_ENCODED - Group struct
type RGV_O15_ORDER_ENCODING_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RGV_O15_ORDER_GIVE - Group struct
type RGV_O15_ORDER_GIVE struct {
	PharmacyTreatmentGive RXG `hl7:"TAG=RXG"`
	TimingGive []RGV_O15_ORDER_GIVE_TIMING_GIVE `hl7:"GROUP"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	Observation []RGV_O15_ORDER_GIVE_OBSERVATION `hl7:"GROUP"`
}

// RGV_O15_ORDER_GIVE_TIMING_GIVE - Group struct
type RGV_O15_ORDER_GIVE_TIMING_GIVE struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RGV_O15_ORDER_GIVE_OBSERVATION - Group struct
type RGV_O15_ORDER_GIVE_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RGV_O15 - Pharmacy/Treatment Give
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/RGV_O15
type RGV_O15 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RGV_O15_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RGV_O15_ORDER `hl7:"GROUP"`
}

