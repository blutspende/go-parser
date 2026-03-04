package hl7v251

// OM5 - Observation Batteries (Sets)
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/OM5
type OM5 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	TestObservationsIncludedWithinAnOrderedTestBattery []CE `hl7:"POS=3"`
	ObservationIDSuffixes string `hl7:"POS=4"`
}

