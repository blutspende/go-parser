package hl7v27

// MSA - Message Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/MSA
type MSA struct {
	AcknowledgmentCode string `hl7:"POS=2"`
	MessageControlID string `hl7:"POS=3"`
	ExpectedSequenceNumber *int `hl7:"POS=5"`
	MessageWaitingNumber *int `hl7:"POS=8"`
	MessageWaitingPriority string `hl7:"POS=9"`
}

