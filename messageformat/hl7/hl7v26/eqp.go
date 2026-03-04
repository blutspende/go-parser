package hl7v26

import "time"

// EQP - Equipment/log Service
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/EQP
type EQP struct {
	EventType CWE `hl7:"POS=2"`
	FileName string `hl7:"POS=3"`
	StartDateTime time.Time `hl7:"POS=4"`
	EndDateTime time.Time `hl7:"POS=5"`
	TransactionData string `hl7:"POS=6"`
}

