package hl7v27

import "time"

// PDC - Product Detail Country
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/PDC
type PDC struct {
	ManufacturerDistributor []XON `hl7:"POS=2"`
	Country CWE `hl7:"POS=3"`
	BrandName string `hl7:"POS=4"`
	DeviceFamilyName string `hl7:"POS=5"`
	GenericName CWE `hl7:"POS=6"`
	ModelIdentifier []string `hl7:"POS=7"`
	CatalogueIdentifier string `hl7:"POS=8"`
	OtherIdentifier []string `hl7:"POS=9"`
	ProductCode CWE `hl7:"POS=10"`
	MarketingBasis string `hl7:"POS=11"`
	MarketingApprovalID string `hl7:"POS=12"`
	LabeledShelfLife CQ `hl7:"POS=13"`
	ExpectedShelfLife CQ `hl7:"POS=14"`
	DateFirstMarketed time.Time `hl7:"POS=15"`
	DateLastMarketed time.Time `hl7:"POS=16"`
}

