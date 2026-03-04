package hl7v26

// PSS - Product/Service Section
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/PSS
type PSS struct {
	ProviderProductServiceSectionNumber EI `hl7:"POS=2"`
	PayerProductServiceSectionNumber EI `hl7:"POS=3"`
	ProductServiceSectionSequenceNumber int `hl7:"POS=4;ATR=sequence"`
	BilledAmount CP `hl7:"POS=5"`
	SectionDescriptionOrHeading string `hl7:"POS=6"`
}

