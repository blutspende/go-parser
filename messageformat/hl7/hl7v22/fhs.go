package hl7v22

import "time"

// FHS - File Header
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/FHS
type FHS struct {
	FileFieldSeparator string `hl7:"POS=2"`
	FileEncodingCharacters string `hl7:"POS=3"`
	FileSendingApplication string `hl7:"POS=4"`
	FileSendingFacility string `hl7:"POS=5"`
	FileReceivingApplication string `hl7:"POS=6"`
	FileReceivingFacility string `hl7:"POS=7"`
	FileCreationDateTime time.Time `hl7:"POS=8"`
	FileSecurity string `hl7:"POS=9"`
	FileNameID string `hl7:"POS=10"`
	FileHeaderComment string `hl7:"POS=11"`
	FileControlID string `hl7:"POS=12"`
	ReferenceFileControlID string `hl7:"POS=13"`
}

