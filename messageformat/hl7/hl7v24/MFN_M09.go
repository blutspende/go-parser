package hl7v24

// MFN_M09_MF_TEST_CATEGORICAL - Group struct
type MFN_M09_MF_TEST_CATEGORICAL struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	MfTestCatDetail MFN_M09_MF_TEST_CATEGORICAL_MF_TEST_CAT_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M09_MF_TEST_CATEGORICAL_MF_TEST_CAT_DETAIL - Group struct
type MFN_M09_MF_TEST_CATEGORICAL_MF_TEST_CAT_DETAIL struct {
	CategoricalServiceTestObservation OM3 `hl7:"TAG=OM3"`
	ObservationsThatRequireSpecimens []OM4 `hl7:"TAG=OM4;ATR=optional"`
}

// MFN_M09 - Master files notification - Test/observation (categorical) master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFN_M09
type MFN_M09 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfTestCategorical []MFN_M09_MF_TEST_CATEGORICAL `hl7:"GROUP"`
}

