package hl7v27

// PPR_PC2_PATIENT_VISIT - Group struct
type PPR_PC2_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPR_PC2_PROBLEM - Group struct
type PPR_PC2_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Problem_role []PPR_PC2_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	Pathway []PPR_PC2_PROBLEM_PATHWAY `hl7:"GROUP;ATR=optional"`
	Problem_observation []PPR_PC2_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
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
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPR_PC2_PROBLEM_GOAL - Group struct
type PPR_PC2_PROBLEM_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Goal_role []PPR_PC2_PROBLEM_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	Goal_observation []PPR_PC2_PROBLEM_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPR_PC2_PROBLEM_GOAL_GOAL_ROLE - Group struct
type PPR_PC2_PROBLEM_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPR_PC2_PROBLEM_GOAL_GOAL_OBSERVATION - Group struct
type PPR_PC2_PROBLEM_GOAL_GOAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPR_PC2_PROBLEM_ORDER - Group struct
type PPR_PC2_PROBLEM_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL - Group struct
type PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Order_observation []PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPR_PC2_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPR_PC2 - PC/ problem update
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/PPR_PC2
type PPR_PC2 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	Patient_visit PPR_PC2_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Problem []PPR_PC2_PROBLEM `hl7:"GROUP"`
}

