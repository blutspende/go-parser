package hl7v231

// MSA - Message acknowledgment segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/MSA
type MSA struct {
	AcknowledgementCode string `hl7:"POS=2"`
	MessageControlID string `hl7:"POS=3"`
	TextMessage string `hl7:"POS=4"`
	ExpectedSequenceNumber *int `hl7:"POS=5"`
	DelayedAcknowledgmentType string `hl7:"POS=6"`
	ErrorCondition CE `hl7:"POS=7"`
}

