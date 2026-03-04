package hl7v251

// PPP_PCC_PATIENT_VISIT - Group struct
type PPP_PCC_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPP_PCC_PATHWAY - Group struct
type PPP_PCC_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	PathwayRole []PPP_PCC_PATHWAY_PATHWAY_ROLE `hl7:"GROUP;ATR=optional"`
	Problem []PPP_PCC_PATHWAY_PROBLEM `hl7:"GROUP;ATR=optional"`
}

// PPP_PCC_PATHWAY_PATHWAY_ROLE - Group struct
type PPP_PCC_PATHWAY_PATHWAY_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM - Group struct
type PPP_PCC_PATHWAY_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PPP_PCC_PATHWAY_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PPP_PCC_PATHWAY_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Goal []PPP_PCC_PATHWAY_PROBLEM_GOAL `hl7:"GROUP;ATR=optional"`
	Order []PPP_PCC_PATHWAY_PROBLEM_ORDER `hl7:"GROUP;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_PROBLEM_ROLE - Group struct
type PPP_PCC_PATHWAY_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PPP_PCC_PATHWAY_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_GOAL - Group struct
type PPP_PCC_PATHWAY_PROBLEM_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PPP_PCC_PATHWAY_PROBLEM_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalObservation []PPP_PCC_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_GOAL_GOAL_ROLE - Group struct
type PPP_PCC_PATHWAY_PROBLEM_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION - Group struct
type PPP_PCC_PATHWAY_PROBLEM_GOAL_GOAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_ORDER - Group struct
type PPP_PCC_PATHWAY_PROBLEM_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetail PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL - Group struct
type PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructions ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPP_PCC_PATHWAY_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPP_PCC - Pathway (problem-oriented) update
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/PPP_PCC
type PPP_PCC struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PPP_PCC_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Pathway []PPP_PCC_PATHWAY `hl7:"GROUP"`
}

