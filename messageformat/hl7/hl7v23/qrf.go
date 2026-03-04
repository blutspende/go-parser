package hl7v23

import "time"

// QRF - Query filter segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/QRF
type QRF struct {
	WhereSubjectFilter []string `hl7:"POS=2"`
	WhenDataStartDateTime time.Time `hl7:"POS=3"`
	WhenDataEndDateTime time.Time `hl7:"POS=4"`
	WhatUserQualifier []string `hl7:"POS=5"`
	OtherQrySubjectFilter []string `hl7:"POS=6"`
	WhichDateTimeQualifier []string `hl7:"POS=7"`
	WhichDateTimeStatusQualifier []string `hl7:"POS=8"`
	DateTimeSelectionQualifier []string `hl7:"POS=9"`
	WhenQuantityTimingQualifier TQ `hl7:"POS=10"`
}

