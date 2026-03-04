package hl7v26

import "time"

// PSH - Product Summary Header
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/PSH
type PSH struct {
	ReportType string `hl7:"POS=2"`
	ReportFormIdentifier string `hl7:"POS=3"`
	ReportDate time.Time `hl7:"POS=4"`
	ReportIntervalStartDate time.Time `hl7:"POS=5"`
	ReportIntervalEndDate time.Time `hl7:"POS=6"`
	QuantityManufactured CQ `hl7:"POS=7"`
	QuantityDistributed CQ `hl7:"POS=8"`
	QuantityDistributedMethod string `hl7:"POS=9"`
	QuantityDistributedComment string `hl7:"POS=10"`
	QuantityInUse CQ `hl7:"POS=11"`
	QuantityInUseMethod string `hl7:"POS=12"`
	QuantityInUseComment string `hl7:"POS=13"`
	NumberOfProductExperienceReportsFiledByFacility []*int `hl7:"POS=14"`
	NumberOfProductExperienceReportsFiledByDistributor []*int `hl7:"POS=15"`
}

