package hl7v25

// OM2 - Numeric Observation
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/OM2
type OM2 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	UnitsOfMeasure CE `hl7:"POS=3"`
	RangeOfDecimalPrecision []*int `hl7:"POS=4"`
	CorrespondingSiUnitsOfMeasure CE `hl7:"POS=5"`
	SiConversionFactor string `hl7:"POS=6"`
	ReferenceNormalRangeOrdinalAndContinuousObservations []RFR `hl7:"POS=7"`
	CriticalRangeForOrdinalAndContinuousObservations []RFR `hl7:"POS=8"`
	AbsoluteRangeForOrdinalAndContinuousObservations RFR `hl7:"POS=9"`
	DeltaCheckCriteria []DLT `hl7:"POS=10"`
	MinimumMeaningfulIncrements *int `hl7:"POS=11"`
}

