package hl7v23

// PEX_P07_VISIT - Group struct
type PEX_P07_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PEX_P07_EXPERIENCE - Group struct
type PEX_P07_EXPERIENCE struct {
	ProductExperienceSender PES `hl7:"TAG=PES"`
	PexObservation []PEX_P07_EXPERIENCE_PEX_OBSERVATION `hl7:"GROUP"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION struct {
	ProductExperienceObservation PEO `hl7:"TAG=PEO"`
	PexCause []PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE `hl7:"GROUP"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE struct {
	PossibleCausalRelationship PCR `hl7:"TAG=PCR"`
	RxOrder PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ORDER `hl7:"GROUP;ATR=optional"`
	RxAdministration []PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ADMINISTRATION `hl7:"GROUP;ATR=optional"`
	ProblemDetail []PRB `hl7:"TAG=PRB;ATR=optional"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	AssociatedPerson PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON `hl7:"GROUP;ATR=optional"`
	Study []PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_STUDY `hl7:"GROUP;ATR=optional"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ORDER - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ORDER struct {
	PharmacyEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ADMINISTRATION - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_RX_ADMINISTRATION struct {
	PharmacyAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON struct {
	NextOfKin NK1 `hl7:"TAG=NK1"`
	AssociatedRxOrder PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ORDER `hl7:"GROUP;ATR=optional"`
	AssociatedRxAdmin []PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ADMIN `hl7:"GROUP;ATR=optional"`
	ProblemDetail []PRB `hl7:"TAG=PRB;ATR=optional"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ORDER - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ORDER struct {
	PharmacyEncodedOrderSegment RXE `hl7:"TAG=RXE"`
	PharmacyRouteSegment []RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ADMIN - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_ASSOCIATED_PERSON_ASSOCIATED_RX_ADMIN struct {
	PharmacyAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
}

// PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_STUDY - Group struct
type PEX_P07_EXPERIENCE_PEX_OBSERVATION_PEX_CAUSE_STUDY struct {
	ClinicalStudyRegistration CSR `hl7:"TAG=CSR"`
	ClinicalStudyPhase []CSP `hl7:"TAG=CSP;ATR=optional"`
}

// PEX_P07 - Unsolicited initial individual product experience report
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/PEX_P07
type PEX_P07 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Visit PEX_P07_VISIT `hl7:"GROUP;ATR=optional"`
	Experience []PEX_P07_EXPERIENCE `hl7:"GROUP"`
}

