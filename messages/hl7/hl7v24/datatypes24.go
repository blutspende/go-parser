package hl7v24

import "time"

// HL7 v2.4 - CX - Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CX
type CX struct {
	Id                                         string    `hl7:"1" json:"Id,omitempty"`
	CheckDigit                                 string    `hl7:"2" json:"CheckDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD        `hl7:"3" json:"CodeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	AssigningAuthority                         HD        `hl7:"4" json:"AssigningAuthority,omitempty"`
	IdentifierTypeCode                         string    `hl7:"5" json:"IdentifierTypeCode,omitempty"`
	AssigningFacility                          HD        `hl7:"6" json:"AssigningFacility,omitempty"`
	EffectiveDate                              time.Time `hl7:"7" json:"EffectiveDate,omitempty"`
	ExpirationDate                             time.Time `hl7:"8" json:"ExpirationDate,omitempty"`
}

// HL7 v2.4 - CWE - Coded With Exceptions
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CWE
type CWE struct {
	Identifier                     string `hl7:"1" json:"Identifier,omitempty"`
	Text                           string `hl7:"2" json:"Text,omitempty"`
	NameOfCodingSystem             string `hl7:"3" json:"NameOfCodingSystem,omitempty"`
	AlternateIdentifier            string `hl7:"4" json:"AlternateIdentifier,omitempty"`
	AlternateText                  string `hl7:"5" json:"AlternateText,omitempty"`
	NameOfAlternateCodingSystem    string `hl7:"6" json:"NameOfAlternateCodingSystem,omitempty"`
	CodingSystemVersionId          string `hl7:"7" json:"CodingSystemVersionId,omitempty"`
	AlternateCodingSystemVersionId string `hl7:"8" json:"AlternateCodingSystemVersionId,omitempty"`
	OriginalText                   string `hl7:"9" json:"OriginalText,omitempty"`
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CE
type CE struct {
	Identifier                  string `hl7:"1" json:"Identifier,omitempty"`
	Text                        string `hl7:"2" json:"Text,omitempty"`
	NameOfCodingSystem          string `hl7:"3" json:"NameOfCodingSystem,omitempty"`
	AlternateIdentifier         string `hl7:"4" json:"AlternateIdentifier,omitempty"`
	AlternateText               string `hl7:"5" json:"AlternateText,omitempty"`
	NameOfAlternateCodingSystem string `hl7:"6" json:"NameOfAlternateCodingSystem,omitempty"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLN
type DLN struct {
	DriverLicenseNumber         string    `hl7:"1" json:"DriverLicenseNumber,omitempty"`
	IssuingStateProvinceCountry string    `hl7:"2" json:"IssuingStateProvinceCountry,omitempty"`
	ExpirationDate              time.Time `hl7:"3,shortdate" json:"ExpirationDate,omitempty"`
}

// DR - Date/time Range
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DR
type DR struct {
	RangeStartDateTime    time.Time `hl7:"1,longdate" json:"RangeStartDateTime,omitempty"`
	RangeStartEndDateTime time.Time `hl7:"2,longdate" json:"RangeStartEndDateTime,omitempty"`
}

// HL7 v2.4 - ELD - Error
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/ELD
type ELD struct {
	SegmentID            string `hl7:"1" json:"SegmentID,omitempty"`
	Sequence             int    `hl7:"2" json:"Sequence,omitempty"`
	FieldPosition        int    `hl7:"3" json:"FieldPosition,omitempty"`
	CodeIdentifyingError CE     `hl7:"4" json:"CodeIdentifyingError,omitempty"`
}

// FN - Family Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/FN
type FN struct {
	Surname              string `hl7:"1" json:"Surname,omitempty"`
	OwnSurnamePrefix     string `hl7:"2" json:"OwnSurnamePrefix,omitempty"`
	OwnSurname           string `hl7:"3" json:"OwnSurname,omitempty"`
	SurnamePrefixPartner string `hl7:"4" json:"SurnamePrefixPartner,omitempty"`
	SurnamePartner       string `hl7:"5" json:"SurnamePartner,omitempty"`
}

// HL7 v2.4 - HD - Hierarchic Designator
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/HD
type HD struct {
	NamespaceId     string `hl7:"1" json:"NamespaceId,omitempty"`
	UniversalId     string `hl7:"2" json:"UniversalId,omitempty"`
	UniversalIdType string `hl7:"3" json:"UniversalIdType,omitempty"`
}

// HL7 v2.4 - NA - Numeric Array
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NA
type NA struct {
	Value1 float32 `hl7:"1" json:"Value1,omitempty"`
	Value2 float32 `hl7:"2" json:"Value2,omitempty"`
	Value3 float32 `hl7:"3" json:"Value3,omitempty"`
	Value4 float32 `hl7:"4" json:"Value4,omitempty"`
}

// HL7 v2.4 - PI - Person Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PI
type PI struct {
	IDNumber            string `hl7:"1" json:"IDNumber,omitempty"`
	TypeOfIDNumber      string `hl7:"2" json:"TypeOfIDNumber,omitempty"`
	OtherQualifyingInfo string `hl7:"3" json:"OtherQualifyingInfo,omitempty"`
}

// HL7 v2.4 - SN - Structured Numeric
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SN
type SN struct {
	// TODO: figure out why it was all "1"s in the original
	Comparator        string  `hl7:"1" json:"Comparator,omitempty"`
	Num1              float32 `hl7:"2" json:"Num1,omitempty"`
	SeparatorOrSuffix string  `hl7:"3" json:"SeparatorOrSuffix,omitempty"`
	Num2              float32 `hl7:"4" json:"Num2,omitempty"`
}

// HL7 v2.4 - VID - Version Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/VID
type VID struct {
	VersionID                string `hl7:"1" json:"VersionID,omitempty"`
	InternationalizationCode CE     `hl7:"2" json:"InternationalizationCode,omitempty"`
	InternationalVersionId   CE     `hl7:"3" json:"InternationalVersionId,omitempty"`
}

// HL7 v2.4 - XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XPN
type XPN struct {
	FamilyName         FN     `hl7:"1" json:"FamilyName,omitempty"`
	GivenName          string `hl7:"2" json:"GivenName,omitempty"`
	SecondGivenName    string `hl7:"3" json:"SecondGivenName,omitempty"`
	Suffix             string `hl7:"4" json:"Suffix,omitempty"`
	Prefix             string `hl7:"5" json:"Prefix,omitempty"`
	Degree             string `hl7:"6" json:"Degree,omitempty"`
	NameTypeCode       string `hl7:"7" json:"NameTypeCode,omitempty"`
	NameRepresentation string `hl7:"8" json:"NameRepresentation,omitempty"`
	NameContext        CE     `hl7:"9" json:"NameContext,omitempty"`
	NameValidityRange  DR     `hl7:"10" json:"NameValidityRange,omitempty"`
	NameAssemblyOrder  string `hl7:"11" json:"NameAssemblyOrder,omitempty"`
}

// HL7 v2.4 - XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XAD
type XAD struct {
	StreetAddress              string `hl7:"1" json:"StreetAddress,omitempty"`
	OtherDesignation           string `hl7:"2" json:"OtherDesignation,omitempty"`
	City                       string `hl7:"3" json:"City,omitempty"`
	StateOrProvince            string `hl7:"4" json:"StateOrProvince,omitempty"`
	ZipOrPostalCode            string `hl7:"5" json:"ZipOrPostalCode,omitempty"`
	Country                    string `hl7:"6" json:"Country,omitempty"`
	AddressType                string `hl7:"7" json:"AddressType,omitempty"`
	OtherGeographicDesignation string `hl7:"8" json:"OtherGeographicDesignation,omitempty"`
	CountyCode                 string `hl7:"9" json:"CountyCode,omitempty"`
	CensusTract                string `hl7:"10" json:"CensusTract,omitempty"`
	AddressRepresentationCode  string `hl7:"11" json:"AddressRepresentationCode,omitempty"`
	AddressValidityRange       DR     `hl7:"12" json:"AddressValidityRange,omitempty"`
}

// HL7 v2.4 - XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XTN
type XTN struct {
	TelephoneNumber                string `hl7:"1" json:"TelephoneNumber,omitempty"`
	TelecommunicationUseCode       string `hl7:"2" json:"TelecommunicationUseCode,omitempty"`
	TelecommunicationEquipmentType string `hl7:"3" json:"TelecommunicationEquipmentType,omitempty"`
	EmailAddress                   string `hl7:"4" json:"EmailAddress,omitempty"`
	CountryCode                    int    `hl7:"5" json:"CountryCode,omitempty"`
	AreaCode                       int    `hl7:"6" json:"AreaCode,omitempty"`
	PhoneNumber                    int    `hl7:"7" json:"PhoneNumber,omitempty"`
	Extension                      int    `hl7:"8" json:"Extension,omitempty"`
	AnyText                        string `hl7:"9" json:"AnyText,omitempty"`
}

// HL7 v2.4 - XON - Extended Composite Name And Identification Number For Organizations
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XON
type XON struct {
	OrganizationName                           string  `hl7:"1" json:"OrganizationName,omitempty"`
	OrganizationNameTypeCode                   string  `hl7:"2" json:"OrganizationNameTypeCode,omitempty"`
	IdNumber                                   float32 `hl7:"3" json:"IdNumber,omitempty"`
	CheckDigit                                 string  `hl7:"4" json:"CheckDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string  `hl7:"5" json:"CodeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	AssigningAuthority                         HD      `hl7:"6" json:"AssigningAuthority,omitempty"`
	IdentifyerTypeCode                         string  `hl7:"7" json:"IdentifyerTypeCode,omitempty"`
	AssigningFAcilityId                        HD      `hl7:"8" json:"AssigningFAcilityId,omitempty"`
	NameRepresentationCode                     string  `hl7:"9" json:"NameRepresentationCode,omitempty"`
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XCN
type XCN struct {
	ID                                         string `hl7:"1" json:"ID,omitempty"`
	FamilyName                                 FN     `hl7:"2" json:"FamilyName,omitempty"`
	GivenName                                  string `hl7:"3" json:"GivenName,omitempty"`
	MiddleInitialOrName                        string `hl7:"4" json:"MiddleInitialOrName,omitempty"`
	Suffix                                     string `hl7:"5" json:"Suffix,omitempty"`
	Prefix                                     string `hl7:"6" json:"Prefix,omitempty"`
	Degree                                     string `hl7:"7" json:"Degree,omitempty"`
	SourceTable                                string `hl7:"8" json:"SourceTable,omitempty"`
	AssigningAuthority                         HD     `hl7:"9" json:"AssigningAuthority,omitempty"`
	NameType                                   string `hl7:"10" json:"NameType,omitempty"`
	IdentifierCheckDigit                       string `hl7:"11" json:"IdentifierCheckDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"12" json:"CodeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	IdentifierTypeCode                         string `hl7:"13" json:"IdentifierTypeCode,omitempty"`
	AssigningFacility                          HD     `hl7:"14" json:"AssigningFacility,omitempty"`
	NameRepresentationCode                     string `hl7:"15" json:"NameRepresentationCode,omitempty"`
	NameContext                                CE     `hl7:"16" json:"NameContext,omitempty"`
	NameValidityRange                          DR     `hl7:"17" json:"NameValidityRange,omitempty"`
	NameAssemblyOrder                          string `hl7:"18" json:"NameAssemblyOrder,omitempty"`
}

// HL7 v2.4 - AUI - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/AUI
type AUI struct {
	AuthorizationNumber string    `hl7:"1" json:"AuthorizationNumber,omitempty"`
	Date                time.Time `hl7:"2,shortdate" json:"Date,omitempty"`
	Source              string    `hl7:"3" json:"Source,omitempty"`
}

// HL7 v2.4 - RMC - Room Coverage
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/RMC
type RMC struct {
	RoomType       string  `hl7:"1" json:"RoomType,omitempty"`
	AmountType     string  `hl7:"2" json:"AmountType,omitempty"`
	CoverageAmount float32 `hl7:"3" json:"CoverageAmount,omitempty"`
}

// HL7 v2.4 - PTA - Policy Type
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PTA
type PTA struct {
	PolicyType  string  `hl7:"1" json:"PolicyType,omitempty"`
	AmountClass string  `hl7:"2" json:"AmountClass,omitempty"`
	Amount      float32 `hl7:"3" json:"Amount,omitempty"`
}

// HL7 v2.4 - DDI - Daily Deductible
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DDI
type DDI struct {
	DelayDays    float32 `hl7:"1" json:"DelayDays,omitempty"`
	Amount       float32 `hl7:"2" json:"Amount,omitempty"`
	NumberOfDays float32 `hl7:"3" json:"NumberOfDays,omitempty"`
}

// HL7 v2.4 - SPS - Specimen Source
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SPS
type SPS struct {
	SourceNameOrCode             CE     `hl7:"1" json:"SourceNameOrCode,omitempty"`
	Additives                    string `hl7:"2" json:"Additives,omitempty"`
	Freetext                     string `hl7:"3" json:"Freetext,omitempty"`
	BodySite                     CE     `hl7:"4" json:"BodySite,omitempty"`
	SiteModifier                 CE     `hl7:"5" json:"SiteModifier,omitempty"`
	CollectionModifierMethodCode CE     `hl7:"6" json:"CollectionModifierMethodCode,omitempty"`
	Role                         CE     `hl7:"7" json:"Role,omitempty"`
}

// HL7 v2.4 - MOC - Charge To Practise
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MOC
type MOC struct {
	DollarAmount MO `hl7:"1" json:"DollarAmount,omitempty"`
	ChargeCode   CE `hl7:"2" json:"ChargeCode,omitempty"`
}

// HL7 v2.4 - PRL - Parent Result Link
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PRL
type PRL struct {
	ObservationIdentifierOfParentResult CE     `hl7:"1" json:"ObservationIdentifierOfParentResult,omitempty"`
	SubIDOfParentResult                 string `hl7:"2" json:"SubIDOfParentResult,omitempty"`
	ObservationResultFromParent         string `hl7:"3" json:"ObservationResultFromParent,omitempty"`
}

// HL7 v2.4 - NDL - Observing Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NDL
type NDL struct {
	OPName             CN        `hl7:"1" json:"OPName,omitempty"`
	StartDatetime      time.Time `hl7:"2" json:"StartDatetime,omitempty"`
	EndDatetime        time.Time `hl7:"3" json:"EndDatetime,omitempty"`
	PointOfCare        string    `hl7:"4" json:"PointOfCare,omitempty"`
	Room               string    `hl7:"5" json:"Room,omitempty"`
	Bed                string    `hl7:"6" json:"Bed,omitempty"`
	Facility           HD        `hl7:"7" json:"Facility,omitempty"`
	LocationStatus     string    `hl7:"8" json:"LocationStatus,omitempty"`
	PersonLocationType string    `hl7:"9" json:"PersonLocationType,omitempty"`
	Building           string    `hl7:"10" json:"Building,omitempty"`
	Floor              string    `hl7:"11" json:"Floor,omitempty"`
}

// HL7 v2.4 - MO - Money
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MO
type MO struct {
	Quantity     int    `hl7:"1" json:"Quantity,omitempty"`
	Denomination string `hl7:"2" json:"Denomination,omitempty"`
}

// HL7 v2.4 - CN - Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CN
type CN struct {
	IDNumber            string `hl7:"1" json:"IDNumber,omitempty"`
	FamilyName          string `hl7:"2" json:"FamilyName,omitempty"`
	GivenName           string `hl7:"3" json:"GivenName,omitempty"`
	MiddleInitialOrName string `hl7:"4" json:"MiddleInitialOrName,omitempty"`
	Suffix              string `hl7:"5" json:"Suffix,omitempty"`
	Prefix              string `hl7:"6" json:"Prefix,omitempty"`
	Degree              string `hl7:"7" json:"Degree,omitempty"`
	SourceTable         string `hl7:"8" json:"SourceTable,omitempty"`
	AssigningAuthority  string `hl7:"9" json:"AssigningAuthority,omitempty"`
}

// HL7 v2.4 - JCC - Job Code/class
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/JCC
type JCC struct {
	JobCode  string `hl7:"1" json:"JobCode,omitempty"`
	JobClass string `hl7:"2" json:"JobClass,omitempty"`
}

// HL7 v2.4 - CP - Composite Price
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CP
type CP struct {
	Price      float32 `hl7:"1" json:"Price,omitempty"`
	PriceType  string  `hl7:"2" json:"PriceType,omitempty"`
	FromValue  float32 `hl7:"3" json:"FromValue,omitempty"`
	ToValue    float32 `hl7:"4" json:"ToValue,omitempty"`
	RangeUnits CE      `hl7:"5" json:"RangeUnits,omitempty"`
	RangeType  string  `hl7:"6" json:"RangeType,omitempty"`
}

type Sex string

const (
	Female  Sex = "F"
	Male    Sex = "M"
	Other   Sex = "O"
	Unknown Sex = "U"
)

// HL7 v2.4 - PL - Person Location
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PL
type PL struct {
	PointOfCare         string `hl7:"1" json:"PointOfCare,omitempty"`
	Room                string `hl7:"2" json:"Room,omitempty"`
	Bed                 string `hl7:"3" json:"Bed,omitempty"`
	Facility            HD     `hl7:"4" json:"Facility,omitempty"`
	LocationStatus      string `hl7:"5" json:"LocationStatus,omitempty"`
	PersonLocationType  string `hl7:"6" json:"PersonLocationType,omitempty"`
	Building            string `hl7:"7" json:"Building,omitempty"`
	Floor               string `hl7:"8" json:"Floor,omitempty"`
	LocationDescription string `hl7:"9" json:"LocationDescription,omitempty"`
}

// HL7 v2.4 - DLD - Discharge Location
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLD
type DLD struct {
	DischargeLocation string    `hl7:"1" json:"DischargeLocation,omitempty"`
	EffectiveDate     time.Time `hl7:"2" json:"EffectiveDate,omitempty"`
}

// HL7 v2.3 - FC - Financial Class
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string    `hl7:"1" json:"FinancialClass,omitempty"`
	EffectiveDate  time.Time `hl7:"2" json:"EffectiveDate,omitempty"`
}

// HL7 v2.4 - EIP - Parent Order
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EIP
type EIP struct {
	ParentsPlacerOrderNumber string `hl7:"1" json:"ParentsPlacerOrderNumber,omitempty"`
	ParentsFillerOrderNumber string `hl7:"2" json:"ParentsFillerOrderNumber,omitempty"`
}

// HL7 v2.4 - TQ - Timing Quantity
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/TQ
type TQ struct {
	Quantity          CQ        `hl7:"1" json:"Quantity,omitempty"`
	Interval          RI        `hl7:"2" json:"Interval,omitempty"`
	Duration          string    `hl7:"3" json:"Duration,omitempty"`
	StartDatetime     time.Time `hl7:"4" json:"StartDatetime,omitempty"`
	EndDatetime       time.Time `hl7:"5" json:"EndDatetime,omitempty"`
	Priority          string    `hl7:"6" json:"Priority,omitempty"`
	Condition         string    `hl7:"7" json:"Condition,omitempty"`
	Text              string    `hl7:"8" json:"Text,omitempty"`
	Conjunction       string    `hl7:"9" json:"Conjunction,omitempty"`
	OrderSequencing   OSD       `hl7:"10" json:"OrderSequencing,omitempty"`
	OccurenceDuration CE        `hl7:"11" json:"OccurenceDuration,omitempty"`
	TotalOccurences   float32   `hl7:"12" json:"TotalOccurences,omitempty"`
}

// CQ - Composite Quantity With Units
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CQ
type CQ struct {
	Quantity int    `hl7:"1" json:"Quantity,omitempty"`
	Units    string `hl7:"2" json:"Units,omitempty"`
}

// HL7 v2.4 - RI - Repeat Interval
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/RI
type RI struct {
	RepeatPattern        string `hl7:"1" json:"RepeatPattern,omitempty"`
	ExplicitTimeInterval string `hl7:"2" json:"ExplicitTimeInterval,omitempty"`
}

// HL7 v2.4 - OSD - Order Sequence
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/OSD
type OSD struct {
	SequenceResultsFlag               string `hl7:"1" json:"SequenceResultsFlag,omitempty"`
	PlacerOrderNumberEntityIdentifier string `hl7:"2" json:"PlacerOrderNumberEntityIdentifier,omitempty"`
	PlacerOrderNumberNamespaceID      string `hl7:"3" json:"PlacerOrderNumberNamespaceID,omitempty"`
	FillerOrderNumberEntityIdentifier string `hl7:"4" json:"FillerOrderNumberEntityIdentifier,omitempty"`
	FillerOrderNumberNamespaceID      string `hl7:"5" json:"FillerOrderNumberNamespaceID,omitempty"`
	SequenceConditionValue            string `hl7:"6" json:"SequenceConditionValue,omitempty"`
	MaximumNumberOfRepeats            int    `hl7:"7" json:"MaximumNumberOfRepeats,omitempty"`
	PlacerOrderNumberUniversalID      string `hl7:"8" json:"PlacerOrderNumberUniversalID,omitempty"`
	PlacerOrderNumberUniversalIDType  string `hl7:"9" json:"PlacerOrderNumberUniversalIDType,omitempty"`
	FillerOrderNumberUniversalID      string `hl7:"10" json:"FillerOrderNumberUniversalID,omitempty"`
	FillerOrderNumberUniversalIDType  string `hl7:"11" json:"FillerOrderNumberUniversalIDType,omitempty"`
}

// HL7 v2.4 - EI - Entity Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"1" json:"EntityIdentifier,omitempty"`
	NamespaceId      string `hl7:"2" json:"NamespaceId,omitempty"`
	UniversalId      string `hl7:"3" json:"UniversalId,omitempty"`
	UniversalIdType  string `hl7:"4" json:"UniversalIdType,omitempty"`
}

/*
// CM_PEN - Penalty
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PEN
type CM_PEN struct {
	PenaltyType   string  `hl7:"1" json:",omitempty"`
	PenaltyAmount float32 `hl7:"2" json:",omitempty"`
}*/

// HL7 v2.4 - DTN - Day Type And Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DTN
type DTN struct {
	DayType      string  `hl7:"1" json:"DayType,omitempty"`
	NumberOfDays float32 `hl7:"2" json:"NumberOfDays,omitempty"`
}

// HL7 v2.4 - PCF - Pre-certification Required
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PCF
type PCF struct {
	PreCertificationPatient  string    `hl7:"1" json:"PreCertificationPatient,omitempty"`
	PreCertificationRequired string    `hl7:"2" json:"PreCertificationRequired,omitempty"`
	PreCertificationWindow   time.Time `hl7:"3,longdate" json:"PreCertificationWindow,omitempty"`
}

// HL7 v2.4 - LA1 - Location With Address Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/LA1
type LA1 struct {
	PointOfCare        string `hl7:"1" json:"PointOfCare,omitempty"`
	Room               string `hl7:"2" json:"Room,omitempty"`
	Bed                string `hl7:"3" json:"Bed,omitempty"`
	Facility           HD     `hl7:"4" json:"Facility,omitempty"`
	LocationStatus     string `hl7:"5" json:"LocationStatus,omitempty"`
	PersonLocationType string `hl7:"6" json:"PersonLocationType,omitempty"`
	Building           string `hl7:"7" json:"Building,omitempty"`
	Floor              string `hl7:"8" json:"Floor,omitempty"`
	Address            AD     `hl7:"9" json:"Address,omitempty"`
}

// HL7 v2.4 - AD - Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/AD
type AD struct {
	StreetAddress              string `hl7:"1" json:"StreetAddress,omitempty"`
	OtherDesignation           string `hl7:"2" json:"OtherDesignation,omitempty"`
	City                       string `hl7:"3" json:"City,omitempty"`
	StateOrProvince            string `hl7:"4" json:"StateOrProvince,omitempty"`
	PostalCode                 string `hl7:"5" json:"PostalCode,omitempty"`
	Country                    string `hl7:"6" json:"Country,omitempty"`
	AddressType                string `hl7:"7" json:"AddressType,omitempty"`
	OtherGeographicDesignation string `hl7:"8" json:"OtherGeographicDesignation,omitempty"`
}

// HL7 v2.4 - CCD - Charge Time
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CCD
type CCD struct {
	WhenToChargeCode string    `hl7:"2" json:"WhenToChargeCode,omitempty"`
	DateTime         time.Time `hl7:"3,longdate" json:"DateTime,omitempty"`
}

// HL7 v2.4 - CK - Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CK
type CK struct {
	IDNumber           int    `hl7:"1" json:"IDNumber,omitempty"`
	CheckDigit         string `hl7:"2" json:"CheckDigit,omitempty"`
	CodeIdentifyingCC  string `hl7:"3" json:"CodeIdentifyingCC,omitempty"`
	AssigningAuthority HD     `hl7:"4" json:"AssigningAuthority,omitempty"`
}

// HL7 v2.4 - MOP - Money Or Percentage
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MOP
type MOP struct {
	MoneyOrPercentageIndicator string  `hl7:"1" json:"MoneyOrPercentageIndicator,omitempty"`
	MoneyOrPercentageQuantity  float32 `hl7:"2" json:"MoneyOrPercentageQuantity,omitempty"`
}
