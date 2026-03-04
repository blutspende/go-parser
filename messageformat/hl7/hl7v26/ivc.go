package hl7v26

import "time"

// IVC - Invoice
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/IVC
type IVC struct {
	ProviderInvoiceNumber EI `hl7:"POS=2"`
	PayerInvoiceNumber EI `hl7:"POS=3"`
	ContractAgreementNumber EI `hl7:"POS=4"`
	InvoiceControl string `hl7:"POS=5"`
	InvoiceReason string `hl7:"POS=6"`
	InvoiceType string `hl7:"POS=7"`
	InvoiceDateTime time.Time `hl7:"POS=8"`
	InvoiceAmount CP `hl7:"POS=9"`
	PaymentTerms string `hl7:"POS=10"`
	ProviderOrganization XON `hl7:"POS=11"`
	PayerOrganization XON `hl7:"POS=12"`
	Attention XCN `hl7:"POS=13"`
	LastInvoiceIndicator string `hl7:"POS=14"`
	InvoiceBookingPeriod time.Time `hl7:"POS=15"`
	Origin string `hl7:"POS=16"`
	InvoiceFixedAmount CP `hl7:"POS=17"`
	SpecialCosts CP `hl7:"POS=18"`
	AmountForDoctorsTreatment CP `hl7:"POS=19"`
	ResponsiblePhysician XCN `hl7:"POS=20"`
	CostCenter CX `hl7:"POS=21"`
	InvoicePrepaidAmount CP `hl7:"POS=22"`
	TotalInvoiceAmountWithoutPrepaidAmount CP `hl7:"POS=23"`
	TotalAmountOfVat CP `hl7:"POS=24"`
	VatRatesApplied []*int `hl7:"POS=25"`
	BenefitGroup string `hl7:"POS=26"`
	ProviderTaxID string `hl7:"POS=27"`
	PayerTaxID string `hl7:"POS=28"`
	ProviderTaxStatus string `hl7:"POS=29"`
	PayerTaxStatus string `hl7:"POS=30"`
	SalesTaxID string `hl7:"POS=31"`
}

