package hl7v27

import "time"

// MFA - Master File Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/MFA
type MFA struct {
	RecordLevelEventCode string `hl7:"POS=2"`
	MfnControlID string `hl7:"POS=3"`
	EventCompletionDateTime time.Time `hl7:"POS=4"`
	MfnRecordLevelErrorReturn CWE `hl7:"POS=5"`
	PrimaryKeyValueMfa []string `hl7:"POS=6"`
	PrimaryKeyValueTypeMfa []string `hl7:"POS=7"`
}

