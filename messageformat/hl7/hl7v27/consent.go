package hl7v27

import "time"

// CON - Consent Segment
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/CON
type CON struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ConsentType CWE `hl7:"POS=3"`
	ConsentFormIDAndVersion string `hl7:"POS=4"`
	ConsentFormNumber EI `hl7:"POS=5"`
	ConsentText []string `hl7:"POS=6"`
	SubjectSpecificConsentText []string `hl7:"POS=7"`
	ConsentBackgroundInformation []string `hl7:"POS=8"`
	SubjectSpecificConsentBackgroundText []string `hl7:"POS=9"`
	ConsenterImposedLimitations []string `hl7:"POS=10"`
	ConsentMode CNE `hl7:"POS=11"`
	ConsentStatus CNE `hl7:"POS=12"`
	ConsentDiscussionDateTime time.Time `hl7:"POS=13"`
	ConsentDecisionDateTime time.Time `hl7:"POS=14"`
	ConsentEffectiveDateTime time.Time `hl7:"POS=15"`
	ConsentEndDateTime time.Time `hl7:"POS=16"`
	SubjectCompetenceIndicator string `hl7:"POS=17"`
	TranslatorAssistanceIndicator string `hl7:"POS=18"`
	LanguageTranslatedTo CWE `hl7:"POS=19"`
	InformationalMaterialSuppliedIndicator string `hl7:"POS=20"`
	ConsentBypassReason CWE `hl7:"POS=21"`
	ConsentDisclosureLevel string `hl7:"POS=22"`
	ConsentNonDisclosureReason CWE `hl7:"POS=23"`
	NonSubjectConsenterReason CWE `hl7:"POS=24"`
	ConsenterID []XPN `hl7:"POS=25"`
	RelationshipToSubject []CWE `hl7:"POS=26"`
}

