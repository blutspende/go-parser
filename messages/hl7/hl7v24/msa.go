package hl7v24

// HL7 v2.4 - MSA - Message Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSA
type MSA struct {
	AcknowledgementCode       string `hl7:"POS=2" json:"AcknowledgementCode,omitempty"`
	MessageControlID          string `hl7:"POS=3" json:"MessageControlID,omitempty"`
	TextMessage               string `hl7:"POS=4" json:"TextMessage,omitempty"`
	ExpectedSequenceNumber    *int   `hl7:"POS=5" json:"ExpectedSequenceNumber,omitempty"`
	DelayedAcknowledgmentType string `hl7:"POS=6" json:"DelayedAcknowledgmentType,omitempty"`
	ErrorCondition            CE     `hl7:"POS=7" json:"ErrorCondition,omitempty"`
}
