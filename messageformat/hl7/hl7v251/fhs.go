package hl7v251

import "time"

// FHS - File Header Segment
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/FHS
type FHS struct {
	FileFieldSeparator string `hl7:"POS=2"`
	FileEncodingCharacters string `hl7:"POS=3"`
	FileSendingApplication HD `hl7:"POS=4"`
	FileSendingFacility HD `hl7:"POS=5"`
	FileReceivingApplication HD `hl7:"POS=6"`
	FileReceivingFacility HD `hl7:"POS=7"`
	FileCreationDateTime time.Time `hl7:"POS=8"`
	FileSecurity string `hl7:"POS=9"`
	FileNameID string `hl7:"POS=10"`
	FileHeaderComment string `hl7:"POS=11"`
	FileControlID string `hl7:"POS=12"`
	ReferenceFileControlID string `hl7:"POS=13"`
}

