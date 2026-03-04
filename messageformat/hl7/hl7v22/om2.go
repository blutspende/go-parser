package hl7v22

// OM2 - Numeric Observation
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/OM2
type OM2 struct {
	SegmentTypeID string `hl7:"POS=2"`
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=3"`
	UnitsOfMeasure CE `hl7:"POS=4"`
	RangeOfDecimalPrecision []*int `hl7:"POS=5"`
	CorrespondingSiUnitsOfMeasure CE `hl7:"POS=6"`
	SiConversionFactor string `hl7:"POS=7"`
	ReferenceNormalRangeOrdinalContinuousObservations CM_RFR `hl7:"POS=8"`
	CriticalRangeForOrdinalAndContinuousObservations CM_RANGE `hl7:"POS=9"`
	AbsoluteRangeForOrdinalAndContinuousObservations CM_ABS_RANGE `hl7:"POS=10"`
	DeltaCheckCriteria []CM_DLT `hl7:"POS=11"`
	MinimumMeaningfulIncrements *int `hl7:"POS=12"`
}

