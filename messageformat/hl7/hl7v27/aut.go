package hl7v27

import "time"

// AUT - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/AUT
type AUT struct {
	AuthorizingPayorPlanID CWE `hl7:"POS=2"`
	AuthorizingPayorCompanyID CWE `hl7:"POS=3"`
	AuthorizingPayorCompanyName string `hl7:"POS=4"`
	AuthorizationEffectiveDate time.Time `hl7:"POS=5"`
	AuthorizationExpirationDate time.Time `hl7:"POS=6"`
	AuthorizationIdentifier EI `hl7:"POS=7"`
	ReimbursementLimit CP `hl7:"POS=8"`
	RequestedNumberOfTreatments CQ `hl7:"POS=9"`
	AuthorizedNumberOfTreatments CQ `hl7:"POS=10"`
	ProcessDate time.Time `hl7:"POS=11"`
	RequestedDisciplines []CWE `hl7:"POS=12"`
	AuthorizedDisciplines []CWE `hl7:"POS=13"`
}

