package hl7v26

import "time"

// MO - Money
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/MO
type MO struct {
	Quantity *int `hl7:"POS=1"`
	Denomination string `hl7:"POS=2"`
}

// ERL - Error Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/ERL
type ERL struct {
	SegmentID string `hl7:"POS=1"`
	SegmentSequence *int `hl7:"POS=2"`
	FieldPosition *int `hl7:"POS=3"`
	FieldRepetition *int `hl7:"POS=4"`
	ComponentNumber *int `hl7:"POS=5"`
	SubComponentNumber *int `hl7:"POS=6"`
}

// LA1 - Location with Address Variation 1
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/LA1
type LA1 struct {
	PointOfCare string `hl7:"POS=1"`
	Room string `hl7:"POS=2"`
	Bed string `hl7:"POS=3"`
	Facility HD `hl7:"POS=4"`
	LocationStatus string `hl7:"POS=5"`
	PatientLocationType string `hl7:"POS=6"`
	Building string `hl7:"POS=7"`
	Floor string `hl7:"POS=8"`
	Address AD `hl7:"POS=9"`
}

// CQ - Composite Quantity with Units
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CQ
type CQ struct {
	Quantity *int `hl7:"POS=1"`
	Units CWE `hl7:"POS=2"`
}

// NA - Numeric Array
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/NA
type NA struct {
	Value1 *int `hl7:"POS=1"`
	Value2 *int `hl7:"POS=2"`
	Value3 *int `hl7:"POS=3"`
	Value4 *int `hl7:"POS=4"`
}

// FC - Financial Class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/FC
type FC struct {
	FinancialClassCode string `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// SAD - Street Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/SAD
type SAD struct {
	StreetOrMailingAddress string `hl7:"POS=1"`
	StreetName string `hl7:"POS=2"`
	DwellingNumber string `hl7:"POS=3"`
}

// OSD - Order Sequence Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/OSD
type OSD struct {
	SequenceResultsFlag string `hl7:"POS=1"`
	PlacerOrderNumberEntityIdentifier string `hl7:"POS=2"`
	PlacerOrderNumberNamespaceID string `hl7:"POS=3"`
	FillerOrderNumberEntityIdentifier string `hl7:"POS=4"`
	FillerOrderNumberNamespaceID string `hl7:"POS=5"`
	SequenceConditionValue string `hl7:"POS=6"`
	MaximumNumberOfRepeats *int `hl7:"POS=7"`
	PlacerOrderNumberUniversalID string `hl7:"POS=8"`
	PlacerOrderNumberUniversalIdtype string `hl7:"POS=9"`
	FillerOrderNumberUniversalID string `hl7:"POS=10"`
	FillerOrderNumberUniversalIdtype string `hl7:"POS=11"`
}

// DIN - Date and Institution Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/DIN
type DIN struct {
	Date time.Time `hl7:"POS=1"`
	InstitutionName CWE `hl7:"POS=2"`
}

// CF - Coded Element with Formatted Values
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CF
type CF struct {
	Identifier string `hl7:"POS=1"`
	FormattedText string `hl7:"POS=2"`
	NameOfCodingSystem string `hl7:"POS=3"`
	AlternateIdentifier string `hl7:"POS=4"`
	AlternateFormattedText string `hl7:"POS=5"`
	NameOfAlternateCodingSystem string `hl7:"POS=6"`
}

// EIP - Entity Identifier Pair
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/EIP
type EIP struct {
	PlacerAssignedIdentifier EI `hl7:"POS=1"`
	FillerAssignedIdentifier EI `hl7:"POS=2"`
}

// SCV - Scheduling Class Value Pair
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/SCV
type SCV struct {
	ParameterClass CWE `hl7:"POS=1"`
	ParameterValue string `hl7:"POS=2"`
}

// CD - Channel Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CD
type CD struct {
	ChannelIdentifier WVI `hl7:"POS=1"`
	WaveformSource WVS `hl7:"POS=2"`
	ChannelSensitivityAndUnits CSU `hl7:"POS=3"`
	ChannelCalibrationParameters CCP `hl7:"POS=4"`
	ChannelSamplingFrequency *int `hl7:"POS=5"`
	MinimumAndMaximumDataValues NR `hl7:"POS=6"`
}

// CE - Coded Element
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CE
type CE struct {
}

// PIP - Practitioner Institutional Privileges
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/PIP
type PIP struct {
	Privilege CWE `hl7:"POS=1"`
	PrivilegeClass CWE `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
	ActivationDate time.Time `hl7:"POS=4;ATR=date"`
	Facility EI `hl7:"POS=5"`
}

// RI - Repeat Interval
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/RI
type RI struct {
	RepeatPattern string `hl7:"POS=1"`
	ExplicitTimeInterval string `hl7:"POS=2"`
}

// LA2 - Location with Address Variation 2
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/LA2
type LA2 struct {
	PointOfCare string `hl7:"POS=1"`
	Room string `hl7:"POS=2"`
	Bed string `hl7:"POS=3"`
	Facility HD `hl7:"POS=4"`
	LocationStatus string `hl7:"POS=5"`
	PatientLocationType string `hl7:"POS=6"`
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

// XCN - Extended Composite ID Number and Name for Persons
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/XCN
type XCN struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName FN `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=4"`
	Suffix string `hl7:"POS=5"`
	Prefix string `hl7:"POS=6"`
	Degree string `hl7:"POS=7"`
	SourceTable string `hl7:"POS=8"`
	AssigningAuthority HD `hl7:"POS=9"`
	NameTypeCode string `hl7:"POS=10"`
	IdentifierCheckDigit string `hl7:"POS=11"`
	CheckDigitScheme string `hl7:"POS=12"`
	IdentifierTypeCode string `hl7:"POS=13"`
	AssigningFacility HD `hl7:"POS=14"`
	NameRepresentationCode string `hl7:"POS=15"`
	NameContext CWE `hl7:"POS=16"`
	NameValidityRange DR `hl7:"POS=17"`
	NameAssemblyOrder string `hl7:"POS=18"`
	EffectiveDate time.Time `hl7:"POS=19"`
	ExpirationDate time.Time `hl7:"POS=20"`
	ProfessionalSuffix string `hl7:"POS=21"`
	AssigningJurisdiction CWE `hl7:"POS=22"`
	AssigningAgencyOrDepartment CWE `hl7:"POS=23"`
}

// QSC - Query Selection Criteria
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/QSC
type QSC struct {
	SegmentFieldName string `hl7:"POS=1"`
	RelationalOperator string `hl7:"POS=2"`
	Value string `hl7:"POS=3"`
	RelationalConjunction string `hl7:"POS=4"`
}

// ED - Encapsulated Data
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/ED
type ED struct {
	SourceApplication HD `hl7:"POS=1"`
	TypeOfData string `hl7:"POS=2"`
	DataSubtype string `hl7:"POS=3"`
	Encoding string `hl7:"POS=4"`
	Data string `hl7:"POS=5"`
}

// DLD - Discharge Location and Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/DLD
type DLD struct {
	DischargeToLocation CWE `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// PLN - Practitioner License or Other ID Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/PLN
type PLN struct {
	IDNumber string `hl7:"POS=1"`
	TypeOfIDNumber string `hl7:"POS=2"`
	StateOtherQualifyingInformation string `hl7:"POS=3"`
	ExpirationDate time.Time `hl7:"POS=4;ATR=date"`
}

// CNN - Composite ID Number and Name Simplified
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CNN
type CNN struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName string `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=4"`
	Suffix string `hl7:"POS=5"`
	Prefix string `hl7:"POS=6"`
	Degree string `hl7:"POS=7"`
	SourceTable string `hl7:"POS=8"`
	AssigningAuthorityNamespaceID string `hl7:"POS=9"`
	AssigningAuthorityUniversalID string `hl7:"POS=10"`
	AssigningAuthorityUniversalIdtype string `hl7:"POS=11"`
}

// OCD - Occurrence Code and Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/OCD
type OCD struct {
	OccurrenceCode CNE `hl7:"POS=1"`
	OccurrenceDate time.Time `hl7:"POS=2;ATR=date"`
}

// DLN - Driver_s License Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/DLN
type DLN struct {
	LicenseNumber string `hl7:"POS=1"`
	IssuingStateProvinceCountry string `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
}

// CSU - Channel Sensitivity
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CSU
type CSU struct {
	ChannelSensitivity *int `hl7:"POS=1"`
	UnitOfMeasureIdentifier string `hl7:"POS=2"`
	UnitOfMeasureDescription string `hl7:"POS=3"`
	UnitOfMeasureCodingSystem string `hl7:"POS=4"`
	AlternateUnitOfMeasureIdentifier string `hl7:"POS=5"`
	AlternateUnitOfMeasureDescription string `hl7:"POS=6"`
	AlternateUnitOfMeasureCodingSystem string `hl7:"POS=7"`
}

// JCC - Job Code/Class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/JCC
type JCC struct {
	JobCode string `hl7:"POS=1"`
	JobClass string `hl7:"POS=2"`
	JobDescriptionText string `hl7:"POS=3"`
}

// DTN - Day Type and Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/DTN
type DTN struct {
	DayType string `hl7:"POS=1"`
	NumberOfDays *int `hl7:"POS=2"`
}

// WVS - Waveform Source
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/WVS
type WVS struct {
	SourceOneName string `hl7:"POS=1"`
	SourceTwoName string `hl7:"POS=2"`
}

// MSG - Message Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/MSG
type MSG struct {
	MessageCode string `hl7:"POS=1"`
	TriggerEvent string `hl7:"POS=2"`
	MessageStructure string `hl7:"POS=3"`
}

// DR - Date/Time Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/DR
type DR struct {
	RangeStartDateTime time.Time `hl7:"POS=1"`
	RangeEndDateTime time.Time `hl7:"POS=2"`
}

// CNE - Coded with No Exceptions
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CNE
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

// CP - Composite Price
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CP
type CP struct {
	Price MO `hl7:"POS=1"`
	PriceType string `hl7:"POS=2"`
	FromValue *int `hl7:"POS=3"`
	ToValue *int `hl7:"POS=4"`
	RangeUnits CWE `hl7:"POS=5"`
	RangeType string `hl7:"POS=6"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/XTN
type XTN struct {
	TelecommunicationUseCode string `hl7:"POS=2"`
	TelecommunicationEquipmentType string `hl7:"POS=3"`
	CommunicationAddress string `hl7:"POS=4"`
	CountryCode *int `hl7:"POS=5"`
	AreaCityCode *int `hl7:"POS=6"`
	LocalNumber *int `hl7:"POS=7"`
	Extension *int `hl7:"POS=8"`
	AnyText string `hl7:"POS=9"`
	ExtensionPrefix string `hl7:"POS=10"`
	SpeedDialCode string `hl7:"POS=11"`
	UnformattedTelephoneNumber string `hl7:"POS=12"`
	EffectiveStartDate time.Time `hl7:"POS=13"`
	ExpirationDate time.Time `hl7:"POS=14"`
	ExpirationReason CWE `hl7:"POS=15"`
	ProtectionCode CWE `hl7:"POS=16"`
	SharedTelecommunicationIdentifier EI `hl7:"POS=17"`
	PreferenceOrder *int `hl7:"POS=18"`
}

// MOP - Money or Percentage
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/MOP
type MOP struct {
	MoneyOrPercentageIndicator string `hl7:"POS=1"`
	MoneyOrPercentageQuantity *int `hl7:"POS=2"`
	CurrencyDenomination string `hl7:"POS=3"`
}

// CX - Extended Composite ID with Check Digit
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CX
type CX struct {
	IDNumber string `hl7:"POS=1"`
	IdentifierCheckDigit string `hl7:"POS=2"`
	CheckDigitScheme string `hl7:"POS=3"`
	AssigningAuthority HD `hl7:"POS=4"`
	IdentifierTypeCode string `hl7:"POS=5"`
	AssigningFacility HD `hl7:"POS=6"`
	EffectiveDate time.Time `hl7:"POS=7;ATR=date"`
	ExpirationDate time.Time `hl7:"POS=8;ATR=date"`
	AssigningJurisdiction CWE `hl7:"POS=9"`
	AssigningAgencyOrDepartment CWE `hl7:"POS=10"`
}

// ICD - Insurance Certification Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/ICD
type ICD struct {
	CertificationPatientType string `hl7:"POS=1"`
	CertificationRequired string `hl7:"POS=2"`
	DateTimeCertificationRequired time.Time `hl7:"POS=3"`
}

// CCP - Channel Calibration Parameters
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CCP
type CCP struct {
	ChannelCalibrationSensitivityCorrectionFactor *int `hl7:"POS=1"`
	ChannelCalibrationBaseline *int `hl7:"POS=2"`
	ChannelCalibrationTimeSkew *int `hl7:"POS=3"`
}

// NDL - Name with Date and Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/NDL
type NDL struct {
	Name CNN `hl7:"POS=1"`
	StartDateTime time.Time `hl7:"POS=2"`
	EndDateTime time.Time `hl7:"POS=3"`
	PointOfCare string `hl7:"POS=4"`
	Room string `hl7:"POS=5"`
	Bed string `hl7:"POS=6"`
	Facility HD `hl7:"POS=7"`
	LocationStatus string `hl7:"POS=8"`
	PatientLocationType string `hl7:"POS=9"`
	Building string `hl7:"POS=10"`
	Floor string `hl7:"POS=11"`
}

// RCD - Row Column Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/RCD
type RCD struct {
	SegmentFieldName string `hl7:"POS=1"`
	Hl7DataType string `hl7:"POS=2"`
	MaximumColumnWidth *int `hl7:"POS=3"`
}

// AD - Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/AD
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

// AUI - Authorization Information
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/AUI
type AUI struct {
	AuthorizationNumber string `hl7:"POS=1"`
	Date time.Time `hl7:"POS=2;ATR=date"`
	Source string `hl7:"POS=3"`
}

// MOC - Money and Code
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/MOC
type MOC struct {
	MonetaryAmount MO `hl7:"POS=1"`
	ChargeCode CWE `hl7:"POS=2"`
}

// PT - Processing Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/PT
type PT struct {
	ProcessingID string `hl7:"POS=1"`
	ProcessingMode string `hl7:"POS=2"`
}

// DDI - Daily Deductible Information
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/DDI
type DDI struct {
	DelayDays *int `hl7:"POS=1"`
	MonetaryAmount MO `hl7:"POS=2"`
	NumberOfDays *int `hl7:"POS=3"`
}

// TQ - Timing Quantity
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/TQ
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
	OrderSequencing OSD `hl7:"POS=10"`
	OccurrenceDuration CWE `hl7:"POS=11"`
	TotalOccurrences *int `hl7:"POS=12"`
}

// VR - Value Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/VR
type VR struct {
	FirstDataCodeValue string `hl7:"POS=1"`
	LastDataCodeValue string `hl7:"POS=2"`
}

// FN - Family Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/FN
type FN struct {
	Surname string `hl7:"POS=1"`
	OwnSurnamePrefix string `hl7:"POS=2"`
	OwnSurname string `hl7:"POS=3"`
	SurnamePrefixFromPartnerSpouse string `hl7:"POS=4"`
	SurnameFromPartnerSpouse string `hl7:"POS=5"`
}

// RMC - Room Coverage
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/RMC
type RMC struct {
	RoomType string `hl7:"POS=1"`
	AmountType string `hl7:"POS=2"`
	CoverageAmount *int `hl7:"POS=3"`
	MoneyOrPercentage MOP `hl7:"POS=4"`
}

// PPN - Performing Person Time Stamp
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/PPN
type PPN struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName FN `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=4"`
	Suffix string `hl7:"POS=5"`
	Prefix string `hl7:"POS=6"`
	Degree string `hl7:"POS=7"`
	SourceTable string `hl7:"POS=8"`
	AssigningAuthority HD `hl7:"POS=9"`
	NameTypeCode string `hl7:"POS=10"`
	IdentifierCheckDigit string `hl7:"POS=11"`
	CheckDigitScheme string `hl7:"POS=12"`
	IdentifierTypeCode string `hl7:"POS=13"`
	AssigningFacility HD `hl7:"POS=14"`
	DateTimeActionPerformed time.Time `hl7:"POS=15"`
	NameRepresentationCode string `hl7:"POS=16"`
	NameContext CWE `hl7:"POS=17"`
	NameValidityRange DR `hl7:"POS=18"`
	NameAssemblyOrder string `hl7:"POS=19"`
	EffectiveDate time.Time `hl7:"POS=20"`
	ExpirationDate time.Time `hl7:"POS=21"`
	ProfessionalSuffix string `hl7:"POS=22"`
	AssigningJurisdiction CWE `hl7:"POS=23"`
	AssigningAgencyOrDepartment CWE `hl7:"POS=24"`
}

// XON - Extended Composite Name and Identification Number for Organizations
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/XON
type XON struct {
	OrganizationName string `hl7:"POS=1"`
	OrganizationNameTypeCode string `hl7:"POS=2"`
	IDNumber *int `hl7:"POS=3"`
	IdentifierCheckDigit *int `hl7:"POS=4"`
	CheckDigitScheme string `hl7:"POS=5"`
	AssigningAuthority HD `hl7:"POS=6"`
	IdentifierTypeCode string `hl7:"POS=7"`
	AssigningFacility HD `hl7:"POS=8"`
	NameRepresentationCode string `hl7:"POS=9"`
	OrganizationIdentifier string `hl7:"POS=10"`
}

// RPT - Repeat Pattern
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/RPT
type RPT struct {
	RepeatPatternCode CWE `hl7:"POS=1"`
	CalendarAlignment string `hl7:"POS=2"`
	PhaseRangeBeginValue *int `hl7:"POS=3"`
	PhaseRangeEndValue *int `hl7:"POS=4"`
	PeriodQuantity *int `hl7:"POS=5"`
	PeriodUnits string `hl7:"POS=6"`
	InstitutionSpecifiedTime string `hl7:"POS=7"`
	Event string `hl7:"POS=8"`
	EventOffsetQuantity *int `hl7:"POS=9"`
	EventOffsetUnits string `hl7:"POS=10"`
	GeneralTimingSpecification string `hl7:"POS=11"`
}

// QIP - Query Input Parameter List
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/QIP
type QIP struct {
	SegmentFieldName string `hl7:"POS=1"`
	Values string `hl7:"POS=2"`
}

// PRL - Parent Result Link
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/PRL
type PRL struct {
	ParentObservationIdentifier CWE `hl7:"POS=1"`
	ParentObservationSubIdentifier string `hl7:"POS=2"`
	ParentObservationValueDescriptor string `hl7:"POS=3"`
}

// DLT - Delta
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/DLT
type DLT struct {
	NormalRange NR `hl7:"POS=1"`
	NumericThreshold *int `hl7:"POS=2"`
	ChangeComputation string `hl7:"POS=3"`
	DaysRetained *int `hl7:"POS=4"`
}

// PTA - Policy Type and Amount
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/PTA
type PTA struct {
	PolicyType string `hl7:"POS=1"`
	AmountClass string `hl7:"POS=2"`
	MoneyOrPercentageQuantity *int `hl7:"POS=3"`
	MoneyOrPercentage MOP `hl7:"POS=4"`
}

// XPN - Extended Person Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/XPN
type XPN struct {
	FamilyName FN `hl7:"POS=1"`
	GivenName string `hl7:"POS=2"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=3"`
	Suffix string `hl7:"POS=4"`
	Prefix string `hl7:"POS=5"`
	Degree string `hl7:"POS=6"`
	NameTypeCode string `hl7:"POS=7"`
	NameRepresentationCode string `hl7:"POS=8"`
	NameContext CWE `hl7:"POS=9"`
	NameValidityRange DR `hl7:"POS=10"`
	NameAssemblyOrder string `hl7:"POS=11"`
	EffectiveDate time.Time `hl7:"POS=12"`
	ExpirationDate time.Time `hl7:"POS=13"`
	ProfessionalSuffix string `hl7:"POS=14"`
}

// SPD - Specialty Description
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/SPD
type SPD struct {
	SpecialtyName string `hl7:"POS=1"`
	GoverningBoard string `hl7:"POS=2"`
	EligibleOrCertified string `hl7:"POS=3"`
	DateOfCertification time.Time `hl7:"POS=4;ATR=date"`
}

// UVC - UB Value Code and Amount
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/UVC
type UVC struct {
	ValueCode CNE `hl7:"POS=1"`
	ValueAmount MO `hl7:"POS=2"`
}

// CWE - Coded with Exceptions
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CWE
type CWE struct {
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

// SRT - Sort Order
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/SRT
type SRT struct {
	SortByField string `hl7:"POS=1"`
	Sequencing string `hl7:"POS=2"`
}

// TS - Time Stamp
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/TS
type TS struct {
}

// RP - Reference Pointer
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/RP
type RP struct {
	Pointer string `hl7:"POS=1"`
	ApplicationID HD `hl7:"POS=2"`
	TypeOfData string `hl7:"POS=3"`
	Subtype string `hl7:"POS=4"`
}

// RFR - Reference Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/RFR
type RFR struct {
	NumericRange NR `hl7:"POS=1"`
	AdministrativeSex string `hl7:"POS=2"`
	AgeRange NR `hl7:"POS=3"`
	GestationalAgeRange NR `hl7:"POS=4"`
	Species string `hl7:"POS=5"`
	RaceSubspecies string `hl7:"POS=6"`
	Conditions string `hl7:"POS=7"`
}

// ELD - Error Location and Description
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/ELD
type ELD struct {
	SegmentID string `hl7:"POS=1"`
	SegmentSequence *int `hl7:"POS=2"`
	FieldPosition *int `hl7:"POS=3"`
	CodeIdentifyingError CWE `hl7:"POS=4"`
}

// OSP - Occurrence Span Code and Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/OSP
type OSP struct {
	OccurrenceSpanCode CNE `hl7:"POS=1"`
	OccurrenceSpanStartDate time.Time `hl7:"POS=2;ATR=date"`
	OccurrenceSpanStopDate time.Time `hl7:"POS=3;ATR=date"`
}

// VID - Version Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/VID
type VID struct {
	VersionID string `hl7:"POS=1"`
	InternationalizationCode CWE `hl7:"POS=2"`
	InternationalVersionID CWE `hl7:"POS=3"`
}

// HD - Hierarchic Designator
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/HD
type HD struct {
	NamespaceID string `hl7:"POS=1"`
	UniversalID string `hl7:"POS=2"`
	UniversalIdtype string `hl7:"POS=3"`
}

// SPS - Specimen Source
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/SPS
type SPS struct {
	SpecimenSourceNameOrCode CWE `hl7:"POS=1"`
	Additives CWE `hl7:"POS=2"`
	SpecimenCollectionMethod string `hl7:"POS=3"`
	BodySite CWE `hl7:"POS=4"`
	SiteModifier CWE `hl7:"POS=5"`
	CollectionMethodModifierCode CWE `hl7:"POS=6"`
	SpecimenRole CWE `hl7:"POS=7"`
}

// EI - Entity Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"POS=1"`
	NamespaceID string `hl7:"POS=2"`
	UniversalID string `hl7:"POS=3"`
	UniversalIdtype string `hl7:"POS=4"`
}

// PL - Person Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/PL
type PL struct {
	PointOfCare string `hl7:"POS=1"`
	Room string `hl7:"POS=2"`
	Bed string `hl7:"POS=3"`
	Facility HD `hl7:"POS=4"`
	LocationStatus string `hl7:"POS=5"`
	PersonLocationType string `hl7:"POS=6"`
	Building string `hl7:"POS=7"`
	Floor string `hl7:"POS=8"`
	LocationDescription string `hl7:"POS=9"`
	ComprehensiveLocationIdentifier EI `hl7:"POS=10"`
	AssigningAuthorityForLocation HD `hl7:"POS=11"`
}

// MA - Multiplexed Array
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/MA
type MA struct {
	SampleYFromChannel1 *int `hl7:"POS=1"`
	SampleYFromChannel2 *int `hl7:"POS=2"`
	SampleYFromChannel3 *int `hl7:"POS=3"`
	SampleYFromChannel4 *int `hl7:"POS=4"`
}

// SN - Structured Numeric
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/SN
type SN struct {
	Comparator string `hl7:"POS=1"`
	Num1 *int `hl7:"POS=2"`
	SeparatorSuffix string `hl7:"POS=3"`
	Num2 *int `hl7:"POS=4"`
}

// CCD - Charge Code and Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/CCD
type CCD struct {
	InvocationEvent string `hl7:"POS=1"`
	DateTime time.Time `hl7:"POS=2"`
}

// WVI - Channel Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/WVI
type WVI struct {
	ChannelNumber *int `hl7:"POS=1"`
	ChannelName string `hl7:"POS=2"`
}

// NR - Numeric Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/NR
type NR struct {
	LowValue *int `hl7:"POS=1"`
	HighValue *int `hl7:"POS=2"`
}

// XAD - Extended Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/XAD
type XAD struct {
	StreetAddress SAD `hl7:"POS=1"`
	OtherDesignation string `hl7:"POS=2"`
	City string `hl7:"POS=3"`
	StateOrProvince string `hl7:"POS=4"`
	ZipOrPostalCode string `hl7:"POS=5"`
	Country string `hl7:"POS=6"`
	AddressType string `hl7:"POS=7"`
	OtherGeographicDesignation string `hl7:"POS=8"`
	CountyParishCode string `hl7:"POS=9"`
	CensusTract string `hl7:"POS=10"`
	AddressRepresentationCode string `hl7:"POS=11"`
	AddressValidityRange DR `hl7:"POS=12"`
	EffectiveDate time.Time `hl7:"POS=13"`
	ExpirationDate time.Time `hl7:"POS=14"`
	ExpirationReason CWE `hl7:"POS=15"`
	TemporaryIndicator string `hl7:"POS=16"`
	BadAddressIndicator string `hl7:"POS=17"`
	AddressUsage string `hl7:"POS=18"`
	Addressee string `hl7:"POS=19"`
	Comment string `hl7:"POS=20"`
	PreferenceOrder *int `hl7:"POS=21"`
	ProtectionCode CWE `hl7:"POS=22"`
	AddressIdentifier EI `hl7:"POS=23"`
}

// VH - Visiting Hours
// https://hl7-definition-staging.caristix.com/v2/HL7v2.6/DataTypes/VH
type VH struct {
	StartDayRange string `hl7:"POS=1"`
	EndDayRange string `hl7:"POS=2"`
	StartHourRange string `hl7:"POS=3"` // TODO - handle TM type parsing
	EndHourRange string `hl7:"POS=4"` // TODO - handle TM type parsing
}

