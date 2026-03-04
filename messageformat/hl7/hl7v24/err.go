package hl7v24

// ERR - Error
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ERR
type ERR struct {
	ErrorCodeAndLocation []ELD `hl7:"POS=2"`
}

