package hl7v23

import "time"

// MSH - Message header segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MSH
type MSH struct {
	SendingApplication HD `hl7:"POS=3"`
	SendingFacility HD `hl7:"POS=4"`
	ReceivingApplication HD `hl7:"POS=5"`
	ReceivingFacility HD `hl7:"POS=6"`
	DateTimeOfMessage time.Time `hl7:"POS=7"`
	Security string `hl7:"POS=8"`
	MessageType CM_MSG `hl7:"POS=9"`
	MessageControlID string `hl7:"POS=10"`
	ProcessingID PT `hl7:"POS=11"`
	VersionID string `hl7:"POS=12"`
	SequenceNumber *int `hl7:"POS=13"`
	ContinuationPointer string `hl7:"POS=14"`
	AcceptAcknowledgementType string `hl7:"POS=15"`
	ApplicationAcknowledgementType string `hl7:"POS=16"`
	CountryCode string `hl7:"POS=17"`
	CharacterSet []string `hl7:"POS=18"`
	PrincipalLanguageOfMessage CE `hl7:"POS=19"`
}

