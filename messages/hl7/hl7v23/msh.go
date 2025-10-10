package hl7v23

import "time"

// MSH - Message header segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MSH
type MSH struct {
	//FieldSeparator                 string    `hl7:"1,fieldseparator" json:"fieldSeparator,omitempty"`
	//EncodingCharacters             string    `hl7:"1,delimiter" json:"encodingCharacters,omitempty"`
	SendingApplication             HD        `hl7:"3" json:"sendingApplication,omitempty"`
	SendingFacility                HD        `hl7:"4" json:"sendingFacility,omitempty"`
	ReceivingApplication           HD        `hl7:"5" json:"receivingApplication,omitempty"`
	ReceivingFacility              HD        `hl7:"6" json:"receivingFacility,omitempty"`
	DateTimeOfMessage              time.Time `hl7:"7,longdate" json:"dateTimeOfMessage,omitempty"`
	Security                       string    `hl7:"8" json:"security,omitempty"`
	MessageType                    string    `hl7:"9.1" json:"messageType,omitempty"`
	MessageTriggerEvent            string    `hl7:"9.2" json:"messageTriggerEvent,omitempty"`
	MessageControlID               string    `hl7:"10" json:"messageControlID,omitempty"`
	ProcessingID                   string    `hl7:"11" json:"processingID,omitempty"`
	VersionID                      string    `hl7:"12" json:"versionID,omitempty"`
	SequenceNumber                 int       `hl7:"13" json:"sequenceNumber,omitempty"`
	ContinuationPointer            string    `hl7:"14" json:"continuationPointer,omitempty"`
	AcceptAcknowledgementType      string    `hl7:"15" json:"acceptAcknowledgementType,omitempty"`
	ApplicationAcknowledgementType string    `hl7:"16" json:"applicationAcknowledgementType,omitempty"`
	CountryCode                    string    `hl7:"17" json:"countryCode,omitempty"`
	CharacterSet                   []string  `hl7:"18" json:"characterSet,omitempty"`
	PrincipalLanguageOfMessage     CE        `hl7:"19" json:"principalLanguageOfMessage,omitempty"`
}
