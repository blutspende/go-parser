package hl7v23

// ERR - Error segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ERR
type ERR struct {
	ErrorCodeAndLocation []CM_ELD `hl7:"POS=2"`
}

