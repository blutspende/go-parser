package hl7v23

// MFN_M09_MF_TEST_CATEGORICAL - Group struct
type MFN_M09_MF_TEST_CATEGORICAL struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	GeneralFieldsThatApplyToMostObservations OM1 `hl7:"TAG=OM1"`
	MfTestCatDetail MFN_M09_MF_TEST_CATEGORICAL_MF_TEST_CAT_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M09_MF_TEST_CATEGORICAL_MF_TEST_CAT_DETAIL - Group struct
type MFN_M09_MF_TEST_CATEGORICAL_MF_TEST_CAT_DETAIL struct {
	CategoricalTestObservation OM3 `hl7:"TAG=OM3"`
	ObservationsThatRequireSpecimens []OM4 `hl7:"TAG=OM4;ATR=optional"`
}

// MFN_M09 - Test/observation (categorical) master file
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFN_M09
type MFN_M09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfTestCategorical []MFN_M09_MF_TEST_CATEGORICAL `hl7:"GROUP"`
}

