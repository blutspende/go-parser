package hl7v23

import "time"

// Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CX
type CX struct {
	Id                                         string `hl7:"1" json:"id,omitempty"`
	CheckDigit                                 string `hl7:"2" json:"checkDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD     `hl7:"3" json:"codeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	AssigningAuthority                         string `hl7:"4" json:"assigningAuthority,omitempty"`
	AssigningFacility                          HD     `hl7:"5" json:"assigningFacility,omitempty"`
}

// XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XPN
type XPN struct {
	FamilyName          string `hl7:"1" json:"familyName,omitempty"`
	GivenName           string `hl7:"2" json:"givenName,omitempty"`
	MiddleInitialOrName string `hl7:"3" json:"middleInitialOrName,omitempty"`
	Suffix              string `hl7:"4" json:"suffix,omitempty"`
	Prefix              string `hl7:"5" json:"prefix,omitempty"`
	Degree              string `hl7:"6" json:"degree,omitempty"`
	NameTypeCode        string `hl7:"7" json:"nameTypeCode,omitempty"`
	NameRepresentation  string `hl7:"8" json:"nameRepresentation,omitempty"`
}

// XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XAD
type XAD struct {
	StreetAddress              string `hl7:"1" json:"streetAddress,omitempty"`
	OtherDesignation           string `hl7:"2" json:"otherDesignation,omitempty"`
	City                       string `hl7:"3" json:"city,omitempty"`
	StateOrProvince            string `hl7:"4" json:"stateOrProvince,omitempty"`
	ZipOrPostalCode            string `hl7:"5" json:"zipOrPostalCode,omitempty"`
	Country                    string `hl7:"6" json:"country,omitempty"`
	AddressType                string `hl7:"7" json:"addressType,omitempty"`
	OtherGeographicDesignation string `hl7:"8" json:"otherGeographicDesignation,omitempty"`
	CountyCode                 string `hl7:"9" json:"countyCode,omitempty"`
	CensusTract                string `hl7:"10" json:"censusTract,omitempty"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XTN
type XTN struct {
	TelephoneNumber                string `hl7:"1" json:"telephoneNumber,omitempty"`
	TelecommunicationUseCode       string `hl7:"2" json:"telecommunicationUseCode,omitempty"`
	TelecommunicationEquipmentType string `hl7:"3" json:"telecommunicationEquipmentType,omitempty"`
	EmailAddress                   string `hl7:"4" json:"emailAddress,omitempty"`
	CountryCode                    int    `hl7:"5" json:"countryCode,omitempty"`
	AreaCode                       int    `hl7:"6" json:"areaCode,omitempty"`
	PhoneNumber                    int    `hl7:"7" json:"phoneNumber,omitempty"`
	Extension                      int    `hl7:"8" json:"extension,omitempty"`
	AnyText                        string `hl7:"9" json:"anyText,omitempty"`
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CE
type CE struct {
	Identifier                  string `hl7:"1" json:"identifier,omitempty"`
	Text                        string `hl7:"2" json:"text,omitempty"`
	NameOFCodingSystem          string `hl7:"3" json:"nameOFCodingSystem,omitempty"`
	AlternateIdentifier         string `hl7:"4" json:"alternateIdentifier,omitempty"`
	AlternateText               string `hl7:"5" json:"alternateText,omitempty"`
	NameOfAlternateCodingSystem string `hl7:"6" json:"nameOfAlternateCodingSystem,omitempty"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/DLN
type DLN struct {
	DriverLicenseNumber         string `hl7:"1" json:"driverLicenseNumber,omitempty"`
	IssuingStateProvinceCountry string `hl7:"2" json:"issuingStateProvinceCountry,omitempty"`
	ExpirationDate              string `hl7:"3" json:"expirationDate,omitempty"`
}

// XON - Extended Composite Name And ID For Organizations
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XON
type XON struct {
	OrganizationName                           string  `hl7:"1" json:"organizationName,omitempty"`
	OrganizationNameTypeCode                   string  `hl7:"2" json:"organizationNameTypeCode,omitempty"`
	IdNumber                                   float32 `hl7:"3" json:"idNumber,omitempty"`
	CheckDigit                                 string  `hl7:"4" json:"checkDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string  `hl7:"5" json:"codeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	AssigningAuthority                         HD      `hl7:"6" json:"assigningAuthority,omitempty"`
	IdentifierTypeCode                         string  `hl7:"7" json:"identifierTypeCode,omitempty"`
	AssigningFAcilityId                        HD      `hl7:"8" json:"assigningFAcilityId,omitempty"`
	NameRepresentationCode                     string  `hl7:"9" json:"nameRepresentationCode,omitempty"`
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XCN
type XCN struct {
	ID                                         string `hl7:"1" json:"id,omitempty"`
	FamilyName                                 string `hl7:"2" json:"familyName,omitempty"`
	GivenName                                  string `hl7:"3" json:"givenName,omitempty"`
	MiddleInitialOrName                        string `hl7:"4" json:"middleInitialOrName,omitempty"`
	Suffix                                     string `hl7:"5" json:"suffix,omitempty"`
	Prefix                                     string `hl7:"6" json:"prefix,omitempty"`
	Degree                                     string `hl7:"7" json:"degree,omitempty"`
	SourceTable                                string `hl7:"8" json:"sourceTable,omitempty"`
	AssigningAuthority                         HD     `hl7:"9" json:"assigningAuthority,omitempty"`
	NameType                                   string `hl7:"10" json:"nameType,omitempty"`
	IdentifierCheckDigit                       string `hl7:"11" json:"identifierCheckDigit,omitempty"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"12" json:"codeIdentifyingTheCheckDigitSchemeEmployed,omitempty"`
	IdentifierTypeCode                         string `hl7:"13" json:"identifierTypeCode,omitempty"`
	AssigningFacilityID                        HD     `hl7:"14" json:"assigningFacilityID,omitempty"`
}

// HD - Hierarchic Designator
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/HD
type HD struct {
	NamespaceId     string `hl7:"1" json:"namespaceId,omitempty"`
	UniversalId     string `hl7:"2" json:"universalId,omitempty"`
	UniversalIdType string `hl7:"3" json:"universalIdType,omitempty"`
}

// CM_AUI - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_AUI
type CM_AUI struct {
	AuthorizationNumber string    `hl7:"1" json:"authorizationNumber,omitempty"`
	Date                time.Time `hl7:"2,shortdate" json:"date,omitempty"`
	Source              string    `hl7:"3" json:"source,omitempty"`
}

// CM_RMC - Room Coverage
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_RMC
type CM_RMC struct {
	RoomType       string  `hl7:"1" json:"roomType,omitempty"`
	AmountType     string  `hl7:"2" json:"amountType,omitempty"`
	CoverageAmount float32 `hl7:"3" json:"coverageAmount,omitempty"`
}

// CM_PTA - Policy Type
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PTA
type CM_PTA struct {
	PolicyType  string  `hl7:"1" json:"policyType,omitempty"`
	AmountClass string  `hl7:"2" json:"amountClass,omitempty"`
	Amount      float32 `hl7:"3" json:"amount,omitempty"`
}

// CM_DDI - Daily Deductible
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DDI
type CM_DDI struct {
	DelayDays    float32 `hl7:"1" json:"delayDays,omitempty"`
	Amount       float32 `hl7:"2" json:"amount,omitempty"`
	NumberOfDays float32 `hl7:"3" json:"numberOfDays,omitempty"`
}

// CM_SPS - Specimen Source
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_SPS
type CM_SPS struct {
	SpecimenSourceNameOrCode     CE     `hl7:"1" json:"specimenSourceNameOrCode,omitempty"`
	Additives                    string `hl7:"2" json:"additives,omitempty"`
	Freetext                     string `hl7:"3" json:"freetext,omitempty"`
	BodySite                     CE     `hl7:"4" json:"bodySite,omitempty"`
	SiteModifier                 CE     `hl7:"5" json:"siteModifier,omitempty"`
	CollectionModifierMethodCode CE     `hl7:"6" json:"collectionModifierMethodCode,omitempty"`
}

// CM_MOC - Charge To Practise
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_MOC
type CM_MOC struct {
	DollarAmount MO `hl7:"1" json:"dollarAmount,omitempty"`
	ChargeCode   CE `hl7:"2" json:"chargeCode,omitempty"`
}

// CM_PRL - Parent Result Link
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PRL
type CM_PRL struct {
	ObservationIdentifierOfParentResult CE     `hl7:"1" json:"observationIdentifierOfParentResult,omitempty"`
	SubIDOfParentResult                 string `hl7:"2" json:"subIDOfParentResult,omitempty"`
	ObservationResultFromParent         string `hl7:"3" json:"observationResultFromParent,omitempty"`
}

// CM_NDL - Observing Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_NDL
type CM_NDL struct {
	OPName             CN        `hl7:"1" json:"opName,omitempty"`
	StartDatetime      time.Time `hl7:"2" json:"startDatetime,omitempty"`
	EndDatetime        time.Time `hl7:"3" json:"endDatetime,omitempty"`
	PointOfCare        string    `hl7:"4" json:"pointOfCare,omitempty"`
	Room               string    `hl7:"5" json:"room,omitempty"`
	Bed                string    `hl7:"6" json:"bed,omitempty"`
	Facility           HD        `hl7:"7" json:"facility,omitempty"`
	LocationStatus     string    `hl7:"8" json:"locationStatus,omitempty"`
	PersonLocationType string    `hl7:"9" json:"personLocationType,omitempty"`
	Building           string    `hl7:"10" json:"building,omitempty"`
	Floor              string    `hl7:"11" json:"floor,omitempty"`
}

// MO - Money
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/MO
type MO struct {
	Quantity     int    `hl7:"1" json:"quantity,omitempty"`
	Denomination string `hl7:"2" json:"denomination,omitempty"`
}

// CN - Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CN
type CN struct {
	IDNumber            string `hl7:"1" json:"idNumber,omitempty"`
	FamilyName          string `hl7:"2" json:"familyName,omitempty"`
	GivenName           string `hl7:"3" json:"givenName,omitempty"`
	MiddleInitialOrName string `hl7:"4" json:"middleInitialOrName,omitempty"`
	Suffix              string `hl7:"5" json:"suffix,omitempty"`
	Prefix              string `hl7:"6" json:"prefix,omitempty"`
	Degree              string `hl7:"7" json:"degree,omitempty"`
	SourceTable         string `hl7:"8" json:"sourceTable,omitempty"`
	AssigningAuthority  string `hl7:"9" json:"assigningAuthority,omitempty"`
}

// JCC - Job Code Class
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/JCC
type JCC struct {
	JobCode  string `hl7:"1" json:"jobCode,omitempty"`
	JobClass string `hl7:"2" json:"jobClass,omitempty"`
}

// HL7 v2.3 - CP - Composite Price
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CP
type CP struct {
	Price      float32 `hl7:"1" json:"price,omitempty"`
	PriceType  string  `hl7:"2" json:"priceType,omitempty"`
	FromValue  float32 `hl7:"3" json:"fromValue,omitempty"`
	ToValue    float32 `hl7:"4" json:"toValue,omitempty"`
	RangeUnits CE      `hl7:"5" json:"rangeUnits,omitempty"`
	RangeType  string  `hl7:"6" json:"rangeType,omitempty"`
}

type Sex string

const (
	Female  Sex = "F"
	Male    Sex = "M"
	Other   Sex = "O"
	Unknown Sex = "U"
)

// PL - Person Location
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/PL
type PL struct {
	PointOfCare        string `hl7:"1" json:"pointOfCare,omitempty"`
	Room               string `hl7:"2" json:"room,omitempty"`
	Bed                string `hl7:"3" json:"bed,omitempty"`
	Facility           HD     `hl7:"4" json:"facility,omitempty"`
	LocationStatus     string `hl7:"5" json:"locationStatus,omitempty"`
	PersonLocationType string `hl7:"6" json:"personLocationType,omitempty"`
	Building           string `hl7:"7" json:"building,omitempty"`
	Floor              string `hl7:"8" json:"floor,omitempty"`
	LocationType       string `hl7:"9" json:"locationType,omitempty"`
}

// CM_DLD - Discharge Location
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DLD
type CM_DLD struct {
	DischargeLocation string    `hl7:"1" json:"dischargeLocation,omitempty"`
	EffectiveDate     time.Time `hl7:"2" json:"effectiveDate,omitempty"`
}

// FC - Financial Class
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string    `hl7:"1" json:"financialClass,omitempty"`
	EffectiveDate  time.Time `hl7:"2" json:"effectiveDate,omitempty"`
}

// CM_EIP - Parent Order
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_EIP
type CM_EIP struct {
	ParentsPlacerOrderNumber string `hl7:"1" json:"parentsPlacerOrderNumber,omitempty"`
	ParentsFillerOrderNumber string `hl7:"2" json:"parentsFillerOrderNumber,omitempty"`
}

// TQ - Timing Quantity
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/TQ
type TQ struct {
	Quantity        CQ        `hl7:"1" json:"quantity,omitempty"`
	Interval        RI        `hl7:"2" json:"interval,omitempty"`
	Duration        string    `hl7:"3" json:"duration,omitempty"`
	StartDatetime   time.Time `hl7:"4" json:"startDatetime,omitempty"`
	EndDatetime     time.Time `hl7:"5" json:"endDatetime,omitempty"`
	Priority        string    `hl7:"6" json:"priority,omitempty"`
	Condition       string    `hl7:"7" json:"condition,omitempty"`
	Text            string    `hl7:"8" json:"text,omitempty"`
	Conjunction     string    `hl7:"9" json:"conjunction,omitempty"`
	OrderSequencing CM_OSD    `hl7:"10" json:"orderSequencing,omitempty"`
}

// CQ - Composite Quantity With Units
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CQ
type CQ struct {
	Quantity int    `hl7:"1" json:"quantity,omitempty"`
	Units    string `hl7:"2" json:"units,omitempty"`
}

// RI - Repeat Interval
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/RI
type RI struct {
	RepeatPattern        string `hl7:"1" json:"repeatPattern,omitempty"`
	ExplicitTimeInterval string `hl7:"2" json:"explicitTimeInterval,omitempty"`
}

// CM_OSD - Order Sequence
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_OSD
type CM_OSD struct {
	SequenceResultsFlag               string `hl7:"1" json:"sequenceResultsFlag,omitempty"`
	PlacerOrderNumberEntityIdentifier string `hl7:"2" json:"placerOrderNumberEntityIdentifier,omitempty"`
	PlacerOrderNumberNamespaceID      string `hl7:"3" json:"placerOrderNumberNamespaceID,omitempty"`
	FillerOrderNumberEntityIdentifier string `hl7:"4" json:"fillerOrderNumberEntityIdentifier,omitempty"`
	FillerOrderNumberNamespaceID      string `hl7:"5" json:"fillerOrderNumberNamespaceID,omitempty"`
	SequenceConditionValue            string `hl7:"6" json:"sequenceConditionValue,omitempty"`
	MaximumNumberOfRepeats            int    `hl7:"7" json:"maximumNumberOfRepeats,omitempty"`
	PlacerOrderNumberUniversalID      string `hl7:"8" json:"placerOrderNumberUniversalID,omitempty"`
	PlacerOrderNumberUniversalIDType  string `hl7:"9" json:"placerOrderNumberUniversalIDType,omitempty"`
	FillerOrderNumberUniversalID      string `hl7:"10" json:"fillerOrderNumberUniversalID,omitempty"`
	FillerOrderNumberUniversalIDType  string `hl7:"11" json:"fillerOrderNumberUniversalIDType,omitempty"`
}

// EI - Entity Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"1" json:"entityIdentifier,omitempty"`
	NamespaceId      string `hl7:"2" json:"namespaceId,omitempty"`
	UniversalId      string `hl7:"3" json:"universalId,omitempty"`
	UniversalIdType  string `hl7:"4" json:"universalIdType,omitempty"`
}

// CM_PEN - Penalty
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PEN
type CM_PEN struct {
	PenaltyType   string  `hl7:"1" json:"penaltyType,omitempty"`
	PenaltyAmount float32 `hl7:"2" json:"penaltyAmount,omitempty"`
}

// CM_DTN - Day Type And Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DTN
type CM_DTN struct {
	DayType      string  `hl7:"1" json:"dayType,omitempty"`
	NumberOfDays float32 `hl7:"2" json:"numberOfDays,omitempty"`
}

// CM_PCF - Pre-certification Required
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PCF
type CM_PCF struct {
	PreCertificationPatient  string `hl7:"1" json:"preCertificationPatient,omitempty"`
	PreCertificationRequired string `hl7:"2" json:"preCertificationRequired,omitempty"`
	PreCertificationWindow   string `hl7:"3" json:"preCertificationWindow,omitempty"`
}

// LA1 - Location With Address Information
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/LA1
type LA1 struct {
	PointOfCare        string `hl7:"1" json:"pointOfCare,omitempty"`
	Room               string `hl7:"2" json:"room,omitempty"`
	Bed                string `hl7:"3" json:"bed,omitempty"`
	Facility           HD     `hl7:"4" json:"facility,omitempty"`
	LocationStatus     string `hl7:"5" json:"locationStatus,omitempty"`
	PersonLocationType string `hl7:"6" json:"personLocationType,omitempty"`
	Building           string `hl7:"7" json:"building,omitempty"`
	Floor              string `hl7:"8" json:"floor,omitempty"`
	Address            AD     `hl7:"9" json:"address,omitempty"`
}

// HL7 v2.3 - AD - Address
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/AD
type AD struct {
	StreetAddress              string `hl7:"1" json:"streetAddress,omitempty"`
	OtherDesignation           string `hl7:"2" json:"otherDesignation,omitempty"`
	City                       string `hl7:"3" json:"city,omitempty"`
	StateOrProvince            string `hl7:"4" json:"stateOrProvince,omitempty"`
	PostalCode                 string `hl7:"5" json:"postalCode,omitempty"`
	Country                    string `hl7:"6" json:"country,omitempty"`
	AddressType                string `hl7:"7" json:"addressType,omitempty"`
	OtherGeographicDesignation string `hl7:"8" json:"otherGeographicDesignation,omitempty"`
}

// CM_CCD - Charge Time
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_CCD
type CM_CCD struct {
	WhenToChargeCode string    `hl7:"1" json:"whenToChargeCode,omitempty"`
	DateTime         time.Time `hl7:"2,longdate" json:"dateTime,omitempty"`
}

// CK - Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CK
type CK struct {
	IDNumber           int    `hl7:"1" json:"idNumber,omitempty"`
	CheckDigit         string `hl7:"2" json:"checkDigit,omitempty"`
	CodeIdentifyingCC  string `hl7:"3" json:"codeIdentifyingCC,omitempty"`
	AssigningAuthority HD     `hl7:"4" json:"assigningAuthority,omitempty"`
}
