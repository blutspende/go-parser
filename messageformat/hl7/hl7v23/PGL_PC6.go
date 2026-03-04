package hl7v23

// PGL_PC6_PATIENT_VISIT - Group struct
type PGL_PC6_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PGL_PC6_GOAL - Group struct
type PGL_PC6_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PGL_PC6_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	Pathway []PGL_PC6_GOAL_PATHWAY `hl7:"GROUP;ATR=optional"`
	Observation []PGL_PC6_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Problem []PGL_PC6_GOAL_PROBLEM `hl7:"GROUP;ATR=optional"`
	Order []PGL_PC6_GOAL_ORDER `hl7:"GROUP;ATR=optional"`
}

// PGL_PC6_GOAL_GOAL_ROLE - Group struct
type PGL_PC6_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC6_GOAL_PATHWAY - Group struct
type PGL_PC6_GOAL_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC6_GOAL_OBSERVATION - Group struct
type PGL_PC6_GOAL_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PGL_PC6_GOAL_PROBLEM - Group struct
type PGL_PC6_GOAL_PROBLEM struct {
	ProblemDetail PRB `hl7:"TAG=PRB"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PGL_PC6_GOAL_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PGL_PC6_GOAL_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PGL_PC6_GOAL_PROBLEM_PROBLEM_ROLE - Group struct
type PGL_PC6_GOAL_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC6_GOAL_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PGL_PC6_GOAL_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PGL_PC6_GOAL_ORDER - Group struct
type PGL_PC6_GOAL_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail PGL_PC6_GOAL_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PGL_PC6_GOAL_ORDER_ORDER_DETAIL - Group struct
type PGL_PC6_GOAL_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PGL_PC6_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PGL_PC6_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PGL_PC6_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PGL_PC6_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyPrescriptionOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PGL_PC6_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PGL_PC6_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC6 - Goal add
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/PGL_PC6
type PGL_PC6 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PGL_PC6_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Goal []PGL_PC6_GOAL `hl7:"GROUP"`
}

