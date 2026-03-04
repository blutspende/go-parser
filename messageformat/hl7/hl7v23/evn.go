package hl7v23

import "time"

// EVN - Event type
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/EVN
type EVN struct {
	EventTypeCode string `hl7:"POS=2"`
	RecordedDateTime time.Time `hl7:"POS=3"`
	DateTimePlannedEvent time.Time `hl7:"POS=4"`
	EventReasonCode string `hl7:"POS=5"`
	OperatorID XCN `hl7:"POS=6"`
	EventOccurred time.Time `hl7:"POS=7"`
}

