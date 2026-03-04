package hl7v251

import "time"

// RF1 - Referral Information
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/RF1
type RF1 struct {
	ReferralStatus CE `hl7:"POS=2"`
	ReferralPriority CE `hl7:"POS=3"`
	ReferralType CE `hl7:"POS=4"`
	ReferralDisposition []CE `hl7:"POS=5"`
	ReferralCategory CE `hl7:"POS=6"`
	OriginatingReferralIdentifier EI `hl7:"POS=7"`
	EffectiveDate time.Time `hl7:"POS=8"`
	ExpirationDate time.Time `hl7:"POS=9"`
	ProcessDate time.Time `hl7:"POS=10"`
	ReferralReason []CE `hl7:"POS=11"`
	ExternalReferralIdentifier []EI `hl7:"POS=12"`
}

