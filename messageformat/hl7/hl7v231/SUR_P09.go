package hl7v231

// SUR_P09_FACILITY - Group struct
type SUR_P09_FACILITY struct {
	FacilitySegment FAC `hl7:"TAG=FAC"`
	Product []SUR_P09_FACILITY_PRODUCT `hl7:"GROUP"`
	ProductSummaryHeaderSegment PSH `hl7:"TAG=PSH"`
	FacilityDetail []SUR_P09_FACILITY_FACILITY_DETAIL `hl7:"GROUP"`
}

// SUR_P09_FACILITY_PRODUCT - Group struct
type SUR_P09_FACILITY_PRODUCT struct {
	ProductSummaryHeaderSegment PSH `hl7:"TAG=PSH"`
	ProductDetailCountrySegment PDC `hl7:"TAG=PDC"`
}

// SUR_P09_FACILITY_FACILITY_DETAIL - Group struct
type SUR_P09_FACILITY_FACILITY_DETAIL struct {
	FacilitySegment FAC `hl7:"TAG=FAC"`
	ProductDetailCountrySegment PDC `hl7:"TAG=PDC"`
	NotesAndCommentsSegment NTE `hl7:"TAG=NTE"`
}

// SUR_P09 - Summary product experience report
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/SUR_P09
type SUR_P09 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Facility []SUR_P09_FACILITY `hl7:"GROUP"`
}

