package hl7v27

// PSG - Product/service Group
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/PSG
type PSG struct {
	ProviderProductServiceGroupNumber EI `hl7:"POS=2"`
	PayerProductServiceGroupNumber EI `hl7:"POS=3"`
	ProductServiceGroupSequenceNumber int `hl7:"POS=4;ATR=sequence"`
	AdjudicateAsGroup string `hl7:"POS=5"`
	ProductServiceGroupBilledAmount CP `hl7:"POS=6"`
	ProductServiceGroupDescription string `hl7:"POS=7"`
}

