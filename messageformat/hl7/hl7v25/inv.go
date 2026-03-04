package hl7v25

import "time"

// INV - Inventory Detail
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/INV
type INV struct {
	SubstanceIdentifier CE `hl7:"POS=2"`
	SubstanceStatus []CE `hl7:"POS=3"`
	SubstanceType CE `hl7:"POS=4"`
	InventoryContainerIdentifier CE `hl7:"POS=5"`
	ContainerCarrierIdentifier CE `hl7:"POS=6"`
	PositionOnCarrier CE `hl7:"POS=7"`
	InitialQuantity *int `hl7:"POS=8"`
	CurrentQuantity *int `hl7:"POS=9"`
	AvailableQuantity *int `hl7:"POS=10"`
	ConsumptionQuantity *int `hl7:"POS=11"`
	QuantityUnits CE `hl7:"POS=12"`
	ExpirationDateTime time.Time `hl7:"POS=13"`
	FirstUsedDateTime time.Time `hl7:"POS=14"`
	OnBoardStabilityDuration TQ `hl7:"POS=15"`
	TestFluidIdentifiers []CE `hl7:"POS=16"`
	ManufacturerLotNumber string `hl7:"POS=17"`
	ManufacturerIdentifier CE `hl7:"POS=18"`
	SupplierIdentifier CE `hl7:"POS=19"`
	OnBoardStabilityTime CQ `hl7:"POS=20"`
	TargetValue CQ `hl7:"POS=21"`
}

