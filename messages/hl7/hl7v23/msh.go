package hl7v23

import "time"

// MSH - Message header segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MSH
type MSH struct {
	SendingApplication             HD        `hl7:"POS=3" json:"sendingApplication,omitempty"`
	SendingFacility                HD        `hl7:"POS=4" json:"sendingFacility,omitempty"`
	ReceivingApplication           HD        `hl7:"POS=5" json:"receivingApplication,omitempty"`
	ReceivingFacility              HD        `hl7:"POS=6" json:"receivingFacility,omitempty"`
	DateTimeOfMessage              time.Time `hl7:"POS=7" json:"dateTimeOfMessage,omitempty"`
	Security                       string    `hl7:"POS=8" json:"security,omitempty"`
	MessageType                    string    `hl7:"POS=9.1" json:"messageType,omitempty"`
	MessageTriggerEvent            string    `hl7:"POS=9.2" json:"messageTriggerEvent,omitempty"`
	MessageControlID               string    `hl7:"POS=10" json:"messageControlID,omitempty"`
	ProcessingID                   string    `hl7:"POS=11" json:"processingID,omitempty"`
	VersionID                      string    `hl7:"POS=12" json:"versionID,omitempty"`
	SequenceNumber                 *int      `hl7:"POS=13" json:"sequenceNumber,omitempty"`
	ContinuationPointer            string    `hl7:"POS=14" json:"continuationPointer,omitempty"`
	AcceptAcknowledgementType      string    `hl7:"POS=15" json:"acceptAcknowledgementType,omitempty"`
	ApplicationAcknowledgementType string    `hl7:"POS=16" json:"applicationAcknowledgementType,omitempty"`
	CountryCode                    string    `hl7:"POS=17" json:"countryCode,omitempty"`
	CharacterSet                   []string  `hl7:"POS=18" json:"characterSet,omitempty"`
	PrincipalLanguageOfMessage     CE        `hl7:"POS=19" json:"principalLanguageOfMessage,omitempty"`
}
