package hl7v271

import "time"

// CER - Certificate Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/CER
type CER struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SerialNumber string `hl7:"POS=3"`
	Version string `hl7:"POS=4"`
	GrantingAuthority XON `hl7:"POS=5"`
	IssuingAuthority XCN `hl7:"POS=6"`
	Signature ED `hl7:"POS=7"`
	GrantingCountry string `hl7:"POS=8"`
	GrantingStateProvince CWE `hl7:"POS=9"`
	GrantingCountyParish CWE `hl7:"POS=10"`
	CertificateType CWE `hl7:"POS=11"`
	CertificateDomain CWE `hl7:"POS=12"`
	SubjectID EI `hl7:"POS=13"`
	SubjectName string `hl7:"POS=14"`
	SubjectDirectoryAttributeExtension []CWE `hl7:"POS=15"`
	SubjectPublicKeyInfo CWE `hl7:"POS=16"`
	AuthorityKeyIdentifier CWE `hl7:"POS=17"`
	BasicConstraint string `hl7:"POS=18"`
	CrlDistributionPoint []CWE `hl7:"POS=19"`
	JurisdictionCountry string `hl7:"POS=20"`
	JurisdictionStateProvince CWE `hl7:"POS=21"`
	JurisdictionCountyParish CWE `hl7:"POS=22"`
	JurisdictionBreadth []CWE `hl7:"POS=23"`
	GrantingDate time.Time `hl7:"POS=24"`
	IssuingDate time.Time `hl7:"POS=25"`
	ActivationDate time.Time `hl7:"POS=26"`
	InactivationDate time.Time `hl7:"POS=27"`
	ExpirationDate time.Time `hl7:"POS=28"`
	RenewalDate time.Time `hl7:"POS=29"`
	RevocationDate time.Time `hl7:"POS=30"`
	RevocationReasonCode CWE `hl7:"POS=31"`
	CertificateStatusCode CWE `hl7:"POS=32"`
}

