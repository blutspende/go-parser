package hl7v28

// PPP_PCB_PATIENT_VISIT - Group struct
type PPP_PCB_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPP_PCB_PATHWAY - Group struct
type PPP_PCB_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Pathway_role []PPP_PCB_PATHWAY_PATHWAY_ROLE `hl7:"GROUP;ATR=optional"`
	Problem []PPP_PCB_PATHWAY_PROBLEM `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PATHWAY_ROLE - Group struct
type PPP_PCB_PATHWAY_PATHWAY_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM - Group struct
type PPP_PCB_PATHWAY_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Problem_role []PPP_PCB_PATHWAY_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	Problem_observation []PPP_PCB_PATHWAY_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
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
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_GOAL - Group struct
type PPP_PCB_PATHWAY_PROBLEM_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Goal_role []PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	Goal_observation []PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_ROLE - Group struct
type PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION - Group struct
type PPP_PCB_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL struct {
	Choice PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_CHOICE `hl7:"GROUP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	Order_observation []PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_CHOICE - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_CHOICE struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	AnyHl7Segment Hxx `hl7:"TAG=Hxx;ATR=optional"`
}

// PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPP_PCB_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCB - PC/ pathway (problem-oriented) add
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/PPP_PCB
type PPP_PCB struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	Patient_visit PPP_PCB_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Pathway []PPP_PCB_PATHWAY `hl7:"GROUP"`
}

