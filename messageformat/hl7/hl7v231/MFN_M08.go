package hl7v231

// MFN_M08_MF_TEST_NUMERIC - Group struct
type MFN_M08_MF_TEST_NUMERIC struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	MfNumericObservation MFN_M08_MF_TEST_NUMERIC_MF_NUMERIC_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// MFN_M08_MF_TEST_NUMERIC_MF_NUMERIC_OBSERVATION - Group struct
type MFN_M08_MF_TEST_NUMERIC_MF_NUMERIC_OBSERVATION struct {
	NumericObservationSegment OM2 `hl7:"TAG=OM2;ATR=optional"`
	CategoricalTestObservationSegment OM3 `hl7:"TAG=OM3;ATR=optional"`
	ObservationsThatRequireSpecimensSegment OM4 `hl7:"TAG=OM4;ATR=optional"`
}

// MFN_M08 - Test/observation (numeric) master file
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFN_M08
type MFN_M08 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfTestNumeric []MFN_M08_MF_TEST_NUMERIC `hl7:"GROUP"`
}

