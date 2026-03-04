package hl7v271

// PPG_PCJ_PATIENT_VISIT - Group struct
type PPG_PCJ_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPG_PCJ_PATHWAY - Group struct
type PPG_PCJ_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Pathway_role []PPG_PCJ_PATHWAY_PATHWAY_ROLE `hl7:"GROUP;ATR=optional"`
	Goal []PPG_PCJ_PATHWAY_GOAL `hl7:"GROUP;ATR=optional"`
}

// PPG_PCJ_PATHWAY_PATHWAY_ROLE - Group struct
type PPG_PCJ_PATHWAY_PATHWAY_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL - Group struct
type PPG_PCJ_PATHWAY_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Goal_role []PPG_PCJ_PATHWAY_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	Goal_observation []PPG_PCJ_PATHWAY_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Problem []PPG_PCJ_PATHWAY_GOAL_PROBLEM `hl7:"GROUP;ATR=optional"`
	Order []PPG_PCJ_PATHWAY_GOAL_ORDER `hl7:"GROUP;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_GOAL_ROLE - Group struct
type PPG_PCJ_PATHWAY_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_GOAL_OBSERVATION - Group struct
type PPG_PCJ_PATHWAY_GOAL_GOAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_PROBLEM - Group struct
type PPG_PCJ_PATHWAY_GOAL_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Problem_role []PPG_PCJ_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	Problem_observation []PPG_PCJ_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE - Group struct
type PPG_PCJ_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PPG_PCJ_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_ORDER - Group struct
type PPG_PCJ_PATHWAY_GOAL_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail PPG_PCJ_PATHWAY_GOAL_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_ORDER_ORDER_DETAIL - Group struct
type PPG_PCJ_PATHWAY_GOAL_ORDER_ORDER_DETAIL struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Order_observation []PPG_PCJ_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPG_PCJ_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPG_PCJ_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPG_PCJ - PC/ pathway (goal-oriented) delete
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/PPG_PCJ
type PPG_PCJ struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	Patient_visit PPG_PCJ_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Pathway []PPG_PCJ_PATHWAY `hl7:"GROUP"`
}

