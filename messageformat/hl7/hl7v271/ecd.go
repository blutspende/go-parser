package hl7v271

// ECD - Equipment Command
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/ECD
type ECD struct {
	ReferenceCommandNumber *int `hl7:"POS=2"`
	RemoteControlCommand CWE `hl7:"POS=3"`
	ResponseRequired string `hl7:"POS=4"`
	Parameters []string `hl7:"POS=6"`
}

