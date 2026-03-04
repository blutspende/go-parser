package hl7v231

import "time"

// OBX - Observation/result segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/OBX
type OBX struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ValueType string `hl7:"POS=3"`
	ObservationIdentifier CE `hl7:"POS=4"`
	ObservationSubID string `hl7:"POS=5"`
	ObservationValue []string `hl7:"POS=6"`
	Units CE `hl7:"POS=7"`
	ReferencesRange string `hl7:"POS=8"`
	AbnormalFlags []string `hl7:"POS=9"`
	Probability []*int `hl7:"POS=10"`
	NatureOfAbnormalTest []string `hl7:"POS=11"`
	ObservationResultStatus string `hl7:"POS=12"`
	DateLastObsNormalValues time.Time `hl7:"POS=13"`
	UserDefinedAccessChecks string `hl7:"POS=14"`
	DateTimeOfTheObservation time.Time `hl7:"POS=15"`
	ProducersID CE `hl7:"POS=16"`
	ResponsibleObserver []XCN `hl7:"POS=17"`
	ObservationMethod []CE `hl7:"POS=18"`
}

