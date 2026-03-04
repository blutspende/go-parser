package hl7v22

// ERR - Error
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/ERR
type ERR struct {
	ErrorCodeAndLocation []CM_ELD `hl7:"POS=2"`
}

