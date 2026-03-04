package hl7v23

// PPR_PC2_PATIENT_VISIT - Group struct
type PPR_PC2_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPR_PC2_PROBLEM - Group struct
type PPR_PC2_PROBLEM struct {
	ProblemDetail PRB `hl7:"TAG=PRB"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PPR_PC2_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	Pathway []PPR_PC2_PROBLEM_PATHWAY `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PPR_PC2_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Goal []PPR_PC2_PROBLEM_GOAL `hl7:"GROUP;ATR=optional"`
	Order []PPR_PC2_PROBLEM_ORDER `hl7:"GROUP;ATR=optional"`
}

// PPR_PC2_PROBLEM_PROBLEM_ROLE - Group struct
type PPR_PC2_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPR_PC2_PROBLEM_PATHWAY - Group struct
type PPR_PC2_PROBLEM_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPR_PC2_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PPR_PC2_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPR_PC2_PROBLEM_GOAL - Group struct
type PPR_PC2_PROBLEM_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PPR_PC2_PROBLEM_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalObservation []PPR_PC2_PROBLEM_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPR_PC2_PROBLEM_GOAL_GOAL_ROLE - Group struct
type PPR_PC2_PROBLEM_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPR_PC2_PROBLEM_GOAL_GOAL_OBSERVATION - Group struct
type PPR_PC2_PROBLEM_GOAL_GOAL_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPR_PC2_PROBLEM_ORDER - Group struct
type PPR_PC2_PROBLEM_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL - Group struct
type PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyPrescriptionOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPR_PC2 - Problem update
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/PPR_PC2
type PPR_PC2 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PPR_PC2_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Problem []PPR_PC2_PROBLEM `hl7:"GROUP"`
}

