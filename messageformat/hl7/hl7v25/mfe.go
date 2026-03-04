package hl7v25

import "time"

// MFE - Master File Entry
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/MFE
type MFE struct {
	RecordLevelEventCode string `hl7:"POS=2"`
	MfnControlID string `hl7:"POS=3"`
	EffectiveDateTime time.Time `hl7:"POS=4"`
	PrimaryKeyValueMfe []string `hl7:"POS=5"`
	PrimaryKeyValueType []string `hl7:"POS=6"`
}

