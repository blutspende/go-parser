package hl7v231

// PTR_PCF_PATIENT - Group struct
type PTR_PCF_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisit PTR_PCF_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Pathway []PTR_PCF_PATIENT_PATHWAY `hl7:"GROUP"`
}

// PTR_PCF_PATIENT_PATIENT_VISIT - Group struct
type PTR_PCF_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY - Group struct
type PTR_PCF_PATIENT_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	PathwayRole []PTR_PCF_PATIENT_PATHWAY_PATHWAY_ROLE `hl7:"GROUP;ATR=optional"`
	Problem []PTR_PCF_PATIENT_PATHWAY_PROBLEM `hl7:"GROUP;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PATHWAY_ROLE - Group struct
type PTR_PCF_PATIENT_PATHWAY_PATHWAY_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM struct {
	ProblemDetail PRB `hl7:"TAG=PRB"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PTR_PCF_PATIENT_PATHWAY_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PTR_PCF_PATIENT_PATHWAY_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Goal []PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL `hl7:"GROUP;ATR=optional"`
	Order []PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER `hl7:"GROUP;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_PROBLEM_ROLE - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalObservation []PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL_GOAL_ROLE - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferencesSegment ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PTR_PCF_PATIENT_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PTR_PCF - Pathway (problem-oriented) query response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/PTR_PCF
type PTR_PCF struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	Patient []PTR_PCF_PATIENT `hl7:"GROUP"`
}

