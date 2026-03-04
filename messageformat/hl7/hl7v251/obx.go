package hl7v251

import "time"

// OBX - Observation/Result
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/OBX
type OBX struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ValueType string `hl7:"POS=3"`
	ObservationIdentifier CE `hl7:"POS=4"`
	ObservationSubID string `hl7:"POS=5"`
	ObservationValue []string `hl7:"POS=6"`
	Units CE `hl7:"POS=7"`
	ReferencesRange string `hl7:"POS=8"`
	AbnormalFlags []string `hl7:"POS=9"`
	Probability *int `hl7:"POS=10"`
	NatureOfAbnormalTest []string `hl7:"POS=11"`
	ObservationResultStatus string `hl7:"POS=12"`
	EffectiveDateOfReferenceRange time.Time `hl7:"POS=13"`
	UserDefinedAccessChecks string `hl7:"POS=14"`
	DateTimeOfTheObservation time.Time `hl7:"POS=15"`
	ProducersID CE `hl7:"POS=16"`
	ResponsibleObserver []XCN `hl7:"POS=17"`
	ObservationMethod []CE `hl7:"POS=18"`
	EquipmentInstanceIdentifier []EI `hl7:"POS=19"`
	DateTimeOfTheAnalysis time.Time `hl7:"POS=20"`
	PerformingOrganizationName XON `hl7:"POS=24"`
	PerformingOrganizationAddress XAD `hl7:"POS=25"`
	PerformingOrganizationMedicalDirector XCN `hl7:"POS=26"`
}

