package hl7v24

import "time"

// HL7 v2.4 - MSH - Message Header
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSH
type MSH struct {
	FieldSeparator string `hl7:"1,fieldseparator" json:"FieldSeparator,omitempty"`
	// the fieldseparater is regarded as field#1 which doesnt work out with splitting the fields
	// therefore the fields are all offest -1 to the documentation
	EncodingCharacters                  string    `hl7:"1,delimiter" json:"EncodingCharacters,omitempty"`
	SendingApplication                  HD        `hl7:"2" json:"SendingApplication,omitempty"`
	SendingFacility                     HD        `hl7:"3" json:"SendingFacility,omitempty"`
	ReceivingApplication                HD        `hl7:"4" json:"ReceivingApplication,omitempty"`
	ReceivingFacility                   HD        `hl7:"5" json:"ReceivingFacility,omitempty"`
	DateTimeOfMessage                   time.Time `hl7:"6,longdate" json:"DateTimeOfMessage,omitempty"`
	Security                            string    `hl7:"7" json:"Security,omitempty"`
	MessageType                         string    `hl7:"8.1" json:"MessageType,omitempty"`
	MessageTriggerEvent                 string    `hl7:"8.2" json:"MessageTriggerEvent,omitempty"`
	MessageControlID                    string    `hl7:"9" json:"MessageControlID,omitempty"`
	ProccessingID                       string    `hl7:"10" json:"ProccessingID,omitempty"`
	VersionID                           string    `hl7:"11" json:"VersionID,omitempty"`
	SequenceNumber                      int       `hl7:"12" json:"SequenceNumber,omitempty"`
	ContinuationPointer                 string    `hl7:"13" json:"ContinuationPointer,omitempty"`
	AcceptAcknowledgementType           string    `hl7:"14" json:"AcceptAcknowledgementType,omitempty"`
	ApplicationAcknowledgementType      string    `hl7:"15" json:"ApplicationAcknowledgementType,omitempty"`
	CountryCode                         string    `hl7:"16" json:"CountryCode,omitempty"`
	CharacterSet                        []string  `hl7:"17" json:"CharacterSet,omitempty"`
	PrincipalLanguageOfMessage          CE        `hl7:"18" json:"PrincipalLanguageOfMessage,omitempty"`
	AlternateCharacterSetHandlingScheme string    `hl7:"19" json:"AlternateCharacterSetHandlingScheme,omitempty"`
	ConformanceStatementID              string    `hl7:"20" json:"ConformanceStatementID,omitempty"`
}
