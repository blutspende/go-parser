package hl7v25

import "time"

// PRB - Problem Details
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/PRB
type PRB struct {
	ActionCode string `hl7:"POS=2"`
	ActionDateTime time.Time `hl7:"POS=3"`
	ProblemID CE `hl7:"POS=4"`
	ProblemInstanceID EI `hl7:"POS=5"`
	EpisodeOfCareID EI `hl7:"POS=6"`
	ProblemListPriority *int `hl7:"POS=7"`
	ProblemEstablishedDateTime time.Time `hl7:"POS=8"`
	AnticipatedProblemResolutionDateTime time.Time `hl7:"POS=9"`
	ActualProblemResolutionDateTime time.Time `hl7:"POS=10"`
	ProblemClassification CE `hl7:"POS=11"`
	ProblemManagementDiscipline []CE `hl7:"POS=12"`
	ProblemPersistence CE `hl7:"POS=13"`
	ProblemConfirmationStatus CE `hl7:"POS=14"`
	ProblemLifeCycleStatus CE `hl7:"POS=15"`
	ProblemLifeCycleStatusDateTime time.Time `hl7:"POS=16"`
	ProblemDateOfOnset time.Time `hl7:"POS=17"`
	ProblemOnsetText string `hl7:"POS=18"`
	ProblemRanking CE `hl7:"POS=19"`
	CertaintyOfProblem01 CE `hl7:"POS=20"`
	ProbabilityOfProblem *int `hl7:"POS=21"`
	IndividualAwarenessOfProblem CE `hl7:"POS=22"`
	ProblemPrognosis CE `hl7:"POS=23"`
	IndividualAwarenessOfPrognosis CE `hl7:"POS=24"`
	FamilySignificantOtherAwarenessOfProblemPrognosis string `hl7:"POS=25"`
	SecuritySensitivity CE `hl7:"POS=26"`
}

