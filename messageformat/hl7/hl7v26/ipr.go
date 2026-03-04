package hl7v26

import "time"

// IPR - Invoice Processing Results
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/IPR
type IPR struct {
	IprIdentifier EI `hl7:"POS=2"`
	ProviderCrossReferenceIdentifier EI `hl7:"POS=3"`
	PayerCrossReferenceIdentifier EI `hl7:"POS=4"`
	IprStatus CWE `hl7:"POS=5"`
	IprDateTime time.Time `hl7:"POS=6"`
	AdjudicatedPaidAmount CP `hl7:"POS=7"`
	ExpectedPaymentDateTime time.Time `hl7:"POS=8"`
	IprChecksum string `hl7:"POS=9"`
}

