package hl7v24

import "time"

// HL7 v2.4 - CX - Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CX
type CX struct {
	Id                                         string    `hl7:"0" json:"Id,omitempty"`
	CheckDigit                                 string    `hl7:"1" json:"CheckDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD        `hl7:"2" json:"CodeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	AssigningAuthority                         HD        `hl7:"3" json:"AssigningAuthority,omitempty"`
	IdentifierTypeCode                         string    `hl7:"4" json:"IdentifierTypeCode,omitempty"`
	AssigningFacility                          HD        `hl7:"5" json:"AssigningFacility,omitempty"`
	EffectiveDate                              time.Time `hl7:"6" json:"EffectiveDate,omitempty"`
	ExpirationDate                             time.Time `hl7:"7" json:"ExpirationDate,omitempty"`
}

// HL7 v2.4 - CWE - Coded With Exceptions
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CWE
type CWE struct {
	Identifier                     string `hl7:"0" json:"Identifier,omitempty"`
	Text                           string `hl7:"1" json:"Text,omitempty"`
	NameOfCodingSystem             string `hl7:"2" json:"NameOfCodingSystem,omitempty"`
	AlternateIdentifier            string `hl7:"3" json:"AlternateIdentifier,omitempty"`
	AlternateText                  string `hl7:"4" json:"AlternateText,omitempty"`
	NameOfAlternateCodingSystem    string `hl7:"5" json:"NameOfAlternateCodingSystem,omitempty"`
	CodingSystemVersionId          string `hl7:"6" json:"CodingSystemVersionId,omitempty"`
	AlternateCodingSystemVersionId string `hl7:"7" json:"AlternateCodingSystemVersionId,omitempty"`
	OriginalText                   string `hl7:"8" json:"OriginalText,omitempty"`
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CE
type CE struct {
	Identifier                  string `hl7:"0" json:"Identifier,omitempty"`
	Text                        string `hl7:"1" json:"Text,omitempty"`
	NameOfCodingSystem          string `hl7:"2" json:"NameOfCodingSystem,omitempty"`
	AlternateIdentifier         string `hl7:"3" json:"AlternateIdentifier,omitempty"`
	AlternateText               string `hl7:"4" json:"AlternateText,omitempty"`
	NameOfAlternateCodingSystem string `hl7:"5" json:"NameOfAlternateCodingSystem,omitempty"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLN
type DLN struct {
	DriverLicenseNumber         string    `hl7:"0" json:"DriverLicenseNumber,omitempty"`
	IssuingStateProvinceCountry string    `hl7:"1" json:"IssuingStateProvinceCountry,omitempty"`
	ExpirationDate              time.Time `hl7:"2,shortdate" json:"ExpirationDate,omitempty"`
}

// DR - Date/time Range
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DR
type DR struct {
	RangeStartDateTime    time.Time `hl7:"0,longdate" json:"RangeStartDateTime,omitempty"`
	RangeStartEndDateTime time.Time `hl7:"1,longdate" json:"RangeStartEndDateTime,omitempty"`
}

// HL7 v2.4 - ELD - Error
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/ELD
type ELD struct {
	SegmentID            string `hl7:"0" json:"SegmentID,omitempty"`
	Sequence             int    `hl7:"1" json:"Sequence,omitempty"`
	FieldPosition        int    `hl7:"2" json:"FieldPosition,omitempty"`
	CodeIdentifyingError CE     `hl7:"3" json:"CodeIdentifyingError,omitempty"`
}

// FN - Family Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/FN
type FN struct {
	Surname              string `hl7:"0" json:"Surname,omitempty"`
	OwnSurnamePrefix     string `hl7:"1" json:"OwnSurnamePrefix,omitempty"`
	OwnSurname           string `hl7:"2" json:"OwnSurname,omitempty"`
	SurnamePrefixPartner string `hl7:"3" json:"SurnamePrefixPartner,omitempty"`
	SurnamePartner       string `hl7:"4" json:"SurnamePartner,omitempty"`
}

// HL7 v2.4 - HD - Hierarchic Designator
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/HD
type HD struct {
	NamespaceId     string `hl7:"0" json:"NamespaceId,omitempty"`
	UniversalId     string `hl7:"1" json:"UniversalId,omitempty"`
	UniversalIdType string `hl7:"2" json:"UniversalIdType,omitempty"`
}

// HL7 v2.4 - NA - Numeric Array
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NA
type NA struct {
	Value1 float32 `hl7:"0" json:"Value1,omitempty"`
	Value2 float32 `hl7:"1" json:"Value2,omitempty"`
	Value3 float32 `hl7:"2" json:"Value3,omitempty"`
	Value4 float32 `hl7:"3" json:"Value4,omitempty"`
}

// HL7 v2.4 - PI - Person Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PI
type PI struct {
	IDNumber            string `hl7:"0" json:"IDNumber,omitempty"`
	TypeOfIDNumber      string `hl7:"1" json:"TypeOfIDNumber,omitempty"`
	OtherQualifyingInfo string `hl7:"2" json:"OtherQualifyingInfo,omitempty"`
}

// HL7 v2.4 - SN - Structured Numeric
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SN
type SN struct {
	Comparator        string  `hl7:"0" json:"Comparator,omitempty"`
	Num1              float32 `hl7:"0" json:"Num1,omitempty"`
	SeparatorOrSuffix string  `hl7:"0" json:"SeparatorOrSuffix,omitempty"`
	Num2              float32 `hl7:"0" json:"Num2,omitempty"`
}

// HL7 v2.4 - VID - Version Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/VID
type VID struct {
	VersionID                string `hl7:"0" json:"VersionID,omitempty"`
	InternationalizationCode CE     `hl7:"1" json:"InternationalizationCode,omitempty"`
	InternationalVersionId   CE     `hl7:"2" json:"InternationalVersionId,omitempty"`
}

// HL7 v2.4 - XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XPN
type XPN struct {
	FamilyName         FN     `hl7:"0" json:"FamilyName,omitempty"`
	GivenName          string `hl7:"1" json:"GivenName,omitempty"`
	SecondGivenName    string `hl7:"2" json:"SecondGivenName,omitempty"`
	Suffix             string `hl7:"3" json:"Suffix,omitempty"`
	Prefix             string `hl7:"4" json:"Prefix,omitempty"`
	Degree             string `hl7:"5" json:"Degree,omitempty"`
	NameTypeCode       string `hl7:"6" json:"NameTypeCode,omitempty"`
	NameRepresentation string `hl7:"7" json:"NameRepresentation,omitempty"`
	NameContext        CE     `hl7:"8" json:"NameContext,omitempty"`
	NameValidityRange  DR     `hl7:"9" json:"NameValidityRange,omitempty"`
	NameAssemblyOrder  string `hl7:"10" json:"NameAssemblyOrder,omitempty"`
}

// HL7 v2.4 - XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XAD
type XAD struct {
	StreetAddress              string `hl7:"0" json:"StreetAddress,omitempty"`
	OtherDesignation           string `hl7:"1" json:"OtherDesignation,omitempty"`
	City                       string `hl7:"2" json:"City,omitempty"`
	StateOrProvince            string `hl7:"3" json:"StateOrProvince,omitempty"`
	ZipOrPostalCode            string `hl7:"4" json:"ZipOrPostalCode,omitempty"`
	Country                    string `hl7:"5" json:"Country,omitempty"`
	AddressType                string `hl7:"6" json:"AddressType,omitempty"`
	OtherGeographicDesignation string `hl7:"7" json:"OtherGeographicDesignation,omitempty"`
	CountyCode                 string `hl7:"8" json:"CountyCode,omitempty"`
	CensusTract                string `hl7:"9" json:"CensusTract,omitempty"`
	AddressRepresentationCode  string `hl7:"10" json:"AddressRepresentationCode,omitempty"`
	AddressValidityRange       DR     `hl7:"11" json:"AddressValidityRange,omitempty"`
}

// HL7 v2.4 - XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XTN
type XTN struct {
	TelephoneNumber                string `hl7:"0" json:"TelephoneNumber,omitempty"`
	TelecommunicationUseCode       string `hl7:"1" json:"TelecommunicationUseCode,omitempty"`
	TelecommunicationEquipmentType string `hl7:"2" json:"TelecommunicationEquipmentType,omitempty"`
	EmailAddress                   string `hl7:"3" json:"EmailAddress,omitempty"`
	CountryCode                    int    `hl7:"4" json:"CountryCode,omitempty"`
	AreaCode                       int    `hl7:"5" json:"AreaCode,omitempty"`
	PhoneNumber                    int    `hl7:"6" json:"PhoneNumber,omitempty"`
	Extension                      int    `hl7:"7" json:"Extension,omitempty"`
	AnyText                        string `hl7:"8" json:"AnyText,omitempty"`
}

// HL7 v2.4 - XON - Extended Composite Name And Identification Number For Organizations
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XON
type XON struct {
	OrganizationName                           string  `hl7:"0" json:"OrganizationName,omitempty"`
	OrganizationNameTypeCode                   string  `hl7:"1" json:"OrganizationNameTypeCode,omitempty"`
	IdNumber                                   float32 `hl7:"2" json:"IdNumber,omitempty"`
	CheckDigit                                 string  `hl7:"3" json:"CheckDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string  `hl7:"4" json:"CodeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	AssigningAuthority                         HD      `hl7:"5" json:"AssigningAuthority,omitempty"`
	IdentifyerTypeCode                         string  `hl7:"6" json:"IdentifyerTypeCode,omitempty"`
	AssigningFAcilityId                        HD      `hl7:"7" json:"AssigningFAcilityId,omitempty"`
	NameRepresentationCode                     string  `hl7:"8" json:"NameRepresentationCode,omitempty"`
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XCN
type XCN struct {
	ID                                         string `hl7:"0" json:"ID,omitempty"`
	FamilyName                                 FN     `hl7:"1" json:"FamilyName,omitempty"`
	GivenName                                  string `hl7:"2" json:"GivenName,omitempty"`
	MiddleInitialOrName                        string `hl7:"3" json:"MiddleInitialOrName,omitempty"`
	Suffix                                     string `hl7:"4" json:"Suffix,omitempty"`
	Prefix                                     string `hl7:"5" json:"Prefix,omitempty"`
	Degree                                     string `hl7:"6" json:"Degree,omitempty"`
	SourceTable                                string `hl7:"7" json:"SourceTable,omitempty"`
	AssigningAuthority                         HD     `hl7:"8" json:"AssigningAuthority,omitempty"`
	NameType                                   string `hl7:"9" json:"NameType,omitempty"`
	IdentifierCheckDigit                       string `hl7:"10" json:"IdentifierCheckDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"11" json:"CodeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	IdentifierTypeCode                         string `hl7:"12" json:"IdentifierTypeCode,omitempty"`
	AssigningFacility                          HD     `hl7:"13" json:"AssigningFacility,omitempty"`
	NameRepresentationCode                     string `hl7:"14" json:"NameRepresentationCode,omitempty"`
	NameContext                                CE     `hl7:"15" json:"NameContext,omitempty"`
	NameValidityRange                          DR     `hl7:"16" json:"NameValidityRange,omitempty"`
	NameAssemblyOrder                          string `hl7:"17" json:"NameAssemblyOrder,omitempty"`
}

// HL7 v2.4 - AUI - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/AUI
type AUI struct {
	AuthorizationNumber string    `hl7:"0" json:"AuthorizationNumber,omitempty"`
	Date                time.Time `hl7:"1,shortdate" json:"Date,omitempty"`
	Source              string    `hl7:"2" json:"Source,omitempty"`
}

// HL7 v2.4 - RMC - Room Coverage
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/RMC
type RMC struct {
	RoomType       string  `hl7:"0" json:"RoomType,omitempty"`
	AmountType     string  `hl7:"1" json:"AmountType,omitempty"`
	CoverageAmount float32 `hl7:"2" json:"CoverageAmount,omitempty"`
}

// HL7 v2.4 - PTA - Policy Type
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PTA
type PTA struct {
	PolicyType  string  `hl7:"0" json:"PolicyType,omitempty"`
	AmountClass string  `hl7:"1" json:"AmountClass,omitempty"`
	Amount      float32 `hl7:"2" json:"Amount,omitempty"`
}

// HL7 v2.4 - DDI - Daily Deductible
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DDI
type DDI struct {
	DelayDays    float32 `hl7:"0" json:"DelayDays,omitempty"`
	Amount       float32 `hl7:"1" json:"Amount,omitempty"`
	NumberOfDays float32 `hl7:"2" json:"NumberOfDays,omitempty"`
}

// HL7 v2.4 - SPS - Specimen Source
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SPS
type SPS struct {
	SourceNameOrCode             CE     `hl7:"0" json:"SourceNameOrCode,omitempty"`
	Additives                    string `hl7:"1" json:"Additives,omitempty"`
	Freetext                     string `hl7:"2" json:"Freetext,omitempty"`
	BodySite                     CE     `hl7:"3" json:"BodySite,omitempty"`
	SiteModifier                 CE     `hl7:"4" json:"SiteModifier,omitempty"`
	CollectionModifierMethodCode CE     `hl7:"5" json:"CollectionModifierMethodCode,omitempty"`
	Role                         CE     `hl7:"6" json:"Role,omitempty"`
}

// HL7 v2.4 - MOC - Charge To Practise
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MOC
type MOC struct {
	DollarAmount MO `hl7:"0" json:"DollarAmount,omitempty"`
	ChargeCode   CE `hl7:"1" json:"ChargeCode,omitempty"`
}

// HL7 v2.4 - PRL - Parent Result Link
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PRL
type PRL struct {
	ObservationIdentifierOfParentResult CE     `hl7:"0" json:"ObservationIdentifierOfParentResult,omitempty"`
	SubIDOfParentResult                 string `hl7:"1" json:"SubIDOfParentResult,omitempty"`
	ObservationResultFromParent         string `hl7:"2" json:"ObservationResultFromParent,omitempty"`
}

// HL7 v2.4 - NDL - Observing Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NDL
type NDL struct {
	OPName             CN        `hl7:"0" json:"OPName,omitempty"`
	StartDatetime      time.Time `hl7:"1" json:"StartDatetime,omitempty"`
	EndDatetime        time.Time `hl7:"2" json:"EndDatetime,omitempty"`
	PointOfCare        string    `hl7:"3" json:"PointOfCare,omitempty"`
	Room               string    `hl7:"4" json:"Room,omitempty"`
	Bed                string    `hl7:"5" json:"Bed,omitempty"`
	Facility           HD        `hl7:"6" json:"Facility,omitempty"`
	LocationStatus     string    `hl7:"7" json:"LocationStatus,omitempty"`
	PersonLocationType string    `hl7:"8" json:"PersonLocationType,omitempty"`
	Building           string    `hl7:"9" json:"Building,omitempty"`
	Floor              string    `hl7:"10" json:"Floor,omitempty"`
}

// HL7 v2.4 - MO - Money
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MO
type MO struct {
	Quantity     int    `hl7:"0" json:"Quantity,omitempty"`
	Denomination string `hl7:"1" json:"Denomination,omitempty"`
}

// HL7 v2.4 - CN - Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CN
type CN struct {
	IDNumber            string `hl7:"0" json:"IDNumber,omitempty"`
	FamilyName          string `hl7:"1" json:"FamilyName,omitempty"`
	GivenName           string `hl7:"2" json:"GivenName,omitempty"`
	MiddleInitialOrName string `hl7:"3" json:"MiddleInitialOrName,omitempty"`
	Suffix              string `hl7:"4" json:"Suffix,omitempty"`
	Prefix              string `hl7:"5" json:"Prefix,omitempty"`
	Degree              string `hl7:"6" json:"Degree,omitempty"`
	SourceTable         string `hl7:"7" json:"SourceTable,omitempty"`
	AssigningAuthority  string `hl7:"8" json:"AssigningAuthority,omitempty"`
}

// HL7 v2.4 - JCC - Job Code/class
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/JCC
type JCC struct {
	JobCode  string `hl7:"0" json:"JobCode,omitempty"`
	JobClass string `hl7:"1" json:"JobClass,omitempty"`
}

// HL7 v2.4 - CP - Composite Price
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CP
type CP struct {
	Price      float32 `hl7:"0" json:"Price,omitempty"`
	PriceType  string  `hl7:"1" json:"PriceType,omitempty"`
	FromValue  float32 `hl7:"2" json:"FromValue,omitempty"`
	ToValue    float32 `hl7:"3" json:"ToValue,omitempty"`
	RangeUnits CE      `hl7:"4" json:"RangeUnits,omitempty"`
	RangeType  string  `hl7:"5" json:"RangeType,omitempty"`
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
	PointOfCare         string `hl7:"0" json:"PointOfCare,omitempty"`
	Room                string `hl7:"1" json:"Room,omitempty"`
	Bed                 string `hl7:"2" json:"Bed,omitempty"`
	Facility            HD     `hl7:"3" json:"Facility,omitempty"`
	LocationStatus      string `hl7:"4" json:"LocationStatus,omitempty"`
	PersonLocationType  string `hl7:"5" json:"PersonLocationType,omitempty"`
	Building            string `hl7:"6" json:"Building,omitempty"`
	Floor               string `hl7:"7" json:"Floor,omitempty"`
	LocationDescription string `hl7:"8" json:"LocationDescription,omitempty"`
}

// HL7 v2.4 - DLD - Discharge Location
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLD
type DLD struct {
	DischargeLocation string    `hl7:"0" json:"DischargeLocation,omitempty"`
	EffectiveDate     time.Time `hl7:"1" json:"EffectiveDate,omitempty"`
}

// HL7 v2.3 - FC - Financial Class
//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string    `hl7:"0" json:"FinancialClass,omitempty"`
	EffectiveDate  time.Time `hl7:"1" json:"EffectiveDate,omitempty"`
}

// HL7 v2.4 - EIP - Parent Order
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EIP
type EIP struct {
	ParentsPlacerOrderNumber string `hl7:"0" json:"ParentsPlacerOrderNumber,omitempty"`
	ParentsFillerOrderNumber string `hl7:"1" json:"ParentsFillerOrderNumber,omitempty"`
}

// HL7 v2.4 - TQ - Timing Quantity
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/TQ
type TQ struct {
	Quantity          CQ        `hl7:"0" json:"Quantity,omitempty"`
	Interval          RI        `hl7:"1" json:"Interval,omitempty"`
	Duration          string    `hl7:"2" json:"Duration,omitempty"`
	StartDatetime     time.Time `hl7:"3" json:"StartDatetime,omitempty"`
	EndDatetime       time.Time `hl7:"4" json:"EndDatetime,omitempty"`
	Priority          string    `hl7:"5" json:"Priority,omitempty"`
	Condition         string    `hl7:"6" json:"Condition,omitempty"`
	Text              string    `hl7:"7" json:"Text,omitempty"`
	Conjunction       string    `hl7:"8" json:"Conjunction,omitempty"`
	OrderSequencing   OSD       `hl7:"9" json:"OrderSequencing,omitempty"`
	OccurenceDuration CE        `hl7:"10" json:"OccurenceDuration,omitempty"`
	TotalOccurences   float32   `hl7:"11" json:"TotalOccurences,omitempty"`
}

// CQ - Composite Quantity With Units
//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CQ
type CQ struct {
	Quantity int    `hl7:"0" json:"Quantity,omitempty"`
	Units    string `hl7:"1" json:"Units,omitempty"`
}

// HL7 v2.4 - RI - Repeat Interval
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/RI
type RI struct {
	RepeatPattern        string `hl7:"0" json:"RepeatPattern,omitempty"`
	ExplicitTimeInterval string `hl7:"1" json:"ExplicitTimeInterval,omitempty"`
}

// HL7 v2.4 - OSD - Order Sequence
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/OSD
type OSD struct {
	SequenceResultsFlag               string `hl7:"0" json:"SequenceResultsFlag,omitempty"`
	PlacerOrderNumberEntityIdentifier string `hl7:"1" json:"PlacerOrderNumberEntityIdentifier,omitempty"`
	PlacerOrderNumberNamespaceID      string `hl7:"2" json:"PlacerOrderNumberNamespaceID,omitempty"`
	FillerOrderNumberEntityIdentifier string `hl7:"3" json:"FillerOrderNumberEntityIdentifier,omitempty"`
	FillerOrderNumberNamespaceID      string `hl7:"4" json:"FillerOrderNumberNamespaceID,omitempty"`
	SequenceConditionValue            string `hl7:"5" json:"SequenceConditionValue,omitempty"`
	MaximumNumberOfRepeats            int    `hl7:"6" json:"MaximumNumberOfRepeats,omitempty"`
	PlacerOrderNumberUniversalID      string `hl7:"7" json:"PlacerOrderNumberUniversalID,omitempty"`
	PlacerOrderNumberUniversalIDType  string `hl7:"8" json:"PlacerOrderNumberUniversalIDType,omitempty"`
	FillerOrderNumberUniversalID      string `hl7:"9" json:"FillerOrderNumberUniversalID,omitempty"`
	FillerOrderNumberUniversalIDType  string `hl7:"10" json:"FillerOrderNumberUniversalIDType,omitempty"`
}

// HL7 v2.4 - EI - Entity Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"0" json:"EntityIdentifier,omitempty"`
	NamespaceId      string `hl7:"1" json:"NamespaceId,omitempty"`
	UniversalId      string `hl7:"2" json:"UniversalId,omitempty"`
	UniversalIdType  string `hl7:"3" json:"UniversalIdType,omitempty"`
}

/*
// CM_PEN - Penalty
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PEN
type CM_PEN struct {
	PenaltyType   string  `hl7:"0" json:",omitempty"`
	PenaltyAmount float32 `hl7:"1" json:",omitempty"`
}*/

// HL7 v2.4 - DTN - Day Type And Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DTN
type DTN struct {
	DayType      string  `hl7:"0" json:"DayType,omitempty"`
	NumberOfDays float32 `hl7:"1" json:"NumberOfDays,omitempty"`
}

// HL7 v2.4 - PCF - Pre-certification Required
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PCF
type PCF struct {
	PreCertificationPatient  string    `hl7:"0" json:"PreCertificationPatient,omitempty"`
	PreCertificationRequired string    `hl7:"1" json:"PreCertificationRequired,omitempty"`
	PreCertificationWindow   time.Time `hl7:"2,longdate" json:"PreCertificationWindow,omitempty"`
}

// HL7 v2.4 - LA1 - Location With Address Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/LA1
type LA1 struct {
	PointOfCare        string `hl7:"0" json:"PointOfCare,omitempty"`
	Room               string `hl7:"1" json:"Room,omitempty"`
	Bed                string `hl7:"2" json:"Bed,omitempty"`
	Facility           HD     `hl7:"3" json:"Facility,omitempty"`
	LocationStatus     string `hl7:"4" json:"LocationStatus,omitempty"`
	PersonLocationType string `hl7:"5" json:"PersonLocationType,omitempty"`
	Building           string `hl7:"6" json:"Building,omitempty"`
	Floor              string `hl7:"7" json:"Floor,omitempty"`
	Address            AD     `hl7:"8" json:"Address,omitempty"`
}

// HL7 v2.4 - AD - Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/AD
type AD struct {
	StreetAddress              string `hl7:"0" json:"StreetAddress,omitempty"`
	OtherDesignation           string `hl7:"1" json:"OtherDesignation,omitempty"`
	City                       string `hl7:"2" json:"City,omitempty"`
	StateOrProvince            string `hl7:"3" json:"StateOrProvince,omitempty"`
	PostalCode                 string `hl7:"4" json:"PostalCode,omitempty"`
	Country                    string `hl7:"5" json:"Country,omitempty"`
	AddressType                string `hl7:"6" json:"AddressType,omitempty"`
	OtherGeographicDesignation string `hl7:"7" json:"OtherGeographicDesignation,omitempty"`
}

// HL7 v2.4 - CCD - Charge Time
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CCD
type CCD struct {
	WhenToChargeCode string    `hl7:"1" json:"WhenToChargeCode,omitempty"`
	DateTime         time.Time `hl7:"2,longdate" json:"DateTime,omitempty"`
}

// HL7 v2.4 - CK - Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CK
type CK struct {
	IDNumber           int    `hl7:"0" json:"IDNumber,omitempty"`
	CheckDigit         string `hl7:"1" json:"CheckDigit,omitempty"`
	CodeIdentifyingCC  string `hl7:"2" json:"CodeIdentifyingCC,omitempty"`
	AssigningAuthority HD     `hl7:"3" json:"AssigningAuthority,omitempty"`
}

// HL7 v2.4 - MOP - Money Or Percentage
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MOP
type MOP struct {
	MoneyOrPercentageIndicator string  `hl7:"0" json:"MoneyOrPercentageIndicator,omitempty"`
	MoneyOrPercentageQuantity  float32 `hl7:"1" json:"MoneyOrPercentageQuantity,omitempty"`
}
