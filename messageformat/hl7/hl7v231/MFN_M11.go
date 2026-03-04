package hl7v231

// MFN_M11_MF_TEST_CALCULATED - Group struct
type MFN_M11_MF_TEST_CALCULATED struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	MfTestCalcDetail MFN_M11_MF_TEST_CALCULATED_MF_TEST_CALC_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M11_MF_TEST_CALCULATED_MF_TEST_CALC_DETAIL - Group struct
type MFN_M11_MF_TEST_CALCULATED_MF_TEST_CALC_DETAIL struct {
	ObservationsThatAreCalculatedFromOtherObservationsSegment OM6 `hl7:"TAG=OM6"`
	NumericObservationSegment OM2 `hl7:"TAG=OM2"`
}

// MFN_M11 - Test/calculated observations master file
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/MFN_M11
type MFN_M11 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfTestCalculated []MFN_M11_MF_TEST_CALCULATED `hl7:"GROUP"`
}

