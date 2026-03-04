package hl7v23

// FTS - File Trailer Segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/FTS
type FTS struct {
	FileBatchCount *int `hl7:"POS=2"`
	FileTrailerComment string `hl7:"POS=3"`
}

