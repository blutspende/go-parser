package hl7v27

import "time"

// INV - Inventory Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/INV
type INV struct {
	SubstanceIdentifier CWE `hl7:"POS=2"`
	SubstanceStatus []CWE `hl7:"POS=3"`
	SubstanceType CWE `hl7:"POS=4"`
	InventoryContainerIdentifier CWE `hl7:"POS=5"`
	ContainerCarrierIdentifier CWE `hl7:"POS=6"`
	PositionOnCarrier CWE `hl7:"POS=7"`
	InitialQuantity *int `hl7:"POS=8"`
	CurrentQuantity *int `hl7:"POS=9"`
	AvailableQuantity *int `hl7:"POS=10"`
	ConsumptionQuantity *int `hl7:"POS=11"`
	QuantityUnits CWE `hl7:"POS=12"`
	ExpirationDateTime time.Time `hl7:"POS=13"`
	FirstUsedDateTime time.Time `hl7:"POS=14"`
	TestFluidIdentifiers []CWE `hl7:"POS=16"`
	ManufacturerLotNumber string `hl7:"POS=17"`
	ManufacturerIdentifier CWE `hl7:"POS=18"`
	SupplierIdentifier CWE `hl7:"POS=19"`
	OnBoardStabilityTime CQ `hl7:"POS=20"`
	TargetValue CQ `hl7:"POS=21"`
}

