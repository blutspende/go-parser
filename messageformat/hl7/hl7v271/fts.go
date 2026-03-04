package hl7v271

// FTS - File Trailer
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/FTS
type FTS struct {
	FileBatchCount *int `hl7:"POS=2"`
	FileTrailerComment string `hl7:"POS=3"`
}

