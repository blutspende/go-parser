package hl7v231

// PRR_PC5_PATIENT - Group struct
type PRR_PC5_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisit PRR_PC5_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Problem []PRR_PC5_PATIENT_PROBLEM `hl7:"GROUP"`
}

// PRR_PC5_PATIENT_PATIENT_VISIT - Group struct
type PRR_PC5_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM - Group struct
type PRR_PC5_PATIENT_PROBLEM struct {
	ProblemDetail PRB `hl7:"TAG=PRB"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	ProblemRole []PRR_PC5_PATIENT_PROBLEM_PROBLEM_ROLE `hl7:"GROUP;ATR=optional"`
	ProblemPathway []PRR_PC5_PATIENT_PROBLEM_PROBLEM_PATHWAY `hl7:"GROUP;ATR=optional"`
	ProblemObservation []PRR_PC5_PATIENT_PROBLEM_PROBLEM_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Goal []PRR_PC5_PATIENT_PROBLEM_GOAL `hl7:"GROUP;ATR=optional"`
	Order []PRR_PC5_PATIENT_PROBLEM_ORDER `hl7:"GROUP;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_PROBLEM_ROLE - Group struct
type PRR_PC5_PATIENT_PROBLEM_PROBLEM_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_PROBLEM_PATHWAY - Group struct
type PRR_PC5_PATIENT_PROBLEM_PROBLEM_PATHWAY struct {
	Pathway PTH `hl7:"TAG=PTH"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_PROBLEM_OBSERVATION - Group struct
type PRR_PC5_PATIENT_PROBLEM_PROBLEM_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_GOAL - Group struct
type PRR_PC5_PATIENT_PROBLEM_GOAL struct {
	GoalDetail GOL `hl7:"TAG=GOL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	GoalRole []PRR_PC5_PATIENT_PROBLEM_GOAL_GOAL_ROLE `hl7:"GROUP;ATR=optional"`
	GoalObservation []PRR_PC5_PATIENT_PROBLEM_GOAL_GOAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_GOAL_GOAL_ROLE - Group struct
type PRR_PC5_PATIENT_PROBLEM_GOAL_GOAL_ROLE struct {
	Role ROL `hl7:"TAG=ROL"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_GOAL_GOAL_OBSERVATION - Group struct
type PRR_PC5_PATIENT_PROBLEM_GOAL_GOAL_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_ORDER - Group struct
type PRR_PC5_PATIENT_PROBLEM_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL - Group struct
type PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
	OrderObservation []PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferencesSegment ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION - Group struct
type PRR_PC5_PATIENT_PROBLEM_ORDER_ORDER_DETAIL_ORDER_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Variance []VAR `hl7:"TAG=VAR;ATR=optional"`
}

// PRR_PC5 - Problem response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/PRR_PC5
type PRR_PC5 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	Patient []PRR_PC5_PATIENT `hl7:"GROUP"`
}

