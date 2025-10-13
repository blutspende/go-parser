package hl7v24

import "time"

// HL7 v2.4 - MSH - Message Header
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSH
type MSH struct {
	SendingApplication                  HD        `hl7:"3" json:"SendingApplication,omitempty"`
	SendingFacility                     HD        `hl7:"4" json:"SendingFacility,omitempty"`
	ReceivingApplication                HD        `hl7:"5" json:"ReceivingApplication,omitempty"`
	ReceivingFacility                   HD        `hl7:"6" json:"ReceivingFacility,omitempty"`
	DateTimeOfMessage                   time.Time `hl7:"7,longdate" json:"DateTimeOfMessage,omitempty"`
	Security                            string    `hl7:"8" json:"Security,omitempty"`
	MessageType                         string    `hl7:"9.1" json:"MessageType,omitempty"`
	MessageTriggerEvent                 string    `hl7:"9.2" json:"MessageTriggerEvent,omitempty"`
	MessageControlID                    string    `hl7:"10" json:"MessageControlID,omitempty"`
	ProccessingID                       string    `hl7:"11" json:"ProccessingID,omitempty"`
	VersionID                           string    `hl7:"12" json:"VersionID,omitempty"`
	SequenceNumber                      int       `hl7:"13" json:"SequenceNumber,omitempty"`
	ContinuationPointer                 string    `hl7:"14" json:"ContinuationPointer,omitempty"`
	AcceptAcknowledgementType           string    `hl7:"15" json:"AcceptAcknowledgementType,omitempty"`
	ApplicationAcknowledgementType      string    `hl7:"16" json:"ApplicationAcknowledgementType,omitempty"`
	CountryCode                         string    `hl7:"17" json:"CountryCode,omitempty"`
	CharacterSet                        []string  `hl7:"18" json:"CharacterSet,omitempty"`
	PrincipalLanguageOfMessage          CE        `hl7:"19" json:"PrincipalLanguageOfMessage,omitempty"`
	AlternateCharacterSetHandlingScheme string    `hl7:"20" json:"AlternateCharacterSetHandlingScheme,omitempty"`
	ConformanceStatementID              string    `hl7:"21" json:"ConformanceStatementID,omitempty"`
}
