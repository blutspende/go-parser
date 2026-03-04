package hl7v26

import "time"

// PRB - Problem Details
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/PRB
type PRB struct {
	ActionCode string `hl7:"POS=2"`
	ActionDateTime time.Time `hl7:"POS=3"`
	ProblemID CWE `hl7:"POS=4"`
	ProblemInstanceID EI `hl7:"POS=5"`
	EpisodeOfCareID EI `hl7:"POS=6"`
	ProblemListPriority *int `hl7:"POS=7"`
	ProblemEstablishedDateTime time.Time `hl7:"POS=8"`
	AnticipatedProblemResolutionDateTime time.Time `hl7:"POS=9"`
	ActualProblemResolutionDateTime time.Time `hl7:"POS=10"`
	ProblemClassification CWE `hl7:"POS=11"`
	ProblemManagementDiscipline []CWE `hl7:"POS=12"`
	ProblemPersistence CWE `hl7:"POS=13"`
	ProblemConfirmationStatus CWE `hl7:"POS=14"`
	ProblemLifeCycleStatus CWE `hl7:"POS=15"`
	ProblemLifeCycleStatusDateTime time.Time `hl7:"POS=16"`
	ProblemDateOfOnset time.Time `hl7:"POS=17"`
	ProblemOnsetText string `hl7:"POS=18"`
	ProblemRanking CWE `hl7:"POS=19"`
	CertaintyOfProblem CWE `hl7:"POS=20"`
	ProbabilityOfProblem *int `hl7:"POS=21"`
	IndividualAwarenessOfProblem CWE `hl7:"POS=22"`
	ProblemPrognosis CWE `hl7:"POS=23"`
	IndividualAwarenessOfPrognosis CWE `hl7:"POS=24"`
	FamilySignificantOtherAwarenessOfProblemPrognosis string `hl7:"POS=25"`
	SecuritySensitivity CWE `hl7:"POS=26"`
	ProblemSeverity CWE `hl7:"POS=27"`
	ProblemPerspective CWE `hl7:"POS=28"`
	MoodCode CNE `hl7:"POS=29"`
}

