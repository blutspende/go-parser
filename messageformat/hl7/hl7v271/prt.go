package hl7v271

import "time"

// PRT - Participation Information
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/PRT
type PRT struct {
	ParticipationInstanceID EI `hl7:"POS=2"`
	ActionCode string `hl7:"POS=3"`
	ActionReason CWE `hl7:"POS=4"`
	Participation CWE `hl7:"POS=5"`
	ParticipationPerson []XCN `hl7:"POS=6"`
	ParticipationPersonProviderType CWE `hl7:"POS=7"`
	ParticipantOrganizationUnitType CWE `hl7:"POS=8"`
	ParticipationOrganization []XON `hl7:"POS=9"`
	ParticipantLocation []PL `hl7:"POS=10"`
	ParticipationDevice []EI `hl7:"POS=11"`
	ParticipationBeginDateTimeArrivalTime time.Time `hl7:"POS=12"`
	ParticipationEndDateTimeDepartureTime time.Time `hl7:"POS=13"`
	ParticipationQualitativeDuration CWE `hl7:"POS=14"`
	ParticipationAddress []XAD `hl7:"POS=15"`
	ParticipantTelecommunicationAddress []XTN `hl7:"POS=16"`
}

