package hl7v231

// PPP_PCB_PATIENT_VISIT - Group struct
type PPP_PCB_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPP_PCB_PATHWAY - Group struct
type PPP_PCB_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	PathwayRole []PPP_PCB_PATHWAY_PATHWAY_ROLE `hl7:"GROUP;ATR=optional"`
	Problem []PPP_PCB_PATHWAY_PROBLEM `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PATHWAY_ROLE - Group struct
type PPP_PCB_PATHWAY_PATHWAY_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM - Group struct
type PPP_PCB_PATHWAY_PROBLEM struct {
	ProblemDetail PRB `hl7:"TAG=PRB"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PPP_PCB_PATHWAY_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PPP_PCB_PATHWAY_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Goal []PPP_PCB_PATHWAY_PROBLEM_GOAL `hl7:"GROUP;ATR=optional"`
	Order []PPP_PCB_PATHWAY_PROBLEM_ORDER `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_PROBLEM_ROLE - Group struct
type PPP_PCB_PATHWAY_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PPP_PCB_PATHWAY_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_GOAL - Group struct
type PPP_PCB_PATHWAY_PROBLEM_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalObservation []PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_ROLE - Group struct
type PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION - Group struct
type PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferencesSegment ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCB - Pathway (problem-oriented) add
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/PPP_PCB
type PPP_PCB struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisit PPP_PCB_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Pathway []PPP_PCB_PATHWAY `hl7:"GROUP"`
}

