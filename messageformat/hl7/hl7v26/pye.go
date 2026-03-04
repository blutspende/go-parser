package hl7v26

// PYE - Payee Information
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/PYE
type PYE struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	PayeeType string `hl7:"POS=3"`
	PayeeRelationshipToInvoicePatient string `hl7:"POS=4"`
	PayeeIdentificationList []XON `hl7:"POS=5"`
	PayeePersonName []XPN `hl7:"POS=6"`
	PayeeAddress []XAD `hl7:"POS=7"`
	PaymentMethod string `hl7:"POS=8"`
}

