package hl7v27

import "time"

// SPM - Specimen
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/SPM
type SPM struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SpecimenID EIP `hl7:"POS=3"`
	SpecimenParentIds []EIP `hl7:"POS=4"`
	SpecimenType CWE `hl7:"POS=5"`
	SpecimenTypeModifier []CWE `hl7:"POS=6"`
	SpecimenAdditives []CWE `hl7:"POS=7"`
	SpecimenCollectionMethod CWE `hl7:"POS=8"`
	SpecimenSourceSite CWE `hl7:"POS=9"`
	SpecimenSourceSiteModifier []CWE `hl7:"POS=10"`
	SpecimenCollectionSite CWE `hl7:"POS=11"`
	SpecimenRole []CWE `hl7:"POS=12"`
	SpecimenCollectionAmount CQ `hl7:"POS=13"`
	GroupedSpecimenCount *int `hl7:"POS=14"`
	SpecimenDescription []string `hl7:"POS=15"`
	SpecimenHandlingCode []CWE `hl7:"POS=16"`
	SpecimenRiskCode []CWE `hl7:"POS=17"`
	SpecimenCollectionDateTime DR `hl7:"POS=18"`
	SpecimenReceivedDateTime time.Time `hl7:"POS=19"`
	SpecimenExpirationDateTime time.Time `hl7:"POS=20"`
	SpecimenAvailability string `hl7:"POS=21"`
	SpecimenRejectReason []CWE `hl7:"POS=22"`
	SpecimenQuality CWE `hl7:"POS=23"`
	SpecimenAppropriateness CWE `hl7:"POS=24"`
	SpecimenCondition []CWE `hl7:"POS=25"`
	SpecimenCurrentQuantity CQ `hl7:"POS=26"`
	NumberOfSpecimenContainers *int `hl7:"POS=27"`
	ContainerType CWE `hl7:"POS=28"`
	ContainerCondition CWE `hl7:"POS=29"`
	SpecimenChildRole CWE `hl7:"POS=30"`
	AccessionID []CX `hl7:"POS=31"`
	OtherSpecimenID []CX `hl7:"POS=32"`
	ShipmentID EI `hl7:"POS=33"`
}

