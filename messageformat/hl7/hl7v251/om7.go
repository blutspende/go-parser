package hl7v251

import "time"

// OM7 - Additional Basic Attributes
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/OM7
type OM7 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	UniversalServiceIdentifier CE `hl7:"POS=3"`
	CategoryIdentifier []CE `hl7:"POS=4"`
	CategoryDescription string `hl7:"POS=5"`
	CategorySynonym []string `hl7:"POS=6"`
	EffectiveTestServiceStartDateTime time.Time `hl7:"POS=7"`
	EffectiveTestServiceEndDateTime time.Time `hl7:"POS=8"`
	TestServiceDefaultDurationQuantity *int `hl7:"POS=9"`
	TestServiceDefaultDurationUnits CE `hl7:"POS=10"`
	TestServiceDefaultFrequency string `hl7:"POS=11"`
	ConsentIndicator string `hl7:"POS=12"`
	ConsentIdentifier CE `hl7:"POS=13"`
	ConsentEffectiveStartDateTime time.Time `hl7:"POS=14"`
	ConsentEffectiveEndDateTime time.Time `hl7:"POS=15"`
	ConsentIntervalQuantity *int `hl7:"POS=16"`
	ConsentIntervalUnits CE `hl7:"POS=17"`
	ConsentWaitingPeriodQuantity *int `hl7:"POS=18"`
	ConsentWaitingPeriodUnits CE `hl7:"POS=19"`
	EffectiveDateTimeOfChange time.Time `hl7:"POS=20"`
	EnteredBy XCN `hl7:"POS=21"`
	OrderableAtLocation []PL `hl7:"POS=22"`
	FormularyStatus string `hl7:"POS=23"`
	SpecialOrderIndicator string `hl7:"POS=24"`
	PrimaryKeyValueCdm []CE `hl7:"POS=25"`
}

