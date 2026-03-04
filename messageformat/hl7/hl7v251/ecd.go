package hl7v251

// ECD - Equipment Command
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/ECD
type ECD struct {
	ReferenceCommandNumber *int `hl7:"POS=2"`
	RemoteControlCommand CE `hl7:"POS=3"`
	ResponseRequired string `hl7:"POS=4"`
	RequestedCompletionTime TQ `hl7:"POS=5"`
	Parameters []string `hl7:"POS=6"`
}

