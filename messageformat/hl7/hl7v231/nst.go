package hl7v231

import "time"

// NST - Statistics
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/NST
type NST struct {
	StatisticsAvailable string `hl7:"POS=2"`
	SourceIdentifier string `hl7:"POS=3"`
	SourceType string `hl7:"POS=4"`
	StatisticsStart time.Time `hl7:"POS=5"`
	StatisticsEnd time.Time `hl7:"POS=6"`
	ReceiveCharacterCount *int `hl7:"POS=7"`
	SendCharacterCount *int `hl7:"POS=8"`
	MessagesReceived *int `hl7:"POS=9"`
	MessagesSent *int `hl7:"POS=10"`
	ChecksumErrorsReceived *int `hl7:"POS=11"`
	LengthErrorsReceived *int `hl7:"POS=12"`
	OtherErrorsReceived *int `hl7:"POS=13"`
	ConnectTimeouts *int `hl7:"POS=14"`
	ReceiveTimeouts *int `hl7:"POS=15"`
	NetworkErrors *int `hl7:"POS=16"`
}

