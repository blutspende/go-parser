package hl7v28

// OMD_O03_PATIENT - Group struct
type OMD_O03_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient_visit OMD_O03_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OMD_O03_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OMD_O03_PATIENT_PATIENT_VISIT - Group struct
type OMD_O03_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMD_O03_PATIENT_INSURANCE - Group struct
type OMD_O03_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OMD_O03_ORDER_DIET - Group struct
type OMD_O03_ORDER_DIET struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing_diet []OMD_O03_ORDER_DIET_TIMING_DIET `hl7:"GROUP;ATR=optional"`
	Diet OMD_O03_ORDER_DIET_DIET `hl7:"GROUP;ATR=optional"`
}

// OMD_O03_ORDER_DIET_TIMING_DIET - Group struct
type OMD_O03_ORDER_DIET_TIMING_DIET struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OMD_O03_ORDER_DIET_DIET - Group struct
type OMD_O03_ORDER_DIET_DIET struct {
	DietaryOrdersSupplementsAndPreferences []ODS `hl7:"TAG=ODS"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []OMD_O03_ORDER_DIET_DIET_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// OMD_O03_ORDER_DIET_DIET_OBSERVATION - Group struct
type OMD_O03_ORDER_DIET_DIET_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMD_O03_ORDER_TRAY - Group struct
type OMD_O03_ORDER_TRAY struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing_tray []OMD_O03_ORDER_TRAY_TIMING_TRAY `hl7:"GROUP;ATR=optional"`
	DietTrayInstructions []ODT `hl7:"TAG=ODT"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMD_O03_ORDER_TRAY_TIMING_TRAY - Group struct
type OMD_O03_ORDER_TRAY_TIMING_TRAY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OMD_O03 - Dietary Order
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/OMD_O03
type OMD_O03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OMD_O03_PATIENT `hl7:"GROUP;ATR=optional"`
	Order_diet []OMD_O03_ORDER_DIET `hl7:"GROUP"`
	Order_tray []OMD_O03_ORDER_TRAY `hl7:"GROUP;ATR=optional"`
}

