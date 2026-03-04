package hl7v28

import "time"

// SHP - Shipment
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/SHP
type SHP struct {
	ShipmentID EI `hl7:"POS=2"`
	InternalShipmentID []EI `hl7:"POS=3"`
	ShipmentStatus CWE `hl7:"POS=4"`
	ShipmentStatusDateTime time.Time `hl7:"POS=5"`
	ShipmentStatusReason string `hl7:"POS=6"`
	ShipmentPriority CWE `hl7:"POS=7"`
	ShipmentConfidentiality []CWE `hl7:"POS=8"`
	NumberOfPackagesInShipment *int `hl7:"POS=9"`
	ShipmentCondition []CWE `hl7:"POS=10"`
	ShipmentHandlingCode []CWE `hl7:"POS=11"`
	ShipmentRiskCode []CWE `hl7:"POS=12"`
}

