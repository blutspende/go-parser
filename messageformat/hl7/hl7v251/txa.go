package hl7v251

import "time"

// TXA - Transcription Document Header
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/TXA
type TXA struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	DocumentType string `hl7:"POS=3"`
	DocumentContentPresentation string `hl7:"POS=4"`
	ActivityDateTime time.Time `hl7:"POS=5"`
	PrimaryActivityProviderCodeName []XCN `hl7:"POS=6"`
	OriginationDateTime time.Time `hl7:"POS=7"`
	TranscriptionDateTime time.Time `hl7:"POS=8"`
	EditDateTime []time.Time `hl7:"POS=9"`
	OriginatorCodeName []XCN `hl7:"POS=10"`
	AssignedDocumentAuthenticator []XCN `hl7:"POS=11"`
	TranscriptionistCodeName []XCN `hl7:"POS=12"`
	UniqueDocumentNumber EI `hl7:"POS=13"`
	ParentDocumentNumber EI `hl7:"POS=14"`
	PlacerOrderNumber []EI `hl7:"POS=15"`
	FillerOrderNumber EI `hl7:"POS=16"`
	UniqueDocumentFileName string `hl7:"POS=17"`
	DocumentCompletionStatus string `hl7:"POS=18"`
	DocumentConfidentialityStatus string `hl7:"POS=19"`
	DocumentAvailabilityStatus string `hl7:"POS=20"`
	DocumentStorageStatus string `hl7:"POS=21"`
	DocumentChangeReason string `hl7:"POS=22"`
	AuthenticationPersonTimeStamp []PPN `hl7:"POS=23"`
	DistributedCopiesCodeAndNameOfRecipients []XCN `hl7:"POS=24"`
}

