package hl7v22

import "time"

// EVN - Event Type
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/EVN
type EVN struct {
	EventTypeCode string `hl7:"POS=2"`
	DateTimeOfEvent time.Time `hl7:"POS=3"`
	DateTimePlannedEvent time.Time `hl7:"POS=4"`
	EventReasonCode string `hl7:"POS=5"`
	OperatorID string `hl7:"POS=6"`
}

