package hl7v27

import "time"

// PEO - Product Experience Observation
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/PEO
type PEO struct {
	EventIdentifiersUsed []CWE `hl7:"POS=2"`
	EventSymptomDiagnosisCode []CWE `hl7:"POS=3"`
	EventOnsetDateTime time.Time `hl7:"POS=4"`
	EventExacerbationDateTime time.Time `hl7:"POS=5"`
	EventImprovedDateTime time.Time `hl7:"POS=6"`
	EventEndedDataTime time.Time `hl7:"POS=7"`
	EventLocationOccurredAddress []XAD `hl7:"POS=8"`
	EventQualification []string `hl7:"POS=9"`
	EventSerious string `hl7:"POS=10"`
	EventExpected string `hl7:"POS=11"`
	EventOutcome []string `hl7:"POS=12"`
	PatientOutcome string `hl7:"POS=13"`
	EventDescriptionFromOthers []string `hl7:"POS=14"`
	EventDescriptionFromOriginalReporter []string `hl7:"POS=15"`
	EventDescriptionFromPatient []string `hl7:"POS=16"`
	EventDescriptionFromPractitioner []string `hl7:"POS=17"`
	EventDescriptionFromAutopsy []string `hl7:"POS=18"`
	CauseOfDeath []CWE `hl7:"POS=19"`
	PrimaryObserverName []XPN `hl7:"POS=20"`
	PrimaryObserverAddress []XAD `hl7:"POS=21"`
	PrimaryObserverTelephone []XTN `hl7:"POS=22"`
	PrimaryObserversQualification string `hl7:"POS=23"`
	ConfirmationProvidedBy string `hl7:"POS=24"`
	PrimaryObserverAwareDateTime time.Time `hl7:"POS=25"`
	PrimaryObserversIdentityMayBeDivulged string `hl7:"POS=26"`
}

