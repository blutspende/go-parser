package hl7v231

// ERR - Error segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/ERR
type ERR struct {
	ErrorCodeAndLocation []ELD `hl7:"POS=2"`
}

