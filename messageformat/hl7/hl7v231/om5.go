package hl7v231

// OM5 - Observation batteries (sets)
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/OM5
type OM5 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	TestObservationsIncludedWithinAnOrderedTestBattery []CE `hl7:"POS=3"`
	ObservationIDSuffixes string `hl7:"POS=4"`
}

