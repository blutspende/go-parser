package hl7v231

// OMD_O01_PATIENT - Group struct
type OMD_O01_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit OMD_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OMD_O01_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	GuarantorSegment GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OMD_O01_PATIENT_PATIENT_VISIT - Group struct
type OMD_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OMD_O01_PATIENT_INSURANCE - Group struct
type OMD_O01_PATIENT_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OMD_O01_ORDER_DIET - Group struct
type OMD_O01_ORDER_DIET struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	Diet OMD_O01_ORDER_DIET_DIET `hl7:"GROUP;ATR=optional"`
}

// OMD_O01_ORDER_DIET_DIET - Group struct
type OMD_O01_ORDER_DIET_DIET struct {
	DietaryOrdersSupplementsAndPreferencesSegment []ODS `hl7:"TAG=ODS"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []OMD_O01_ORDER_DIET_DIET_OBSERVATION `hl7:"GROUP"`
}

// OMD_O01_ORDER_DIET_DIET_OBSERVATION - Group struct
type OMD_O01_ORDER_DIET_DIET_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMD_O01_ORDER_TRAY - Group struct
type OMD_O01_ORDER_TRAY struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	DietTrayInstructionsSegment []ODT `hl7:"TAG=ODT"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMD_O01 - Dietary order
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/OMD_O01
type OMD_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OMD_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	OrderDiet []OMD_O01_ORDER_DIET `hl7:"GROUP"`
	OrderTray []OMD_O01_ORDER_TRAY `hl7:"GROUP;ATR=optional"`
}

