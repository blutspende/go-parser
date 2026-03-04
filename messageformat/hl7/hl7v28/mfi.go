package hl7v28

import "time"

// MFI - Master File Identification
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/MFI
type MFI struct {
	MasterFileIdentifier CWE `hl7:"POS=2"`
	MasterFileApplicationIdentifier []HD `hl7:"POS=3"`
	FileLevelEventCode string `hl7:"POS=4"`
	EnteredDateTime time.Time `hl7:"POS=5"`
	EffectiveDateTime time.Time `hl7:"POS=6"`
	ResponseLevelCode string `hl7:"POS=7"`
}

