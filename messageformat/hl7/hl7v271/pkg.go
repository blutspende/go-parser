package hl7v271

import "time"

// PKG - Item Packaging
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/PKG
type PKG struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	PackagingUnits CWE `hl7:"POS=3"`
	DefaultOrderUnitOfMeasureIndicator CNE `hl7:"POS=4"`
	PackageQuantity *int `hl7:"POS=5"`
	Price CP `hl7:"POS=6"`
	FutureItemPrice CP `hl7:"POS=7"`
	FutureItemPriceEffectiveDate time.Time `hl7:"POS=8"`
}

