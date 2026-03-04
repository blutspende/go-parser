package hl7v26

import "time"

// RF1 - Referral Information
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/RF1
type RF1 struct {
	ReferralStatus CWE `hl7:"POS=2"`
	ReferralPriority CWE `hl7:"POS=3"`
	ReferralType CWE `hl7:"POS=4"`
	ReferralDisposition []CWE `hl7:"POS=5"`
	ReferralCategory CWE `hl7:"POS=6"`
	OriginatingReferralIdentifier EI `hl7:"POS=7"`
	EffectiveDate time.Time `hl7:"POS=8"`
	ExpirationDate time.Time `hl7:"POS=9"`
	ProcessDate time.Time `hl7:"POS=10"`
	ReferralReason []CWE `hl7:"POS=11"`
	ExternalReferralIdentifier []EI `hl7:"POS=12"`
	ReferralDocumentationCompletionStatus CWE `hl7:"POS=13"`
}

