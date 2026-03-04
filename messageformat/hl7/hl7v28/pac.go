package hl7v28

// PAC - Shipment Package
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/PAC
type PAC struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	PackageID EI `hl7:"POS=3"`
	ParentPackageID EI `hl7:"POS=4"`
	PositionInParentPackage NA `hl7:"POS=5"`
	PackageType CWE `hl7:"POS=6"`
	PackageCondition []CWE `hl7:"POS=7"`
	PackageHandlingCode []CWE `hl7:"POS=8"`
	PackageRiskCode []CWE `hl7:"POS=9"`
}

