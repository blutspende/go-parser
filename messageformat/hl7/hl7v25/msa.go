package hl7v25

// MSA - Message Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/MSA
type MSA struct {
	AcknowledgmentCode string `hl7:"POS=2"`
	MessageControlID string `hl7:"POS=3"`
	TextMessage string `hl7:"POS=4"`
	ExpectedSequenceNumber *int `hl7:"POS=5"`
	ErrorCondition CE `hl7:"POS=7"`
}

