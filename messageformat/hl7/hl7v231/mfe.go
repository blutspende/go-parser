package hl7v231

import "time"

// MFE - Master file entry segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/MFE
type MFE struct {
	RecordLevelEventCode string `hl7:"POS=2"`
	MfnControlID string `hl7:"POS=3"`
	EffectiveDateTime time.Time `hl7:"POS=4"`
	PrimaryKeyValueMfe []string `hl7:"POS=5"`
	PrimaryKeyValueType []string `hl7:"POS=6"`
}

