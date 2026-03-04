package hl7v231

// PPG_PCH_PATIENT_VISIT - Group struct
type PPG_PCH_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPG_PCH_PATHWAY - Group struct
type PPG_PCH_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	PathwayRole []PPG_PCH_PATHWAY_PATHWAY_ROLE `hl7:"GROUP;ATR=optional"`
	Goal []PPG_PCH_PATHWAY_GOAL `hl7:"GROUP;ATR=optional"`
}

// PPG_PCH_PATHWAY_PATHWAY_ROLE - Group struct
type PPG_PCH_PATHWAY_PATHWAY_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL - Group struct
type PPG_PCH_PATHWAY_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PPG_PCH_PATHWAY_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalObservation []PPG_PCH_PATHWAY_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Problem []PPG_PCH_PATHWAY_GOAL_PROBLEM `hl7:"GROUP;ATR=optional"`
	Order []PPG_PCH_PATHWAY_GOAL_ORDER `hl7:"GROUP;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_GOAL_ROLE - Group struct
type PPG_PCH_PATHWAY_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_GOAL_OBSERVATION - Group struct
type PPG_PCH_PATHWAY_GOAL_GOAL_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_PROBLEM - Group struct
type PPG_PCH_PATHWAY_GOAL_PROBLEM struct {
	ProblemDetail PRB `hl7:"TAG=PRB"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PPG_PCH_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PPG_PCH_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE - Group struct
type PPG_PCH_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PPG_PCH_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_ORDER - Group struct
type PPG_PCH_PATHWAY_GOAL_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL - Group struct
type PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferencesSegment ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPG_PCH_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCH - Pathway (goal-oriented) update
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/PPG_PCH
type PPG_PCH struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisit PPG_PCH_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Pathway []PPG_PCH_PATHWAY `hl7:"GROUP"`
}

