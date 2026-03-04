package hl7v27

// OMP_O09_PATIENT - Group struct
type OMP_O09_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	Additional_demographics OMP_O09_PATIENT_ADDITIONAL_DEMOGRAPHICS `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient_visit OMP_O09_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OMP_O09_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OMP_O09_PATIENT_ADDITIONAL_DEMOGRAPHICS - Group struct
type OMP_O09_PATIENT_ADDITIONAL_DEMOGRAPHICS struct {
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMP_O09_PATIENT_PATIENT_VISIT - Group struct
type OMP_O09_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OMP_O09_PATIENT_INSURANCE - Group struct
type OMP_O09_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OMP_O09_ORDER - Group struct
type OMP_O09_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []OMP_O09_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Component []OMP_O09_ORDER_COMPONENT `hl7:"GROUP;ATR=optional"`
	Observation []OMP_O09_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OMP_O09_ORDER_TIMING - Group struct
type OMP_O09_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OMP_O09_ORDER_COMPONENT - Group struct
type OMP_O09_ORDER_COMPONENT struct {
	PharmacyTreatmentComponentOrder RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMP_O09_ORDER_OBSERVATION - Group struct
type OMP_O09_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMP_O09 - Pharmacy/treatment order
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/OMP_O09
type OMP_O09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OMP_O09_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OMP_O09_ORDER `hl7:"GROUP"`
}

