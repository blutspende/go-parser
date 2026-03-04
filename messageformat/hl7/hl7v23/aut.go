package hl7v23

import "time"

// AUT - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/AUT
type AUT struct {
	AuthorizingPayorPlanCode CE `hl7:"POS=2"`
	AuthorizingPayorCompanyID CE `hl7:"POS=3"`
	AuthorizingPayorCompanyName string `hl7:"POS=4"`
	AuthorizationEffectiveDate time.Time `hl7:"POS=5"`
	AuthorizationExpirationDate time.Time `hl7:"POS=6"`
	AuthorizationIdentifier EI `hl7:"POS=7"`
	ReimbursementLimit CP `hl7:"POS=8"`
	RequestedNumberOfTreatments *int `hl7:"POS=9"`
	AuthorizedNumberOfTreatments *int `hl7:"POS=10"`
	ProcessDate time.Time `hl7:"POS=11"`
}

