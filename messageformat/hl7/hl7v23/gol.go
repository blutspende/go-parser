package hl7v23

import "time"

// GOL - Goal Detail
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/GOL
type GOL struct {
	ActionCode string `hl7:"POS=2"`
	ActionDateTime time.Time `hl7:"POS=3"`
	GoalID CE `hl7:"POS=4"`
	GoalInstanceID EI `hl7:"POS=5"`
	EpisodeOfCareID EI `hl7:"POS=6"`
	GoalListPriority *int `hl7:"POS=7"`
	GoalEstablishedDateTime time.Time `hl7:"POS=8"`
	ExpectedGoalAchievementDateTime time.Time `hl7:"POS=9"`
	GoalClassification CE `hl7:"POS=10"`
	GoalManagementDiscipline CE `hl7:"POS=11"`
	CurrentGoalReviewStatus CE `hl7:"POS=12"`
	CurrentGoalReviewDateTime time.Time `hl7:"POS=13"`
	NextGoalReviewDateTime time.Time `hl7:"POS=14"`
	PreviousGoalReviewDateTime time.Time `hl7:"POS=15"`
	GoalReviewInterval TQ `hl7:"POS=16"`
	GoalEvaluation CE `hl7:"POS=17"`
	GoalEvaluationComment []string `hl7:"POS=18"`
	GoalLifeCycleStatus CE `hl7:"POS=19"`
	GoalLifeCycleStatusDateTime time.Time `hl7:"POS=20"`
	GoalTargetType []CE `hl7:"POS=21"`
	GoalTargetName []XPN `hl7:"POS=22"`
}

