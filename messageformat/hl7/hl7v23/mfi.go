package hl7v23

import "time"

// MFI - Master file identification segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MFI
type MFI struct {
	MasterFileIdentifier CE `hl7:"POS=2"`
	MasterFileApplicationIdentifier HD `hl7:"POS=3"`
	FileLevelEventCode string `hl7:"POS=4"`
	EnteredDateTime time.Time `hl7:"POS=5"`
	EffectiveDateTime time.Time `hl7:"POS=6"`
	ResponseLevelCode string `hl7:"POS=7"`
}

