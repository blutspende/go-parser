package hl7v25

import "time"

// QRF - Original style query filter
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/QRF
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
	SearchConfidenceThreshold *int `hl7:"POS=11"`
}

