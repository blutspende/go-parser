package hl7v26

// OPL_O37_GUARANTOR - Group struct
type OPL_O37_GUARANTOR struct {
	Guarantor GT1 `hl7:"TAG=GT1"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OPL_O37_ORDER - Group struct
type OPL_O37_ORDER struct {
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1"`
	Patient OPL_O37_ORDER_PATIENT `hl7:"GROUP;ATR=optional"`
	Specimen []OPL_O37_ORDER_SPECIMEN `hl7:"GROUP"`
	PriorResult OPL_O37_ORDER_PRIOR_RESULT `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OPL_O37_ORDER_PATIENT - Group struct
type OPL_O37_ORDER_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	Insurance []OPL_O37_ORDER_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OPL_O37_ORDER_PATIENT_INSURANCE - Group struct
type OPL_O37_ORDER_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OPL_O37_ORDER_SPECIMEN - Group struct
type OPL_O37_ORDER_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	Container []OPL_O37_ORDER_SPECIMEN_CONTAINER `hl7:"GROUP;ATR=optional"`
	ObservationRequest []OPL_O37_ORDER_SPECIMEN_OBSERVATION_REQUEST `hl7:"GROUP"`
}

// OPL_O37_ORDER_SPECIMEN_CONTAINER - Group struct
type OPL_O37_ORDER_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// OPL_O37_ORDER_SPECIMEN_OBSERVATION_REQUEST - Group struct
type OPL_O37_ORDER_SPECIMEN_OBSERVATION_REQUEST struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	Timing []OPL_O37_ORDER_SPECIMEN_OBSERVATION_REQUEST_TIMING `hl7:"GROUP;ATR=optional"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// OPL_O37_ORDER_SPECIMEN_OBSERVATION_REQUEST_TIMING - Group struct
type OPL_O37_ORDER_SPECIMEN_OBSERVATION_REQUEST_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OPL_O37_ORDER_PRIOR_RESULT - Group struct
type OPL_O37_ORDER_PRIOR_RESULT struct {
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1"`
	PatientPrior OPL_O37_ORDER_PRIOR_RESULT_PATIENT_PRIOR `hl7:"GROUP;ATR=optional"`
	PatientVisitPrior OPL_O37_ORDER_PRIOR_RESULT_PATIENT_VISIT_PRIOR `hl7:"GROUP;ATR=optional"`
	PatientAllergyInformation AL1 `hl7:"TAG=AL1;ATR=optional"`
	OrderPrior []OPL_O37_ORDER_PRIOR_RESULT_ORDER_PRIOR `hl7:"GROUP"`
}

// OPL_O37_ORDER_PRIOR_RESULT_PATIENT_PRIOR - Group struct
type OPL_O37_ORDER_PRIOR_RESULT_PATIENT_PRIOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
}

// OPL_O37_ORDER_PRIOR_RESULT_PATIENT_VISIT_PRIOR - Group struct
type OPL_O37_ORDER_PRIOR_RESULT_PATIENT_VISIT_PRIOR struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OPL_O37_ORDER_PRIOR_RESULT_ORDER_PRIOR - Group struct
type OPL_O37_ORDER_PRIOR_RESULT_ORDER_PRIOR struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	Timing OPL_O37_ORDER_PRIOR_RESULT_ORDER_PRIOR_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX"`
}

// OPL_O37_ORDER_PRIOR_RESULT_ORDER_PRIOR_TIMING - Group struct
type OPL_O37_ORDER_PRIOR_RESULT_ORDER_PRIOR_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OPL_O37 - Population/Location-Based Laboratory Order
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/OPL_O37
type OPL_O37 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL"`
	Guarantor OPL_O37_GUARANTOR `hl7:"GROUP;ATR=optional"`
	Order []OPL_O37_ORDER `hl7:"GROUP"`
}

