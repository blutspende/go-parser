package hl7v271

// PCE - Patient Charge Cost Center Exceptions
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/PCE
type PCE struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	CostCenterAccountNumber CX `hl7:"POS=3"`
	TransactionCode CWE `hl7:"POS=4"`
	TransactionAmountUnit CP `hl7:"POS=5"`
}

