package hl7v23

import "time"

// MSH - Message header segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MSH
type MSH struct {
	FieldSeparator                 string    `hl7:"1,fieldseparator" json:"fieldSeparator,omitempty"`
	EncodingCharacters             string    `hl7:"1,delimiter" json:"encodingCharacters,omitempty"`
	SendingApplication             HD        `hl7:"2" json:"sendingApplication,omitempty"`
	SendingFacility                HD        `hl7:"3" json:"sendingFacility,omitempty"`
	ReceivingApplication           HD        `hl7:"4" json:"receivingApplication,omitempty"`
	ReceivingFacility              HD        `hl7:"5" json:"receivingFacility,omitempty"`
	DateTimeOfMessage              time.Time `hl7:"6,longdate" json:"dateTimeOfMessage,omitempty"`
	Security                       string    `hl7:"7" json:"security,omitempty"`
	MessageType                    string    `hl7:"8.1" json:"messageType,omitempty"`
	MessageTriggerEvent            string    `hl7:"8.2" json:"messageTriggerEvent,omitempty"`
	MessageControlID               string    `hl7:"9" json:"messageControlID,omitempty"`
	ProcessingID                   string    `hl7:"10" json:"processingID,omitempty"`
	VersionID                      string    `hl7:"11" json:"versionID,omitempty"`
	SequenceNumber                 int       `hl7:"12" json:"sequenceNumber,omitempty"`
	ContinuationPointer            string    `hl7:"13" json:"continuationPointer,omitempty"`
	AcceptAcknowledgementType      string    `hl7:"14" json:"acceptAcknowledgementType,omitempty"`
	ApplicationAcknowledgementType string    `hl7:"15" json:"applicationAcknowledgementType,omitempty"`
	CountryCode                    string    `hl7:"16" json:"countryCode,omitempty"`
	CharacterSet                   []string  `hl7:"17" json:"characterSet,omitempty"`
	PrincipalLanguageOfMessage     CE        `hl7:"18" json:"principalLanguageOfMessage,omitempty"`
}
