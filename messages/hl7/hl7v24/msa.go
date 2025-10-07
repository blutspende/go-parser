package hl7v24

// HL7 v2.4 - MSA - Message Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSA
type MSA struct {
	AcknowledgementCode       string `hl7:"1" json:"AcknowledgementCode,omitempty"`
	MessageControlID          string `hl7:"2" json:"MessageControlID,omitempty"`
	TextMessage               string `hl7:"3" json:"TextMessage,omitempty"`
	ExpectedSequenceNumber    int    `hl7:"4" json:"ExpectedSequenceNumber,omitempty"`
	DelayedAcknowledgmentType string `hl7:"5" json:"DelayedAcknowledgmentType,omitempty"`
	ErrorCondition            CE     `hl7:"6" json:"ErrorCondition,omitempty"`
}
