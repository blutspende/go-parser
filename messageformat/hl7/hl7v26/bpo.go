package hl7v26

import "time"

// BPO - Blood product order
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/BPO
type BPO struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	BpUniversalServiceIdentifier CWE `hl7:"POS=3"`
	BpProcessingRequirements []CWE `hl7:"POS=4"`
	BpQuantity *int `hl7:"POS=5"`
	BpAmount *int `hl7:"POS=6"`
	BpUnits CWE `hl7:"POS=7"`
	BpIntendedUseDateTime time.Time `hl7:"POS=8"`
	BpIntendedDispenseFromLocation PL `hl7:"POS=9"`
	BpIntendedDispenseFromAddress XAD `hl7:"POS=10"`
	BpRequestedDispenseDateTime time.Time `hl7:"POS=11"`
	BpRequestedDispenseToLocation PL `hl7:"POS=12"`
	BpRequestedDispenseToAddress XAD `hl7:"POS=13"`
	BpIndicationForUse []CWE `hl7:"POS=14"`
	BpInformedConsentIndicator string `hl7:"POS=15"`
}

