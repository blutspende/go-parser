package hl7v22

import "time"

// MFA - Master File Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/MFA
type MFA struct {
	RecordLevelEventCode string `hl7:"POS=2"`
	MfnControlID string `hl7:"POS=3"`
	EventCompletionDateTime time.Time `hl7:"POS=4"`
	ErrorReturnCodeAndOrText CE `hl7:"POS=5"`
	PrimaryKeyValue []CE `hl7:"POS=6"`
}

