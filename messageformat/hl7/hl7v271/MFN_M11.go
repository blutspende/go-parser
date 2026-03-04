package hl7v271

// MFN_M11_MF_TEST_CALCULATED - Group struct
type MFN_M11_MF_TEST_CALCULATED struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	Mf_test_calc_detail MFN_M11_MF_TEST_CALCULATED_MF_TEST_CALC_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M11_MF_TEST_CALCULATED_MF_TEST_CALC_DETAIL - Group struct
type MFN_M11_MF_TEST_CALCULATED_MF_TEST_CALC_DETAIL struct {
	ObservationsThatAreCalculatedFromOtherObservations OM6 `hl7:"TAG=OM6"`
	NumericObservation OM2 `hl7:"TAG=OM2"`
}

// MFN_M11 - Test/calculated observations master file
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/MFN_M11
type MFN_M11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_test_calculated []MFN_M11_MF_TEST_CALCULATED `hl7:"GROUP"`
}

