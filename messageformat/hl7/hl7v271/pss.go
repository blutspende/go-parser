package hl7v271

// PSS - Product/service Section
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/PSS
type PSS struct {
	ProviderProductServiceSectionNumber EI `hl7:"POS=2"`
	PayerProductServiceSectionNumber EI `hl7:"POS=3"`
	ProductServiceSectionSequenceNumber int `hl7:"POS=4;ATR=sequence"`
	BilledAmount CP `hl7:"POS=5"`
	SectionDescriptionOrHeading string `hl7:"POS=6"`
}

