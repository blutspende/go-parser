package hl7v231

// PEX_P08_VISIT - Group struct
type PEX_P08_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PEX_P08_EXPERIENCE - Group struct
type PEX_P08_EXPERIENCE struct {
	ProductExperienceSenderSegment PES `hl7:"TAG=PES"`
	PexObservation []PEX_P08_EXPERIENCE_PEX_OBSERVATION `hl7:"GROUP"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION struct {
	ProductExperienceObservationSegment PEO `hl7:"TAG=PEO"`
	PexCause []PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE `hl7:"GROUP"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE struct {
	PossibleCausalRelationshipSegment PCR `hl7:"TAG=PCR"`
	RxOrder PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ORDER `hl7:"GROUP;ATR=optional"`
	RxAdministration []PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ADMINISTRATION `hl7:"GROUP;ATR=optional"`
	ProblemDetail []PRB `hl7:"TAG=PRB;ATR=optional"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	AssociatedPerson PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON `hl7:"GROUP;ATR=optional"`
	Study []PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_STUDY `hl7:"GROUP;ATR=optional"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ORDER - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ORDER struct {
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ADMINISTRATION - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ADMINISTRATION struct {
	PharmacyTreatmentAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON struct {
	NextOfKinAssociatedPartiesSegment NK1 `hl7:"TAG=NK1"`
	AssociatedRxOrder PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ORDER `hl7:"GROUP;ATR=optional"`
	AssociatedRxAdmin []PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ADMIN `hl7:"GROUP;ATR=optional"`
	ProblemDetail []PRB `hl7:"TAG=PRB;ATR=optional"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ORDER - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ORDER struct {
	PharmacyTreatmentEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyTreatmentRouteSegment []RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ADMIN - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ADMIN struct {
	PharmacyTreatmentAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_STUDY - Group struct
type PEX_P08_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_STUDY struct {
	ClinicalStudyRegistrationSegment CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhaseSegment []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// PEX_P08 - Unsolicited update individual product experience report
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/PEX_P08
type PEX_P08 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit PEX_P08_VISIT `hl7:"GROUP;ATR=optional"`
	Experience []PEX_P08_EXPERIENCE `hl7:"GROUP"`
}

