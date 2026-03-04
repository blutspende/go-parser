package hl7v251

// FTS - File Trailer Segment
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/FTS
type FTS struct {
	FileBatchCount *int `hl7:"POS=2"`
	FileTrailerComment string `hl7:"POS=3"`
}

