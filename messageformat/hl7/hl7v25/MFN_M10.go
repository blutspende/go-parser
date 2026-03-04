package hl7v25

// MFN_M10_MF_TEST_BATTERIES - Group struct
type MFN_M10_MF_TEST_BATTERIES struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	MfTestBattDetail MFN_M10_MF_TEST_BATTERIES_MF_TEST_BATT_DETAIL `hl7:"GROUP;ATR=optional"`
}

// MFN_M10_MF_TEST_BATTERIES_MF_TEST_BATT_DETAIL - Group struct
type MFN_M10_MF_TEST_BATTERIES_MF_TEST_BATT_DETAIL struct {
	ObservationBatteriesSets OM5 `hl7:"TAG=OM5"`
	ObservationsThatRequireSpecimens []OM4 `hl7:"TAG=OM4;ATR=optional"`
}

// MFN_M10 - Master File Notification - Test/Observation Batteries
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/MFN_M10
type MFN_M10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfTestBatteries []MFN_M10_MF_TEST_BATTERIES `hl7:"GROUP"`
}

