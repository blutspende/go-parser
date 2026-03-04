package hl7v28

import "time"

// PMT - Payment Information
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/PMT
type PMT struct {
	PaymentRemittanceAdviceNumber EI `hl7:"POS=2"`
	PaymentRemittanceEffectiveDateTime time.Time `hl7:"POS=3"`
	PaymentRemittanceExpirationDateTime time.Time `hl7:"POS=4"`
	PaymentMethod CWE `hl7:"POS=5"`
	PaymentRemittanceDateTime time.Time `hl7:"POS=6"`
	PaymentRemittanceAmount CP `hl7:"POS=7"`
	CheckNumber EI `hl7:"POS=8"`
	PayeeBankIdentification XON `hl7:"POS=9"`
	PayeeTransitNumber string `hl7:"POS=10"`
	PayeeBankAccountID CX `hl7:"POS=11"`
	PaymentOrganization XON `hl7:"POS=12"`
	EsrCodeLine string `hl7:"POS=13"`
}

