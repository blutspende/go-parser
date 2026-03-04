package hl7v251

// PPV_PCA_PATIENT - Group struct
type PPV_PCA_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PPV_PCA_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Goal []PPV_PCA_PATIENT_GOAL `hl7:"GROUP"`
}

// PPV_PCA_PATIENT_PATIENT_VISIT - Group struct
type PPV_PCA_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL - Group struct
type PPV_PCA_PATIENT_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PPV_PCA_PATIENT_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalPathway []PPV_PCA_PATIENT_GOAL_GOAL_PATHWAY `hl7:"GROUP;ATR=optional"`
	GoalObservation []PPV_PCA_PATIENT_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Problem []PPV_PCA_PATIENT_GOAL_PROBLEM `hl7:"GROUP;ATR=optional"`
	Order []PPV_PCA_PATIENT_GOAL_ORDER `hl7:"GROUP;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_GOAL_ROLE - Group struct
type PPV_PCA_PATIENT_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_GOAL_PATHWAY - Group struct
type PPV_PCA_PATIENT_GOAL_GOAL_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_GOAL_OBSERVATION - Group struct
type PPV_PCA_PATIENT_GOAL_GOAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_PROBLEM - Group struct
type PPV_PCA_PATIENT_GOAL_PROBLEM struct {
	ProblemDetails PRB `hl7:"TAG=PRB"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PPV_PCA_PATIENT_GOAL_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PPV_PCA_PATIENT_GOAL_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_PROBLEM_PROBLEM_ROLE - Group struct
type PPV_PCA_PATIENT_GOAL_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PPV_PCA_PATIENT_GOAL_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_ORDER - Group struct
type PPV_PCA_PATIENT_GOAL_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetail PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL - Group struct
type PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructions ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PPV_PCA_PATIENT_GOAL_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PPV_PCA - Patient goal response
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/PPV_PCA
type PPV_PCA struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	Patient []PPV_PCA_PATIENT `hl7:"GROUP"`
}

