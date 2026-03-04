package hl7v23

import "time"

// MFA - Master file acknowledgement segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MFA
type MFA struct {
	RecordLevelEventCode string `hl7:"POS=2"`
	MfnControlID string `hl7:"POS=3"`
	EventCompletionDateTime time.Time `hl7:"POS=4"`
	ErrorReturnCodeAndOrText CE `hl7:"POS=5"`
	PrimaryKeyValue []CE `hl7:"POS=6"`
}

