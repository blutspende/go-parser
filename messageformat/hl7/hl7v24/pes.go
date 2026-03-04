package hl7v24

import "time"

// PES - Product Experience Sender
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PES
type PES struct {
	SenderOrganizationName []XON `hl7:"POS=2"`
	SenderIndividualName []XCN `hl7:"POS=3"`
	SenderAddress []XAD `hl7:"POS=4"`
	SenderTelephone []XTN `hl7:"POS=5"`
	SenderEventIdentifier EI `hl7:"POS=6"`
	SenderSequenceNumber *int `hl7:"POS=7"`
	SenderEventDescription []string `hl7:"POS=8"`
	SenderComment string `hl7:"POS=9"`
	SenderAwareDateTime time.Time `hl7:"POS=10"`
	EventReportDate time.Time `hl7:"POS=11"`
	EventReportTimingType []string `hl7:"POS=12"`
	EventReportSource string `hl7:"POS=13"`
	EventReportedTo []string `hl7:"POS=14"`
}

