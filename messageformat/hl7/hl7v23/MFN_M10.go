package hl7v23

// MFN_M10_MF_TEST_BATTERIES - Group struct
type MFN_M10_MF_TEST_BATTERIES struct {
	MasterFileEntrySegment MFE `hl7:"TAG=MFE"`
	GeneralFieldsThatApplyToMostObservations OM1 `hl7:"TAG=OM1"`
	MfTestBattDetail MFN_M10_MF_TEST_BATTERIES_MF_TEST_BATT_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M10_MF_TEST_BATTERIES_MF_TEST_BATT_DETAIL - Group struct
type MFN_M10_MF_TEST_BATTERIES_MF_TEST_BATT_DETAIL struct {
	ObservationBatteries OM5 `hl7:"TAG=OM5"`
	ObservationsThatRequireSpecimens []OM4 `hl7:"TAG=OM4;ATR=optional"`
}

// MFN_M10 - Test/observation batteries master file
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/MFN_M10
type MFN_M10 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MasterFileIdentificationSegment MFI `hl7:"TAG=MFI"`
	MfTestBatteries []MFN_M10_MF_TEST_BATTERIES `hl7:"GROUP"`
}

