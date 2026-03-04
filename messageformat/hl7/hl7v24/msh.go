package hl7v24

import "time"

// MSH - Message Header
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSH
type MSH struct {
	SendingApplication HD `hl7:"POS=3"`
	SendingFacility HD `hl7:"POS=4"`
	ReceivingApplication HD `hl7:"POS=5"`
	ReceivingFacility HD `hl7:"POS=6"`
	DateTimeOfMessage time.Time `hl7:"POS=7"`
	Security string `hl7:"POS=8"`
	MessageType MSG `hl7:"POS=9"`
	MessageControlID string `hl7:"POS=10"`
	ProcessingID PT `hl7:"POS=11"`
	VersionID VID `hl7:"POS=12"`
	SequenceNumber *int `hl7:"POS=13"`
	ContinuationPointer string `hl7:"POS=14"`
	AcceptAcknowledgmentType string `hl7:"POS=15"`
	ApplicationAcknowledgmentType string `hl7:"POS=16"`
	CountryCode string `hl7:"POS=17"`
	CharacterSet []string `hl7:"POS=18"`
	PrincipalLanguageOfMessage CE `hl7:"POS=19"`
	AlternateCharacterSetHandlingScheme string `hl7:"POS=20"`
	ConformanceStatementID []string `hl7:"POS=21"`
}

