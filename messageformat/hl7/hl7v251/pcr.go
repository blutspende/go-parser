package hl7v251

import "time"

// PCR - Possible Causal Relationship
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/PCR
type PCR struct {
	ImplicatedProduct CE `hl7:"POS=2"`
	GenericProduct string `hl7:"POS=3"`
	ProductClass CE `hl7:"POS=4"`
	TotalDurationOfTherapy CQ `hl7:"POS=5"`
	ProductManufactureDate time.Time `hl7:"POS=6"`
	ProductExpirationDate time.Time `hl7:"POS=7"`
	ProductImplantationDate time.Time `hl7:"POS=8"`
	ProductExplantationDate time.Time `hl7:"POS=9"`
	SingleUseDevice string `hl7:"POS=10"`
	IndicationForProductUse CE `hl7:"POS=11"`
	ProductProblem string `hl7:"POS=12"`
	ProductSerialLotNumber []string `hl7:"POS=13"`
	ProductAvailableForInspection string `hl7:"POS=14"`
	ProductEvaluationPerformed CE `hl7:"POS=15"`
	ProductEvaluationStatus CE `hl7:"POS=16"`
	ProductEvaluationResults CE `hl7:"POS=17"`
	EvaluatedProductSource string `hl7:"POS=18"`
	DateProductReturnedToManufacturer time.Time `hl7:"POS=19"`
	DeviceOperatorQualifications string `hl7:"POS=20"`
	RelatednessAssessment string `hl7:"POS=21"`
	ActionTakenInResponseToTheEvent []string `hl7:"POS=22"`
	EventCausalityObservations []string `hl7:"POS=23"`
	IndirectExposureMechanism []string `hl7:"POS=24"`
}

