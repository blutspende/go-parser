package hl7v271

// MFN_M08_MF_TEST_NUMERIC - Group struct
type MFN_M08_MF_TEST_NUMERIC struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	NumericObservation OM2 `hl7:"TAG=OM2;ATR=optional"`
	CategoricalServiceTestObservation OM3 `hl7:"TAG=OM3;ATR=optional"`
	ObservationsThatRequireSpecimens OM4 `hl7:"TAG=OM4;ATR=optional"`
}

// MFN_M08 - Test/observation (numeric) master file
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/MFN_M08
type MFN_M08 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_test_numeric []MFN_M08_MF_TEST_NUMERIC `hl7:"GROUP"`
}

