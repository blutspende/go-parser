package hl7v26

// MSA - Message Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/MSA
type MSA struct {
	AcknowledgmentCode string `hl7:"POS=2"`
	MessageControlID string `hl7:"POS=3"`
	TextMessage string `hl7:"POS=4"`
	ExpectedSequenceNumber *int `hl7:"POS=5"`
	ErrorCondition CE `hl7:"POS=7"`
	MessageWaitingNumber *int `hl7:"POS=8"`
	MessageWaitingPriority string `hl7:"POS=9"`
}

