package hl7v23

// MSA - Message acknowledgement segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MSA
type MSA struct {
	AcknowledgementCode string `hl7:"POS=2"`
	MessageControlID string `hl7:"POS=3"`
	TextMessage string `hl7:"POS=4"`
	ExpectedSequenceNumber *int `hl7:"POS=5"`
	DelayedAcknowledgementType string `hl7:"POS=6"`
	ErrorCondition CE `hl7:"POS=7"`
}

