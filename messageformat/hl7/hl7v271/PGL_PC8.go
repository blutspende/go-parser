package hl7v271

// PGL_PC8_PATIENT_VISIT - Group struct
type PGL_PC8_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PGL_PC8_GOAL - Group struct
type PGL_PC8_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Goal_role []PGL_PC8_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	Pathway []PGL_PC8_GOAL_PATHWAY `hl7:"GROUP;ATR=optional"`
	Observation []PGL_PC8_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Problem []PGL_PC8_GOAL_PROBLEM `hl7:"GROUP;ATR=optional"`
	Order []PGL_PC8_GOAL_ORDER `hl7:"GROUP;ATR=optional"`
}

// PGL_PC8_GOAL_GOAL_ROLE - Group struct
type PGL_PC8_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC8_GOAL_PATHWAY - Group struct
type PGL_PC8_GOAL_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC8_GOAL_OBSERVATION - Group struct
type PGL_PC8_GOAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PGL_PC8_GOAL_PROBLEM - Group struct
type PGL_PC8_GOAL_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Problem_role []PGL_PC8_GOAL_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	Problem_observation []PGL_PC8_GOAL_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PGL_PC8_GOAL_PROBLEM_PROBLEM_ROLE - Group struct
type PGL_PC8_GOAL_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC8_GOAL_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PGL_PC8_GOAL_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PGL_PC8_GOAL_ORDER - Group struct
type PGL_PC8_GOAL_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail PGL_PC8_GOAL_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PGL_PC8_GOAL_ORDER_ORDER_DETAIL - Group struct
type PGL_PC8_GOAL_ORDER_ORDER_DETAIL struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Order_observation []PGL_PC8_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PGL_PC8_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PGL_PC8_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PGL_PC8 - PC/ goal delete
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/PGL_PC8
type PGL_PC8 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	Patient_visit PGL_PC8_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Goal []PGL_PC8_GOAL `hl7:"GROUP"`
}

