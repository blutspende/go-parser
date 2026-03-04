package hl7v25

import "time"

// BTX - Blood Product Transfusion/Disposition
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/BTX
type BTX struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	BcDonationID EI `hl7:"POS=3"`
	BcComponent CNE `hl7:"POS=4"`
	BcBloodGroup CNE `hl7:"POS=5"`
	CpCommercialProduct CWE `hl7:"POS=6"`
	CpManufacturer XON `hl7:"POS=7"`
	CpLotNumber EI `hl7:"POS=8"`
	BpQuantity *int `hl7:"POS=9"`
	BpAmount *int `hl7:"POS=10"`
	BpUnits CE `hl7:"POS=11"`
	BpTransfusionDispositionStatus CWE `hl7:"POS=12"`
	BpMessageStatus string `hl7:"POS=13"`
	BpDateTimeOfStatus time.Time `hl7:"POS=14"`
	BpAdministrator XCN `hl7:"POS=15"`
	BpVerifier XCN `hl7:"POS=16"`
	BpTransfusionStartDateTimeOfStatus time.Time `hl7:"POS=17"`
	BpTransfusionEndDateTimeOfStatus time.Time `hl7:"POS=18"`
	BpAdverseReactionType []CWE `hl7:"POS=19"`
	BpTransfusionInterruptedReason CWE `hl7:"POS=20"`
}

