package hl7v22

// OM5 - Observation Batteries
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/OM5
type OM5 struct {
	SegmentTypeID string `hl7:"POS=2"`
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=3"`
	TestsObservationsIncludedWithinAnOrderedTestBattery []CE `hl7:"POS=4"`
	ObservationIDSuffixes string `hl7:"POS=5"`
}

