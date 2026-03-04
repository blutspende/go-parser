package hl7v271

// OM5 - Observation Batteries (sets)
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/OM5
type OM5 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	TestObservationsIncludedWithinAnOrderedTestBattery []CWE `hl7:"POS=3"`
	ObservationIDSuffixes string `hl7:"POS=4"`
}

