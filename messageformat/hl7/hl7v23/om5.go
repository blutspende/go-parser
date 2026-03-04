package hl7v23

// OM5 - Observation batteries
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OM5
type OM5 struct {
	SequenceNumber *int `hl7:"POS=2"`
	TestObservationsIncludedWAnOrderedTestBattery []CE `hl7:"POS=3"`
	ObservationIDSuffixes string `hl7:"POS=4"`
}

