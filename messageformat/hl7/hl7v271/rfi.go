package hl7v271

import "time"

// RFI - Request For Information
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/RFI
type RFI struct {
	RequestDate time.Time `hl7:"POS=2"`
	ResponseDueDate time.Time `hl7:"POS=3"`
	PatientConsent string `hl7:"POS=4"`
	DateAdditionalInformationWasSubmitted time.Time `hl7:"POS=5"`
}

