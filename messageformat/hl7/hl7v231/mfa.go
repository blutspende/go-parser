package hl7v231

import "time"

// MFA - Master file acknowledgment segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/MFA
type MFA struct {
	RecordLevelEventCode string `hl7:"POS=2"`
	MfnControlID string `hl7:"POS=3"`
	EventCompletionDateTime time.Time `hl7:"POS=4"`
	MfnRecordLevelErrorReturn CE `hl7:"POS=5"`
	PrimaryKeyValueMfa []CE `hl7:"POS=6"`
	PrimaryKeyValueTypeMfa []string `hl7:"POS=7"`
}

