package hl7v24

import "time"

// CSR - Clinical Study Registration
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CSR
type CSR struct {
	SponsorStudyID EI `hl7:"POS=2"`
	AlternateStudyID EI `hl7:"POS=3"`
	InstitutionRegisteringThePatient CE `hl7:"POS=4"`
	SponsorPatientID CX `hl7:"POS=5"`
	AlternatePatientIDCsr CX `hl7:"POS=6"`
	DateTimeOfPatientStudyRegistration time.Time `hl7:"POS=7"`
	PersonPerformingStudyRegistration []XCN `hl7:"POS=8"`
	StudyAuthorizingProvider []XCN `hl7:"POS=9"`
	DateTimePatientStudyConsentSigned time.Time `hl7:"POS=10"`
	PatientStudyEligibilityStatus CE `hl7:"POS=11"`
	StudyRandomizationDateTime []time.Time `hl7:"POS=12"`
	RandomizedStudyArm []CE `hl7:"POS=13"`
	StratumForStudyRandomization []CE `hl7:"POS=14"`
	PatientEvaluabilityStatus CE `hl7:"POS=15"`
	DateTimeEndedStudy time.Time `hl7:"POS=16"`
	ReasonEndedStudy CE `hl7:"POS=17"`
}

