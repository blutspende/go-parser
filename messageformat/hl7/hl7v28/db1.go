package hl7v28

import "time"

// DB1 - Disability
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/DB1
type DB1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	DisabledPersonCode CWE `hl7:"POS=3"`
	DisabledPersonIdentifier []CX `hl7:"POS=4"`
	DisabilityIndicator string `hl7:"POS=5"`
	DisabilityStartDate time.Time `hl7:"POS=6;ATR=date"`
	DisabilityEndDate time.Time `hl7:"POS=7;ATR=date"`
	DisabilityReturnToWorkDate time.Time `hl7:"POS=8;ATR=date"`
	DisabilityUnableToWorkDate time.Time `hl7:"POS=9;ATR=date"`
}

