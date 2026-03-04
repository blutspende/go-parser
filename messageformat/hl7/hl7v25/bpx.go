package hl7v25

import "time"

// BPX - Blood product dispense status
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/BPX
type BPX struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	BpDispenseStatus CWE `hl7:"POS=3"`
	BpStatus string `hl7:"POS=4"`
	BpDateTimeOfStatus time.Time `hl7:"POS=5"`
	BcDonationID EI `hl7:"POS=6"`
	BcComponent CNE `hl7:"POS=7"`
	BcDonationTypeIntendedUse CNE `hl7:"POS=8"`
	CpCommercialProduct CWE `hl7:"POS=9"`
	CpManufacturer XON `hl7:"POS=10"`
	CpLotNumber EI `hl7:"POS=11"`
	BpBloodGroup CNE `hl7:"POS=12"`
	BcSpecialTesting []CNE `hl7:"POS=13"`
	BpExpirationDateTime time.Time `hl7:"POS=14"`
	BpQuantity *int `hl7:"POS=15"`
	BpAmount *int `hl7:"POS=16"`
	BpUnits CE `hl7:"POS=17"`
	BpUniqueID EI `hl7:"POS=18"`
	BpActualDispensedToLocation PL `hl7:"POS=19"`
	BpActualDispensedToAddress XAD `hl7:"POS=20"`
	BpDispensedToReceiver XCN `hl7:"POS=21"`
	BpDispensingIndividual XCN `hl7:"POS=22"`
}

