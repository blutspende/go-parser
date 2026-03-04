package hl7v271

// ITM - Material Item
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/ITM
type ITM struct {
	ItemIdentifier                        EI     `hl7:"POS=2"`
	ItemDescription                       string `hl7:"POS=3"`
	ItemStatus                            CWE    `hl7:"POS=4"`
	ItemType                              CWE    `hl7:"POS=5"`
	ItemCategory                          CWE    `hl7:"POS=6"`
	SubjectToExpirationIndicator          CNE    `hl7:"POS=7"`
	ManufacturerIdentifier                EI     `hl7:"POS=8"`
	ManufacturerName                      string `hl7:"POS=9"`
	ManufacturerCatalogNumber             string `hl7:"POS=10"`
	ManufacturerLabelerIdentificationCode CWE    `hl7:"POS=11"`
	PatientChargeableIndicator            CNE    `hl7:"POS=12"`
	TransactionCode                       CWE    `hl7:"POS=13"`
	TransactionAmountUnit                 CP     `hl7:"POS=14"`
	StockedItemIndicator                  CNE    `hl7:"POS=15"`
	SupplyRiskCodes                       CWE    `hl7:"POS=16"`
	ApprovingRegulatoryAgency             []XON  `hl7:"POS=17"`
	LatexIndicator                        CNE    `hl7:"POS=18"`
	RulingAct                             []CWE  `hl7:"POS=19"`
	ItemNaturalAccountCode                CWE    `hl7:"POS=20"`
	ApprovedToBuyQuantity                 *int   `hl7:"POS=21"`
	ApprovedToBuyPrice                    MO     `hl7:"POS=22"`
	TaxableItemIndicator                  CNE    `hl7:"POS=23"`
	FreightChargeIndicator                CNE    `hl7:"POS=24"`
	ItemSetIndicator                      CNE    `hl7:"POS=25"`
	ItemSetID                             EI     `hl7:"POS=26"`
	TrackDepartmentUsageIndicator         CNE    `hl7:"POS=27"`
	ProcedureCode                         CNE    `hl7:"POS=28"`
	ProcedureCodeModifier                 []CNE  `hl7:"POS=29"`
	SpecialHandlingCode                   CWE    `hl7:"POS=30"`
}
