package hl7v28

import "time"

// MO - Money
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/MO
type MO struct {
	Quantity *int `hl7:"POS=1"`
	Denomination string `hl7:"POS=2"`
}

// ERL - Error Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/ERL
type ERL struct {
	SegmentID string `hl7:"POS=1"`
	SegmentSequence *int `hl7:"POS=2"`
	FieldPosition *int `hl7:"POS=3"`
	FieldRepetition *int `hl7:"POS=4"`
	ComponentNumber *int `hl7:"POS=5"`
	SubComponentNumber *int `hl7:"POS=6"`
}

// CQ - Composite Quantity With Units
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CQ
type CQ struct {
	Quantity *int `hl7:"POS=1"`
	Units CWE `hl7:"POS=2"`
}

// NA - Numeric Array
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/NA
type NA struct {
	Value1 *int `hl7:"POS=1"`
	Value2 *int `hl7:"POS=2"`
	Value3 *int `hl7:"POS=3"`
	Value4 *int `hl7:"POS=4"`
}

// FC - Financial Class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/FC
type FC struct {
	FinancialClassCode CWE `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// SAD - Street Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/SAD
type SAD struct {
	StreetOrMailingAddress string `hl7:"POS=1"`
	StreetName string `hl7:"POS=2"`
	DwellingNumber string `hl7:"POS=3"`
}

// DIN - Date And Institution Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/DIN
type DIN struct {
	Date time.Time `hl7:"POS=1"`
	InstitutionName CWE `hl7:"POS=2"`
}

// CF - Coded Element With Formatted Values
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CF
type CF struct {
	Identifier string `hl7:"POS=1"`
	FormattedText string `hl7:"POS=2"`
	NameOfCodingSystem string `hl7:"POS=3"`
	AlternateIdentifier string `hl7:"POS=4"`
	AlternateFormattedText string `hl7:"POS=5"`
	NameOfAlternateCodingSystem string `hl7:"POS=6"`
	CodingSystemVersionID string `hl7:"POS=7"`
	AlternateCodingSystemVersionID string `hl7:"POS=8"`
	OriginalText string `hl7:"POS=9"`
	SecondAlternateIdentifier string `hl7:"POS=10"`
	SecondAlternateFormattedText string `hl7:"POS=11"`
	NameOfSecondAlternateCodingSystem string `hl7:"POS=12"`
	SecondAlternateCodingSystemVersionID string `hl7:"POS=13"`
	CodingSystemOid string `hl7:"POS=14"`
	ValueSetOid string `hl7:"POS=15"`
	ValueSetVersionID time.Time `hl7:"POS=16"`
	AlternateCodingSystemOid string `hl7:"POS=17"`
	AlternateValueSetOid string `hl7:"POS=18"`
	AlternateValueSetVersionID time.Time `hl7:"POS=19"`
	SecondAlternateCodingSystemOid string `hl7:"POS=20"`
	SecondAlternateValueSetOid string `hl7:"POS=21"`
	SecondAlternateValueSetVersionID time.Time `hl7:"POS=22"`
}

// EIP - Entity Identifier Pair
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/EIP
type EIP struct {
	PlacerAssignedIdentifier EI `hl7:"POS=1"`
	FillerAssignedIdentifier EI `hl7:"POS=2"`
}

// SCV - Scheduling Class Value Pair
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/SCV
type SCV struct {
	ParameterClass CWE `hl7:"POS=1"`
	ParameterValue string `hl7:"POS=2"`
}

// CD - Channel Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CD
type CD struct {
	ChannelIdentifier WVI `hl7:"POS=1"`
	WaveformSource WVS `hl7:"POS=2"`
	ChannelSensitivityAndUnits CSU `hl7:"POS=3"`
	ChannelCalibrationParameters CCP `hl7:"POS=4"`
	ChannelSamplingFrequency *int `hl7:"POS=5"`
	MinimumAndMaximumDataValues NR `hl7:"POS=6"`
}

// PIP - Practitioner Institutional Privileges
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/PIP
type PIP struct {
	Privilege CWE `hl7:"POS=1"`
	PrivilegeClass CWE `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
	ActivationDate time.Time `hl7:"POS=4;ATR=date"`
	Facility EI `hl7:"POS=5"`
}

// RI - Repeat Interval
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/RI
type RI struct {
	RepeatPattern CWE `hl7:"POS=1"`
	ExplicitTimeInterval string `hl7:"POS=2"`
}

// XCN - Extended Composite ID Number And Name For Persons
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/XCN
type XCN struct {
	PersonIdentifier string `hl7:"POS=1"`
	FamilyName FN `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=4"`
	SuffixEgJrOrIii string `hl7:"POS=5"`
	PrefixEgDr string `hl7:"POS=6"`
	SourceTable CWE `hl7:"POS=8"`
	AssigningAuthority HD `hl7:"POS=9"`
	NameTypeCode string `hl7:"POS=10"`
	IdentifierCheckDigit string `hl7:"POS=11"`
	CheckDigitScheme string `hl7:"POS=12"`
	IdentifierTypeCode string `hl7:"POS=13"`
	AssigningFacility HD `hl7:"POS=14"`
	NameRepresentationCode string `hl7:"POS=15"`
	NameContext CWE `hl7:"POS=16"`
	NameAssemblyOrder string `hl7:"POS=18"`
	EffectiveDate time.Time `hl7:"POS=19"`
	ExpirationDate time.Time `hl7:"POS=20"`
	ProfessionalSuffix string `hl7:"POS=21"`
	AssigningJurisdiction CWE `hl7:"POS=22"`
	AssigningAgencyOrDepartment CWE `hl7:"POS=23"`
	SecurityCheck string `hl7:"POS=24"`
	SecurityCheckScheme string `hl7:"POS=25"`
}

// QSC - Query Selection Criteria
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/QSC
type QSC struct {
	SegmentFieldName string `hl7:"POS=1"`
	RelationalOperator string `hl7:"POS=2"`
	Value string `hl7:"POS=3"`
	RelationalConjunction string `hl7:"POS=4"`
}

// ED - Encapsulated Data
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/ED
type ED struct {
	SourceApplication HD `hl7:"POS=1"`
	TypeOfData string `hl7:"POS=2"`
	DataSubtype string `hl7:"POS=3"`
	Encoding string `hl7:"POS=4"`
	Data string `hl7:"POS=5"`
}

// DLD - Discharge To Location And Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/DLD
type DLD struct {
	DischargeToLocation CWE `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// PLN - Practitioner License Or Other ID Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/PLN
type PLN struct {
	IDNumber string `hl7:"POS=1"`
	TypeOfIDNumber CWE `hl7:"POS=2"`
	StateOtherQualifyingInformation string `hl7:"POS=3"`
	ExpirationDate time.Time `hl7:"POS=4;ATR=date"`
}

// CNN - Composite ID Number And Name Simplified
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CNN
type CNN struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName string `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=4"`
	SuffixEgJrOrIii string `hl7:"POS=5"`
	PrefixEgDr string `hl7:"POS=6"`
	DegreeEgMd string `hl7:"POS=7"`
	SourceTable string `hl7:"POS=8"`
	AssigningAuthorityNamespaceID string `hl7:"POS=9"`
	AssigningAuthorityUniversalID string `hl7:"POS=10"`
	AssigningAuthorityUniversalIdtype string `hl7:"POS=11"`
}

// OCD - Occurrence Code And Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/OCD
type OCD struct {
	OccurrenceCode CNE `hl7:"POS=1"`
	OccurrenceDate time.Time `hl7:"POS=2;ATR=date"`
}

// DLN - Driver's License Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/DLN
type DLN struct {
	LicenseNumber string `hl7:"POS=1"`
	IssuingStateProvinceCountry CWE `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
}

// CSU - Channel Sensitivity And Units
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CSU
type CSU struct {
	ChannelSensitivity *int `hl7:"POS=1"`
	UnitOfMeasureIdentifier string `hl7:"POS=2"`
	UnitOfMeasureDescription string `hl7:"POS=3"`
	UnitOfMeasureCodingSystem string `hl7:"POS=4"`
	AlternateUnitOfMeasureIdentifier string `hl7:"POS=5"`
	AlternateUnitOfMeasureDescription string `hl7:"POS=6"`
	AlternateUnitOfMeasureCodingSystem string `hl7:"POS=7"`
	UnitOfMeasureCodingSystemVersionID string `hl7:"POS=8"`
	AlternateUnitOfMeasureCodingSystemVersionID string `hl7:"POS=9"`
	OriginalText string `hl7:"POS=10"`
	SecondAlternateUnitOfMeasureIdentifier string `hl7:"POS=11"`
	SecondAlternateUnitOfMeasureText string `hl7:"POS=12"`
	NameOfSecondAlternateUnitOfMeasureCodingSy string `hl7:"POS=13"`
	SecondAlternateUnitOfMeasureCodingSystemVer string `hl7:"POS=14"`
	UnitOfMeasureCodingSystemOid string `hl7:"POS=15"`
	UnitOfMeasureValueSetOid string `hl7:"POS=16"`
	UnitOfMeasureValueSetVersionID time.Time `hl7:"POS=17"`
	AlternateUnitOfMeasureCodingSystemOid1 string `hl7:"POS=18"`
	AlternateUnitOfMeasureValueSetOid1 string `hl7:"POS=19"`
	AlternateUnitOfMeasureValueSetVersionID1 time.Time `hl7:"POS=20"`
	AlternateUnitOfMeasureCodingSystemOid2 string `hl7:"POS=21"`
	AlternateUnitOfMeasureValueSetOid2 string `hl7:"POS=22"`
	AlternateUnitOfMeasureValueSetVersionID2 string `hl7:"POS=23"`
}

// JCC - Job Code/class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/JCC
type JCC struct {
	JobCode CWE `hl7:"POS=1"`
	JobClass CWE `hl7:"POS=2"`
	JobDescriptionText string `hl7:"POS=3"`
}

// DTN - Day Type And Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/DTN
type DTN struct {
	DayType CWE `hl7:"POS=1"`
	NumberOfDays *int `hl7:"POS=2"`
}

// WVS - Waveform Source
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/WVS
type WVS struct {
	SourceOneName string `hl7:"POS=1"`
	SourceTwoName string `hl7:"POS=2"`
}

// MSG - Message Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/MSG
type MSG struct {
	MessageCode string `hl7:"POS=1"`
	TriggerEvent string `hl7:"POS=2"`
	MessageStructure string `hl7:"POS=3"`
}

// DR - Date/time Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/DR
type DR struct {
	RangeStartDateTime time.Time `hl7:"POS=1"`
	RangeEndDateTime time.Time `hl7:"POS=2"`
}

// CNE - Coded With No Exceptions
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CNE
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
	SecondAlternateIdentifier string `hl7:"POS=10"`
	SecondAlternateText string `hl7:"POS=11"`
	NameOfSecondAlternateCodingSystem string `hl7:"POS=12"`
	SecondAlternateCodingSystemVersionID string `hl7:"POS=13"`
	CodingSystemOid string `hl7:"POS=14"`
	ValueSetOid string `hl7:"POS=15"`
	ValueSetVersionID time.Time `hl7:"POS=16"`
	AlternateCodingSystemOid string `hl7:"POS=17"`
	AlternateValueSetOid string `hl7:"POS=18"`
	AlternateValueSetVersionID time.Time `hl7:"POS=19"`
	SecondAlternateCodingSystemOid string `hl7:"POS=20"`
	SecondAlternateValueSetOid string `hl7:"POS=21"`
	SecondAlternateValueSetVersionID time.Time `hl7:"POS=22"`
}

// CP - Composite Price
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CP
type CP struct {
	Price MO `hl7:"POS=1"`
	PriceType string `hl7:"POS=2"`
	FromValue *int `hl7:"POS=3"`
	ToValue *int `hl7:"POS=4"`
	RangeUnits CWE `hl7:"POS=5"`
	RangeType string `hl7:"POS=6"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/XTN
type XTN struct {
	TelecommunicationUseCode string `hl7:"POS=2"`
	TelecommunicationEquipmentType string `hl7:"POS=3"`
	CommunicationAddress string `hl7:"POS=4"`
	CountryCode string `hl7:"POS=5"`
	AreaCityCode string `hl7:"POS=6"`
	LocalNumber string `hl7:"POS=7"`
	Extension string `hl7:"POS=8"`
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

// MOP - Money Or Percentage
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/MOP
type MOP struct {
	MoneyOrPercentageIndicator string `hl7:"POS=1"`
	MoneyOrPercentageQuantity *int `hl7:"POS=2"`
	MonetaryDenomination string `hl7:"POS=3"`
}

// CX - Extended Composite IDWith Check Digit
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CX
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
	SecurityCheck string `hl7:"POS=11"`
	SecurityCheckScheme string `hl7:"POS=12"`
}

// ICD - Insurance Certification Definition
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/ICD
type ICD struct {
	CertificationPatientType CWE `hl7:"POS=1"`
	CertificationRequired string `hl7:"POS=2"`
	DateTimeCertificationRequired time.Time `hl7:"POS=3"`
}

// CCP - Channel Calibration Parameters
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CCP
type CCP struct {
	ChannelCalibrationSensitivityCorrectionFactor *int `hl7:"POS=1"`
	ChannelCalibrationBaseline *int `hl7:"POS=2"`
	ChannelCalibrationTimeSkew *int `hl7:"POS=3"`
}

// NDL - Name With Date And Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/NDL
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
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/RCD
type RCD struct {
	SegmentFieldName string `hl7:"POS=1"`
	Hl7DataType string `hl7:"POS=2"`
	MaximumColumnWidth *int `hl7:"POS=3"`
}

// AD - Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/AD
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
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/AUI
type AUI struct {
	AuthorizationNumber string `hl7:"POS=1"`
	Date time.Time `hl7:"POS=2;ATR=date"`
	Source string `hl7:"POS=3"`
}

// MOC - Money And Code
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/MOC
type MOC struct {
	MonetaryAmount MO `hl7:"POS=1"`
	ChargeCode CWE `hl7:"POS=2"`
}

// PT - Processing Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/PT
type PT struct {
	ProcessingID string `hl7:"POS=1"`
	ProcessingMode string `hl7:"POS=2"`
}

// DDI - Daily Deductible Information
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/DDI
type DDI struct {
	DelayDays *int `hl7:"POS=1"`
	MonetaryAmount MO `hl7:"POS=2"`
	NumberOfDays *int `hl7:"POS=3"`
}

// VR - Value Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/VR
type VR struct {
	FirstDataCodeValue string `hl7:"POS=1"`
	LastDataCodeValue string `hl7:"POS=2"`
}

// FN - Family Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/FN
type FN struct {
	Surname string `hl7:"POS=1"`
	OwnSurnamePrefix string `hl7:"POS=2"`
	OwnSurname string `hl7:"POS=3"`
	SurnamePrefixFromPartnerSpouse string `hl7:"POS=4"`
	SurnameFromPartnerSpouse string `hl7:"POS=5"`
}

// RMC - Room Coverage
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/RMC
type RMC struct {
	RoomType CWE `hl7:"POS=1"`
	AmountType CWE `hl7:"POS=2"`
	MoneyOrPercentage MOP `hl7:"POS=4"`
}

// PPN - Performing Person Time Stamp
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/PPN
type PPN struct {
	PersonIdentifier string `hl7:"POS=1"`
	FamilyName FN `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=4"`
	SuffixEgJrOrIii string `hl7:"POS=5"`
	PrefixEgDr string `hl7:"POS=6"`
	SourceTable CWE `hl7:"POS=8"`
	AssigningAuthority HD `hl7:"POS=9"`
	NameTypeCode string `hl7:"POS=10"`
	IdentifierCheckDigit string `hl7:"POS=11"`
	CheckDigitScheme string `hl7:"POS=12"`
	IdentifierTypeCode string `hl7:"POS=13"`
	AssigningFacility HD `hl7:"POS=14"`
	DateTimeActionPerformed time.Time `hl7:"POS=15"`
	NameRepresentationCode string `hl7:"POS=16"`
	NameContext CWE `hl7:"POS=17"`
	NameAssemblyOrder string `hl7:"POS=19"`
	EffectiveDate time.Time `hl7:"POS=20"`
	ExpirationDate time.Time `hl7:"POS=21"`
	ProfessionalSuffix string `hl7:"POS=22"`
	AssigningJurisdiction CWE `hl7:"POS=23"`
	AssigningAgencyOrDepartment CWE `hl7:"POS=24"`
	SecurityCheck string `hl7:"POS=25"`
	SecurityCheckScheme string `hl7:"POS=26"`
}

// XON - Extended Composite Name And Identification Number For Organizations
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/XON
type XON struct {
	OrganizationName string `hl7:"POS=1"`
	OrganizationNameTypeCode CWE `hl7:"POS=2"`
	AssigningAuthority HD `hl7:"POS=6"`
	IdentifierTypeCode string `hl7:"POS=7"`
	AssigningFacility HD `hl7:"POS=8"`
	NameRepresentationCode string `hl7:"POS=9"`
	OrganizationIdentifier string `hl7:"POS=10"`
}

// RPT - Repeat Pattern
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/RPT
type RPT struct {
	RepeatPatternCode CWE `hl7:"POS=1"`
	CalendarAlignment string `hl7:"POS=2"`
	PhaseRangeBeginValue *int `hl7:"POS=3"`
	PhaseRangeEndValue *int `hl7:"POS=4"`
	PeriodQuantity *int `hl7:"POS=5"`
	PeriodUnits CWE `hl7:"POS=6"`
	InstitutionSpecifiedTime string `hl7:"POS=7"`
	Event string `hl7:"POS=8"`
	EventOffsetQuantity *int `hl7:"POS=9"`
	EventOffsetUnits CWE `hl7:"POS=10"`
	GeneralTimingSpecification string `hl7:"POS=11"`
}

// QIP - Query Input Parameter List
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/QIP
type QIP struct {
	SegmentFieldName string `hl7:"POS=1"`
	Values string `hl7:"POS=2"`
}

// PRL - Parent Result Link
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/PRL
type PRL struct {
	ParentObservationIdentifier CWE `hl7:"POS=1"`
	ParentObservationSubIdentifier string `hl7:"POS=2"`
	ParentObservationValueDescriptor string `hl7:"POS=3"`
}

// DLT - Delta
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/DLT
type DLT struct {
	NormalRange NR `hl7:"POS=1"`
	NumericThreshold *int `hl7:"POS=2"`
	ChangeComputation string `hl7:"POS=3"`
	DaysRetained *int `hl7:"POS=4"`
}

// PTA - Policy Type And Amount
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/PTA
type PTA struct {
	PolicyType CWE `hl7:"POS=1"`
	AmountClass CWE `hl7:"POS=2"`
	MoneyOrPercentage MOP `hl7:"POS=4"`
}

// XPN - Extended Person Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/XPN
type XPN struct {
	FamilyName FN `hl7:"POS=1"`
	GivenName string `hl7:"POS=2"`
	SecondAndFurtherGivenNamesOrInitialsThereof string `hl7:"POS=3"`
	SuffixEgJrOrIii string `hl7:"POS=4"`
	PrefixEgDr string `hl7:"POS=5"`
	NameTypeCode string `hl7:"POS=7"`
	NameRepresentationCode string `hl7:"POS=8"`
	NameContext CWE `hl7:"POS=9"`
	NameAssemblyOrder string `hl7:"POS=11"`
	EffectiveDate time.Time `hl7:"POS=12"`
	ExpirationDate time.Time `hl7:"POS=13"`
	ProfessionalSuffix string `hl7:"POS=14"`
	CalledBy string `hl7:"POS=15"`
}

// SPD - Specialty Description
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/SPD
type SPD struct {
	SpecialtyName string `hl7:"POS=1"`
	GoverningBoard string `hl7:"POS=2"`
	EligibleOrCertified string `hl7:"POS=3"`
	DateOfCertification time.Time `hl7:"POS=4;ATR=date"`
}

// UVC - Ub Value Code And Amount
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/UVC
type UVC struct {
	ValueCode CWE `hl7:"POS=1"`
	ValueAmount MO `hl7:"POS=2"`
	NonMonetaryValueAmountQuantity *int `hl7:"POS=3"`
	NonMonetaryValueAmountUnits CWE `hl7:"POS=4"`
}

// CWE - Coded With Exceptions
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CWE
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
	SecondAlternateIdentifier string `hl7:"POS=10"`
	SecondAlternateText string `hl7:"POS=11"`
	NameOfSecondAlternateCodingSystem string `hl7:"POS=12"`
	SecondAlternateCodingSystemVersionID string `hl7:"POS=13"`
	CodingSystemOid string `hl7:"POS=14"`
	ValueSetOid string `hl7:"POS=15"`
	ValueSetVersionID time.Time `hl7:"POS=16"`
	AlternateCodingSystemOid string `hl7:"POS=17"`
	AlternateValueSetOid string `hl7:"POS=18"`
	AlternateValueSetVersionID time.Time `hl7:"POS=19"`
	SecondAlternateCodingSystemOid string `hl7:"POS=20"`
	SecondAlternateValueSetOid string `hl7:"POS=21"`
	SecondAlternateValueSetVersionID time.Time `hl7:"POS=22"`
}

// SRT - Sort Order
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/SRT
type SRT struct {
	SortByField string `hl7:"POS=1"`
	Sequencing string `hl7:"POS=2"`
}

// RP - Reference Pointer
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/RP
type RP struct {
	Pointer string `hl7:"POS=1"`
	ApplicationID HD `hl7:"POS=2"`
	TypeOfData string `hl7:"POS=3"`
	Subtype string `hl7:"POS=4"`
}

// RFR - Reference Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/RFR
type RFR struct {
	NumericRange NR `hl7:"POS=1"`
	AdministrativeSex CWE `hl7:"POS=2"`
	AgeRange NR `hl7:"POS=3"`
	GestationalAgeRange NR `hl7:"POS=4"`
	Species string `hl7:"POS=5"`
	RaceSubspecies string `hl7:"POS=6"`
	Conditions string `hl7:"POS=7"`
}

// OSP - Occurrence Span Code And Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/OSP
type OSP struct {
	OccurrenceSpanCode CNE `hl7:"POS=1"`
	OccurrenceSpanStartDate time.Time `hl7:"POS=2;ATR=date"`
	OccurrenceSpanStopDate time.Time `hl7:"POS=3;ATR=date"`
}

// VID - Version Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/VID
type VID struct {
	VersionID string `hl7:"POS=1"`
	InternationalizationCode CWE `hl7:"POS=2"`
	InternationalVersionID CWE `hl7:"POS=3"`
}

// HD - Hierarchic Designator
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/HD
type HD struct {
	NamespaceID string `hl7:"POS=1"`
	UniversalID string `hl7:"POS=2"`
	UniversalIdtype string `hl7:"POS=3"`
}

// EI - Entity Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"POS=1"`
	NamespaceID string `hl7:"POS=2"`
	UniversalID string `hl7:"POS=3"`
	UniversalIdtype string `hl7:"POS=4"`
}

// PL - Person Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/PL
type PL struct {
	PointOfCare HD `hl7:"POS=1"`
	Room HD `hl7:"POS=2"`
	Bed HD `hl7:"POS=3"`
	Facility HD `hl7:"POS=4"`
	LocationStatus string `hl7:"POS=5"`
	PersonLocationType string `hl7:"POS=6"`
	Building HD `hl7:"POS=7"`
	Floor HD `hl7:"POS=8"`
	LocationDescription string `hl7:"POS=9"`
	ComprehensiveLocationIdentifier EI `hl7:"POS=10"`
	AssigningAuthorityForLocation HD `hl7:"POS=11"`
}

// MA - Multiplexed Array
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/MA
type MA struct {
	SampleYFromChannel1 *int `hl7:"POS=1"`
	SampleYFromChannel2 *int `hl7:"POS=2"`
	SampleYFromChannel3 *int `hl7:"POS=3"`
	SampleYFromChannel4 *int `hl7:"POS=4"`
}

// SN - Structured Numeric
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/SN
type SN struct {
	Comparator string `hl7:"POS=1"`
	Num1 *int `hl7:"POS=2"`
	SeparatorSuffix string `hl7:"POS=3"`
	Num2 *int `hl7:"POS=4"`
}

// CCD - Charge Code And Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/CCD
type CCD struct {
	InvocationEvent string `hl7:"POS=1"`
	DateTime time.Time `hl7:"POS=2"`
}

// WVI - Channel Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/WVI
type WVI struct {
	ChannelNumber *int `hl7:"POS=1"`
	ChannelName string `hl7:"POS=2"`
}

// NR - Numeric Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/NR
type NR struct {
	LowValue *int `hl7:"POS=1"`
	HighValue *int `hl7:"POS=2"`
}

// XAD - Extended Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/XAD
type XAD struct {
	StreetAddress SAD `hl7:"POS=1"`
	OtherDesignation string `hl7:"POS=2"`
	City string `hl7:"POS=3"`
	StateOrProvince string `hl7:"POS=4"`
	ZipOrPostalCode string `hl7:"POS=5"`
	Country string `hl7:"POS=6"`
	AddressType string `hl7:"POS=7"`
	OtherGeographicDesignation string `hl7:"POS=8"`
	CountyParishCode CWE `hl7:"POS=9"`
	CensusTract CWE `hl7:"POS=10"`
	AddressRepresentationCode string `hl7:"POS=11"`
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
// https://hl7-definition-staging.caristix.com/v2/HL7v2.8/DataTypes/VH
type VH struct {
	StartDayRange string `hl7:"POS=1"`
	EndDayRange string `hl7:"POS=2"`
	StartHourRange string `hl7:"POS=3"` // TODO - handle TM type parsing
	EndHourRange string `hl7:"POS=4"` // TODO - handle TM type parsing
}

