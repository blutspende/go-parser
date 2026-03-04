package hl7v271

import "time"

// GOL - Goal Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/GOL
type GOL struct {
	ActionCode string `hl7:"POS=2"`
	ActionDateTime time.Time `hl7:"POS=3"`
	GoalID CWE `hl7:"POS=4"`
	GoalInstanceID EI `hl7:"POS=5"`
	EpisodeOfCareID EI `hl7:"POS=6"`
	GoalListPriority *int `hl7:"POS=7"`
	GoalEstablishedDateTime time.Time `hl7:"POS=8"`
	ExpectedGoalAchieveDateTime time.Time `hl7:"POS=9"`
	GoalClassification CWE `hl7:"POS=10"`
	GoalManagementDiscipline CWE `hl7:"POS=11"`
	CurrentGoalReviewStatus CWE `hl7:"POS=12"`
	CurrentGoalReviewDateTime time.Time `hl7:"POS=13"`
	NextGoalReviewDateTime time.Time `hl7:"POS=14"`
	PreviousGoalReviewDateTime time.Time `hl7:"POS=15"`
	GoalEvaluation CWE `hl7:"POS=17"`
	GoalEvaluationComment []string `hl7:"POS=18"`
	GoalLifeCycleStatus CWE `hl7:"POS=19"`
	GoalLifeCycleStatusDateTime time.Time `hl7:"POS=20"`
	GoalTargetType []CWE `hl7:"POS=21"`
	GoalTargetName []XPN `hl7:"POS=22"`
	MoodCode CNE `hl7:"POS=23"`
}

