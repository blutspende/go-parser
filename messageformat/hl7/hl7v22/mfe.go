package hl7v22

import "time"

// MFE - Master File Entry
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/MFE
type MFE struct {
	RecordLevelEventCode string `hl7:"POS=2"`
	MfnControlID string `hl7:"POS=3"`
	EffectiveDateTime time.Time `hl7:"POS=4"`
	PrimaryKeyValue []CE `hl7:"POS=5"`
}

