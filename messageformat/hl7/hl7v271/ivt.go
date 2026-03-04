package hl7v271

// IVT - Material Location
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/IVT
type IVT struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	InventoryLocationIdentifier EI `hl7:"POS=3"`
	InventoryLocationName string `hl7:"POS=4"`
	SourceLocationIdentifier EI `hl7:"POS=5"`
	SourceLocationName string `hl7:"POS=6"`
	ItemStatus CWE `hl7:"POS=7"`
	BinLocationIdentifier []EI `hl7:"POS=8"`
	OrderPackaging CWE `hl7:"POS=9"`
	IssuePackaging CWE `hl7:"POS=10"`
	DefaultInventoryAssetAccount EI `hl7:"POS=11"`
	PatientChargeableIndicator CNE `hl7:"POS=12"`
	TransactionCode CWE `hl7:"POS=13"`
	TransactionAmountUnit CP `hl7:"POS=14"`
	ItemImportanceCode CWE `hl7:"POS=15"`
	StockedItemIndicator CNE `hl7:"POS=16"`
	ConsignmentItemIndicator CNE `hl7:"POS=17"`
	ReusableItemIndicator CNE `hl7:"POS=18"`
	ReusableCost CP `hl7:"POS=19"`
	SubstituteItemIdentifier []EI `hl7:"POS=20"`
	LatexFreeSubstituteItemIdentifier EI `hl7:"POS=21"`
	RecommendedReorderTheory CWE `hl7:"POS=22"`
	RecommendedSafetyStockDays *int `hl7:"POS=23"`
	RecommendedMaximumDaysInventory *int `hl7:"POS=24"`
	RecommendedOrderPoint *int `hl7:"POS=25"`
	RecommendedOrderAmount *int `hl7:"POS=26"`
	OperatingRoomParLevelIndicator CNE `hl7:"POS=27"`
}

