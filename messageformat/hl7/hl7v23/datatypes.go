package hl7v23

import "time"

// CM_RFR - Reference Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_RFR
type CM_RFR struct {
	ReferenceRange CM_RANGE `hl7:"POS=1"`
	Sex string `hl7:"POS=2"`
	AgeRange CM_RANGE `hl7:"POS=3"`
	AgeGestation CM_RANGE `hl7:"POS=4"`
	Species string `hl7:"POS=5"`
	RaceSubspecies string `hl7:"POS=6"`
	Conditions string `hl7:"POS=7"`
}

// CM_PLN - Practitioner ID Numbers
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_PLN
type CM_PLN struct {
	IDNumber string `hl7:"POS=1"`
	TypeOfIDNumber string `hl7:"POS=2"`
	StateOtherQualifyingInfo string `hl7:"POS=3"`
	ExpirationDate time.Time `hl7:"POS=4;ATR=date"`
}

// CM_NDL - Observing Practitioner
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_NDL
type CM_NDL struct {
	OpName CN `hl7:"POS=1"`
	StartDateTime time.Time `hl7:"POS=2"`
	EndDateTime time.Time `hl7:"POS=3"`
	PointOfCare string `hl7:"POS=4"`
	Room string `hl7:"POS=5"`
	Bed string `hl7:"POS=6"`
	Facility HD `hl7:"POS=7"`
	LocationStatus string `hl7:"POS=8"`
	PersonLocationType string `hl7:"POS=9"`
	Building string `hl7:"POS=10"`
	Floor string `hl7:"POS=11"`
}

// CM_CSU - Channel Sensitivity/units
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_CSU
type CM_CSU struct {
	Sensitivity *int `hl7:"POS=1"`
	UnitIdentifier string `hl7:"POS=2"`
	UnitText string `hl7:"POS=3"`
	UnitsNameOfCodingSystem string `hl7:"POS=4"`
	UnitsAlternateIdentifier string `hl7:"POS=5"`
	UnitsAlternateText string `hl7:"POS=6"`
	UnitsNameOfAlternateCodingSystem string `hl7:"POS=7"`
}

// CM_EIP - Parent Order
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_EIP
type CM_EIP struct {
	ParentsPlacerOrderNumber EI `hl7:"POS=1"`
	ParentsFillerOrderNumber EI `hl7:"POS=2"`
}

// CM_RANGE - Wertebereich
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_RANGE
type CM_RANGE struct {
	LowValue string `hl7:"POS=1"`
	HighValue string `hl7:"POS=2"`
}

// MO - Money
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/MO
type MO struct {
	Quantity *int `hl7:"POS=1"`
	Denomination string `hl7:"POS=2"`
}

// CM_DLT - Delta Check
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_DLT
type CM_DLT struct {
	Range CM_RANGE `hl7:"POS=1"`
	NumericThreshold *int `hl7:"POS=2"`
	Change string `hl7:"POS=3"`
	LengthOfTimeDays *int `hl7:"POS=4"`
}

// LA1 - Location With Address Information (variant 1)
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/LA1
type LA1 struct {
	PointOfCare string `hl7:"POS=1"`
	Room string `hl7:"POS=2"`
	Bed string `hl7:"POS=3"`
	Facility HD `hl7:"POS=4"`
	LocationStatus string `hl7:"POS=5"`
	PersonLocationType string `hl7:"POS=6"`
	Building string `hl7:"POS=7"`
	Floor string `hl7:"POS=8"`
	Address AD `hl7:"POS=9"`
}

// CM_MSG - Message Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_MSG
type CM_MSG struct {
	MessageType string `hl7:"POS=1"`
	TriggerEvent string `hl7:"POS=2"`
}

// CQ - Composite Quantity With Units
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CQ
type CQ struct {
	Quantity *int `hl7:"POS=1"`
	Units string `hl7:"POS=2"`
}

// CM_CCD - Charge Time
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_CCD
type CM_CCD struct {
	WhenToChargeCode string `hl7:"POS=1"`
	DateTime time.Time `hl7:"POS=2"`
}

// CM_WVI - Channel Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_WVI
type CM_WVI struct {
	ChannelNumber *int `hl7:"POS=1"`
	ChannelName string `hl7:"POS=2"`
}

// NA - Numeric Array
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/NA
type NA struct {
	Value1 *int `hl7:"POS=1"`
	Value2 *int `hl7:"POS=2"`
	Value3 *int `hl7:"POS=3"`
	Value4 *int `hl7:"POS=4"`
}

// FC - Financial Class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// CM_PI - Person Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_PI
type CM_PI struct {
	IDNumber string `hl7:"POS=1"`
	TypeOfIDNumber string `hl7:"POS=2"`
	OtherQualifyingInfo string `hl7:"POS=3"`
}

// CF - Coded Element With Formatted Values
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CF
type CF struct {
	Identifier string `hl7:"POS=1"`
	FormattedText string `hl7:"POS=2"`
	NameOfCodingSystem string `hl7:"POS=3"`
	AlternateIdentifier string `hl7:"POS=4"`
	AlternateFormattedText string `hl7:"POS=5"`
	NameOfAlternateCodingSystem string `hl7:"POS=6"`
}

// SCV - Scheduling Class Value Pair
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/SCV
type SCV struct {
	ParameterClass string `hl7:"POS=1"`
	ParameterValue string `hl7:"POS=2"`
}

// CD - Channel Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CD
type CD struct {
	ChannelIdentifier CM_WVI `hl7:"POS=1"`
	WaveformSource CM_CD_ELECTRODE `hl7:"POS=2"`
	ChannelSensitivityUnits CM_CSU `hl7:"POS=3"`
	CalibrationParameters CM_CCP `hl7:"POS=4"`
	SamplingFrequency *int `hl7:"POS=5"`
	MinimumMaximumDataValues CM_MDV `hl7:"POS=6"`
}

// CM_SPD - Specialty
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_SPD
type CM_SPD struct {
	SpecialtyName string `hl7:"POS=1"`
	GoverningBoard string `hl7:"POS=2"`
	EligibleOrCertified string `hl7:"POS=3"`
	DateOfCertification time.Time `hl7:"POS=4;ATR=date"`
}

// CE - Coded Element
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CE
type CE struct {
	Identifier string `hl7:"POS=1"`
	Text string `hl7:"POS=2"`
	NameOfCodingSystem string `hl7:"POS=3"`
	AlternateIdentifier string `hl7:"POS=4"`
	AlternateText string `hl7:"POS=5"`
	NameOfAlternateCodingSystem string `hl7:"POS=6"`
}

// RI - Repeat Interval
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/RI
type RI struct {
	RepeatPattern string `hl7:"POS=1"`
	ExplicitTimeInterval string `hl7:"POS=2"`
}

// LA2 - Location With Address Information (variant 2)
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/LA2
type LA2 struct {
	PointOfCare string `hl7:"POS=1"`
	Room string `hl7:"POS=2"`
	Bed string `hl7:"POS=3"`
	Facility HD `hl7:"POS=4"`
	LocationStatus string `hl7:"POS=5"`
	PersonLocationType string `hl7:"POS=6"`
	Building string `hl7:"POS=7"`
	Floor string `hl7:"POS=8"`
	StreetAddress string `hl7:"POS=9"`
	OtherDesignation string `hl7:"POS=10"`
	City string `hl7:"POS=11"`
	StateOrProvince string `hl7:"POS=12"`
	ZipOrPostalCode string `hl7:"POS=13"`
	Country string `hl7:"POS=14"`
	AddressType string `hl7:"POS=15"`
	OtherGeographicDesignation string `hl7:"POS=16"`
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/XCN
type XCN struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName string `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	MiddleInitialOrName string `hl7:"POS=4"`
	Suffix string `hl7:"POS=5"`
	Prefix string `hl7:"POS=6"`
	Degree string `hl7:"POS=7"`
	SourceTable string `hl7:"POS=8"`
	AssigningAuthority HD `hl7:"POS=9"`
	NameType string `hl7:"POS=10"`
	IdentifierCheckDigit string `hl7:"POS=11"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=12"`
	IdentifierTypeCode string `hl7:"POS=13"`
	AssigningFacilityID HD `hl7:"POS=14"`
}

// PTS - Authentication person, time stamp 
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/PTS
type PTS struct {
	AuthenticationPerson XCN `hl7:"POS=1"`
	AuthenticationTimeStamp time.Time `hl7:"POS=2"`
}

// QSC - Query Selection Criteria
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/QSC
type QSC struct {
	NameOfField string `hl7:"POS=1"`
	RelationalOperator string `hl7:"POS=2"`
	Value string `hl7:"POS=3"`
	RelationalConjunction string `hl7:"POS=4"`
}

// ED - Encapsulated Data
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/ED
type ED struct {
	SourceApplication HD `hl7:"POS=1"`
	TypeOfData string `hl7:"POS=2"`
	DataSubtype string `hl7:"POS=3"`
	Encoding string `hl7:"POS=4"`
	Data string `hl7:"POS=5"`
}

// PN - Person Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/PN
type PN struct {
	FamilyName string `hl7:"POS=1"`
	GivenName string `hl7:"POS=2"`
	MiddleInitialOrName string `hl7:"POS=3"`
	Suffix string `hl7:"POS=4"`
	Prefix string `hl7:"POS=5"`
	Degree string `hl7:"POS=6"`
}

// DLN - Driver's License Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/DLN
type DLN struct {
	DriversLicenseNumber string `hl7:"POS=1"`
	IssuingStateProvinceCountry string `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
}

// CM_DLD - Discharge Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_DLD
type CM_DLD struct {
	DischargeLocation string `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// JCC - Job Code Class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/JCC
type JCC struct {
	JobCode string `hl7:"POS=1"`
	JobClass string `hl7:"POS=2"`
}

// CM_AUI - Authorization Information
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_AUI
type CM_AUI struct {
	AuthorizationNumber string `hl7:"POS=1"`
	Date time.Time `hl7:"POS=2;ATR=date"`
	Source string `hl7:"POS=3"`
}

// CM_OCD - Occurrence
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_OCD
type CM_OCD struct {
	OccurrenceCode CE `hl7:"POS=1"`
	OccurrenceDate time.Time `hl7:"POS=2;ATR=date"`
}

// DR - Date Time Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/DR
type DR struct {
	RangeStartDateTime time.Time `hl7:"POS=1"`
	RangeEndDateTime time.Time `hl7:"POS=2"`
}

// CNE - Coded with No Exceptions
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CNE
type CNE struct {
	Identifier string `hl7:"POS=1"`
	Text string `hl7:"POS=2"`
	NameOfCodingSystem string `hl7:"POS=3"`
	AlternateIdentifier string `hl7:"POS=4"`
	AlternateText string `hl7:"POS=5"`
	NameOfAlternateCodingSystem string `hl7:"POS=6"`
	CodingSystemVersionID string `hl7:"POS=7"`
	AlternateCodingSystemVersionID string `hl7:"POS=8"`
	OriginalText string `hl7:"POS=9"`
}

// CN - Composite ID Number And Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CN
type CN struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName string `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	MiddleInitialOrName string `hl7:"POS=4"`
	Suffix string `hl7:"POS=5"`
	Prefix string `hl7:"POS=6"`
	Degree string `hl7:"POS=7"`
	SourceTable string `hl7:"POS=8"`
	AssigningAuthority string `hl7:"POS=9"`
}

// CP - Composite Price
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CP
type CP struct {
	Price MO `hl7:"POS=1"`
	PriceType string `hl7:"POS=2"`
	FromValue *int `hl7:"POS=3"`
	ToValue *int `hl7:"POS=4"`
	RangeUnits CE `hl7:"POS=5"`
	RangeType string `hl7:"POS=6"`
}

// CM_DTN - Day Type And Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_DTN
type CM_DTN struct {
	DayType string `hl7:"POS=1"`
	NumberOfDays *int `hl7:"POS=2"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/XTN
type XTN struct {
	TelephoneNumber string `hl7:"POS=1"`
	TelecommunicationUseCode string `hl7:"POS=2"`
	TelecommunicationEquipmentType string `hl7:"POS=3"`
	EmailAddress string `hl7:"POS=4"`
	CountryCode *int `hl7:"POS=5"`
	AreaCityCode *int `hl7:"POS=6"`
	PhoneNumber *int `hl7:"POS=7"`
	Extension *int `hl7:"POS=8"`
	AnyText string `hl7:"POS=9"`
}

// CX - Extended Composite ID With Check Digit
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CX
type CX struct {
	ID string `hl7:"POS=1"`
	CheckDigit string `hl7:"POS=2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=3"`
	AssigningAuthority HD `hl7:"POS=4"`
	IdentifierTypeCode string `hl7:"POS=5"`
	AssigningFacility HD `hl7:"POS=6"`
}

// CM_DIN - Activation Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_DIN
type CM_DIN struct {
	Date time.Time `hl7:"POS=1"`
	InstitutionName CE `hl7:"POS=2"`
}

// CM_OSD - Order Sequence
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_OSD
type CM_OSD struct {
	SequenceResultsFlag string `hl7:"POS=1"`
	PlacerOrderNumberEntityIdentifier string `hl7:"POS=2"`
	PlacerOrderNumberNamespaceID string `hl7:"POS=3"`
	FillerOrderNumberEntityIdentifier string `hl7:"POS=4"`
	FillerOrderNumberNamespaceID string `hl7:"POS=5"`
	SequenceConditionValue string `hl7:"POS=6"`
	MaximumNumberOfRepeats *int `hl7:"POS=7"`
	PlacerOrderNumberUniversalID string `hl7:"POS=8"`
	PlacerOrderNumberUniversalIDType string `hl7:"POS=9"`
	FillerOrderNumberUniversalID string `hl7:"POS=10"`
	FillerOrderNumberUniversalIDType string `hl7:"POS=11"`
}

// CM_MDV - Minimum/maximum Data Values
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_MDV
type CM_MDV struct {
	MinimumValue *int `hl7:"POS=1"`
	MaximumValue *int `hl7:"POS=2"`
}

// CM_VR - Value Qualifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_VR
type CM_VR struct {
	FirstDataCodeValue string `hl7:"POS=1"`
	LastDataCodeValue string `hl7:"POS=2"`
}

// RCD - Row Column Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/RCD
type RCD struct {
	Hl7ItemNumber string `hl7:"POS=1"`
	Hl7DateType string `hl7:"POS=2"`
	MaximumColumnWidth *int `hl7:"POS=3"`
}

// AD - Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/AD
type AD struct {
	StreetAddress string `hl7:"POS=1"`
	OtherDesignation string `hl7:"POS=2"`
	City string `hl7:"POS=3"`
	StateOrProvince string `hl7:"POS=4"`
	ZipOrPostalCode string `hl7:"POS=5"`
	Country string `hl7:"POS=6"`
	AddressType string `hl7:"POS=7"`
	OtherGeographicDesignation string `hl7:"POS=8"`
}

// CK - Composite ID With Check Digit
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CK
type CK struct {
	IDNumber *int `hl7:"POS=1"`
	CheckDigit string `hl7:"POS=2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=3"`
	AssigningAuthority HD `hl7:"POS=4"`
}

// CM_OSP - Occurrence Span
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_OSP
type CM_OSP struct {
	OccurrenceSpanCode CE `hl7:"POS=1"`
	OccurrenceSpanStartDate time.Time `hl7:"POS=2;ATR=date"`
	OccurrenceSpanStopDate time.Time `hl7:"POS=3;ATR=date"`
}

// PT - Processing Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/PT
type PT struct {
	ProcessingID string `hl7:"POS=1"`
	ProcessingMode string `hl7:"POS=2"`
}

// CM_CCP - Channel Calibration Parameters
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_CCP
type CM_CCP struct {
	SensitivityCorrectionFactor *int `hl7:"POS=1"`
	Baseline *int `hl7:"POS=2"`
	TimeSkew *int `hl7:"POS=3"`
}

// CM_PTA - Policy Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_PTA
type CM_PTA struct {
	PolicyType string `hl7:"POS=1"`
	AmountClass string `hl7:"POS=2"`
	Amount *int `hl7:"POS=3"`
}

// CM_PIP - Privileges
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_PIP
type CM_PIP struct {
	Privilege CE `hl7:"POS=1"`
	PrivilegeClass CE `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
	ActivationDate time.Time `hl7:"POS=4;ATR=date"`
}

// CM_PEN - Penalty
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_PEN
type CM_PEN struct {
	PenaltyType string `hl7:"POS=1"`
	PenaltyAmount *int `hl7:"POS=2"`
}

// TQ - Timing Quantity
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/TQ
type TQ struct {
	Quantity CQ `hl7:"POS=1"`
	Interval RI `hl7:"POS=2"`
	Duration string `hl7:"POS=3"`
	StartDateTime time.Time `hl7:"POS=4"`
	EndDateTime time.Time `hl7:"POS=5"`
	Priority string `hl7:"POS=6"`
	Condition string `hl7:"POS=7"`
	Text string `hl7:"POS=8"`
	Conjunction string `hl7:"POS=9"`
	OrderSequencing CM_OSD `hl7:"POS=10"`
}

// FN - Family + Last Name Prefix
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/FN
type FN struct {
	FamilyName string `hl7:"POS=1"`
	LastNamePrefix string `hl7:"POS=2"`
}

// PPN - Performing Person Time Stamp
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/PPN
type PPN struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName string `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	MiddleInitialOrName string `hl7:"POS=4"`
	Suffix string `hl7:"POS=5"`
	Prefix string `hl7:"POS=6"`
	Degree string `hl7:"POS=7"`
	SourceTable string `hl7:"POS=8"`
	AssigningAuthority HD `hl7:"POS=9"`
	NameTypeCode string `hl7:"POS=10"`
	IdentifierCheckDigit string `hl7:"POS=11"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=12"`
	IdentifierTypeCode string `hl7:"POS=13"`
	AssigningFacility HD `hl7:"POS=14"`
	DateTimeActionPerformed time.Time `hl7:"POS=15"`
}

// XON - Extended Composite Name And ID For Organizations
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/XON
type XON struct {
	OrganizationName string `hl7:"POS=1"`
	OrganizationNameTypeCode string `hl7:"POS=2"`
	IDNumber *int `hl7:"POS=3"`
	CheckDigit string `hl7:"POS=4"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=5"`
	AssigningAuthority HD `hl7:"POS=6"`
	IdentifierTypeCode string `hl7:"POS=7"`
	AssigningFacilityID HD `hl7:"POS=8"`
}

// QIP - Query Input Parameter List
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/QIP
type QIP struct {
	FieldName string `hl7:"POS=1"`
	Value1value2value3 string `hl7:"POS=2"`
}

// XPN - Extended Person Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/XPN
type XPN struct {
	FamilyName string `hl7:"POS=1"`
	GivenName string `hl7:"POS=2"`
	MiddleInitialOrName string `hl7:"POS=3"`
	Suffix string `hl7:"POS=4"`
	Prefix string `hl7:"POS=5"`
	Degree string `hl7:"POS=6"`
	NameTypeCode string `hl7:"POS=7"`
	NameRepresentationCode string `hl7:"POS=8"`
}

// CM_RMC - Room Coverage
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_RMC
type CM_RMC struct {
	RoomType string `hl7:"POS=1"`
	AmountType string `hl7:"POS=2"`
	CoverageAmount *int `hl7:"POS=3"`
}

// CM_PRL - Parent Result Link
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_PRL
type CM_PRL struct {
	Obx3ObservationIdentifierOfParentResult CE `hl7:"POS=1"`
	Obx4SubIDOfParentResult string `hl7:"POS=2"`
	PartOfObx5ObservationResultFromParent string `hl7:"POS=3"`
}

// CM_SPS - Specimen Source
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_SPS
type CM_SPS struct {
	SpecimenSourceNameOrCode CE `hl7:"POS=1"`
	Additives string `hl7:"POS=2"`
	Freetext string `hl7:"POS=3"`
	BodySite CE `hl7:"POS=4"`
	SiteModifier CE `hl7:"POS=5"`
	CollectionModifierMethodCode CE `hl7:"POS=6"`
}

// TS - Time Stamp
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/TS
type TS struct {
	TimeOfAnEvent string `hl7:"POS=1"`
}

// RP - Reference Pointer
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/RP
type RP struct {
	Pointer string `hl7:"POS=1"`
	ApplicationID HD `hl7:"POS=2"`
	TypeOfData string `hl7:"POS=3"`
	Subtype string `hl7:"POS=4"`
}

// CM_DDI - Daily Deductible
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_DDI
type CM_DDI struct {
	DelayDays *int `hl7:"POS=1"`
	Amount *int `hl7:"POS=2"`
	NumberOfDays *int `hl7:"POS=3"`
}

// ELD - Error
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/ELD
type ELD struct {
	SegmentID string `hl7:"POS=1"`
	Sequence *int `hl7:"POS=2"`
	FieldPosition *int `hl7:"POS=3"`
	CodeIdentifyingError CE `hl7:"POS=4"`
}

// CM_ELD - Error
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_ELD
type CM_ELD struct {
	SegmentID string `hl7:"POS=1"`
	Sequence *int `hl7:"POS=2"`
	FieldPosition *int `hl7:"POS=3"`
	CodeIdentifyingError CE `hl7:"POS=4"`
}

// VID - Version Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/VID
type VID struct {
	VersionID string `hl7:"POS=1"`
	InternationalizationCode CE `hl7:"POS=2"`
	InternationalVersionID CE `hl7:"POS=3"`
}

// HD - Hierarchic Designator
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/HD
type HD struct {
	NamespaceID string `hl7:"POS=1"`
	UniversalID string `hl7:"POS=2"`
	UniversalIDType string `hl7:"POS=3"`
}

// CM_CD_ELECTRODE - Electrode Parameter
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_CD_ELECTRODE
type CM_CD_ELECTRODE struct {
	SourceName1 string `hl7:"POS=1"`
	SourceName2 string `hl7:"POS=2"`
}

// CM_PCF - Pre-certification Required
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_PCF
type CM_PCF struct {
	PreCertificationPatientType string `hl7:"POS=1"`
	PreCertificationRequired string `hl7:"POS=2"`
	PreCertificationWindow time.Time `hl7:"POS=3"`
}

// CM_ABS_RANGE - Absolute Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_ABS_RANGE
type CM_ABS_RANGE struct {
	Range CM_RANGE `hl7:"POS=1"`
	NumericChange *int `hl7:"POS=2"`
	PercentPerChange *int `hl7:"POS=3"`
	Days *int `hl7:"POS=4"`
}

// EI - Entity Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"POS=1"`
	NamespaceID string `hl7:"POS=2"`
	UniversalID string `hl7:"POS=3"`
	UniversalIDType string `hl7:"POS=4"`
}

// PL - Person Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/PL
type PL struct {
	PointOfCare string `hl7:"POS=1"`
	Room string `hl7:"POS=2"`
	Bed string `hl7:"POS=3"`
	Facility HD `hl7:"POS=4"`
	LocationStatus string `hl7:"POS=5"`
	PersonLocationType string `hl7:"POS=6"`
	Building string `hl7:"POS=7"`
	Floor string `hl7:"POS=8"`
	LocationType string `hl7:"POS=9"`
}

// MA - Multiplexed Array
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/MA
type MA struct {
	Sample1FromChannel1 *int `hl7:"POS=1"`
	Sample1FromChannel2 *int `hl7:"POS=2"`
	Sample1FromChannel3 *int `hl7:"POS=3"`
	Sample2FromChannel1 *int `hl7:"POS=4"`
	Sample2FromChannel2 *int `hl7:"POS=5"`
	Sample2FromChannel3 *int `hl7:"POS=6"`
}

// CM_MOC - Charge To Practise
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_MOC
type CM_MOC struct {
	DollarAmount MO `hl7:"POS=1"`
	ChargeCode CE `hl7:"POS=2"`
}

// CM_UVC - Value Code And Amount
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/CM_UVC
type CM_UVC struct {
	ValueCode string `hl7:"POS=1"`
	ValueAmount *int `hl7:"POS=2"`
}

// SN - Structured Numeric
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/SN
type SN struct {
	Comparator string `hl7:"POS=1"`
	Num1 *int `hl7:"POS=2"`
	SeparatorOrSuffix string `hl7:"POS=3"`
	Num2 *int `hl7:"POS=4"`
}

// XAD - Extended Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/XAD
type XAD struct {
	StreetAddress string `hl7:"POS=1"`
	OtherDesignation string `hl7:"POS=2"`
	City string `hl7:"POS=3"`
	StateOrProvince string `hl7:"POS=4"`
	ZipOrPostalCode string `hl7:"POS=5"`
	Country string `hl7:"POS=6"`
	AddressType string `hl7:"POS=7"`
	OtherGeographicDesignation string `hl7:"POS=8"`
	CountyParishCode string `hl7:"POS=9"`
	CensusTract string `hl7:"POS=10"`
}

// VH - Visiting Hours
// https://hl7-definition-staging.caristix.com/v2/HL7v2.3/DataTypes/VH
type VH struct {
	StartDayRange string `hl7:"POS=1"`
	EndDayRange string `hl7:"POS=2"`
	StartHourRange string `hl7:"POS=3"` // TODO - handle TM type parsing
	EndHourRange string `hl7:"POS=4"` // TODO - handle TM type parsing
}

