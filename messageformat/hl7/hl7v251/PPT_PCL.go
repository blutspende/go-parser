package hl7v251

// PPT_PCL_PATIENT - Group struct
type PPT_PCL_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PPT_PCL_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Pathway []PPT_PCL_PATIENT_PATHWAY `hl7:"GROUP"`
}

// PPT_PCL_PATIENT_PATIENT_VISIT - Group struct
type PPT_PCL_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY - Group struct
type PPT_PCL_PATIENT_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	PathwayRole []PPT_PCL_PATIENT_PATHWAY_PATHWAY_ROLE `hl7:"GROUP;ATR=optional"`
	Goal []PPT_PCL_PATIENT_PATHWAY_GOAL `hl7:"GROUP;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_PATHWAY_ROLE - Group struct
type PPT_PCL_PATIENT_PATHWAY_PATHWAY_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PPT_PCL_PATIENT_PATHWAY_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalObservation []PPT_PCL_PATIENT_PATHWAY_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Problem []PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM `hl7:"GROUP;ATR=optional"`
	Order []PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER `hl7:"GROUP;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_GOAL_ROLE - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_GOAL_OBSERVATION - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_GOAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetail PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructions ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPT_PCL_PATIENT_PATHWAY_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPT_PCL - Pathway (goal-oriented) query response
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/PPT_PCL
type PPT_PCL struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	Patient []PPT_PCL_PATIENT `hl7:"GROUP"`
}

