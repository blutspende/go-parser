package hl7v22

import "time"

// CM_RFR - Reference Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_RFR
type CM_RFR struct {
	ReferenceRange CM_RANGE `hl7:"POS=1"`
	Sex string `hl7:"POS=2"`
	AgeRange CM_RANGE `hl7:"POS=3"`
	GestationalAgeRange CM_RANGE `hl7:"POS=4"`
	Species string `hl7:"POS=5"`
	RaceSubspecies string `hl7:"POS=6"`
	TextCondition string `hl7:"POS=7"`
}

// CM_PLN - Practitioner ID Numbers
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PLN
type CM_PLN struct {
	IDNumber string `hl7:"POS=1"`
	TypeOfIDNumberID string `hl7:"POS=2"`
	StateOtherQualifyingInfo string `hl7:"POS=3"`
}

// CM_NDL - Observing Practitioner
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_NDL
type CM_NDL struct {
	InterpreterTechnician CN_PERSON `hl7:"POS=1"`
	StartDateTime time.Time `hl7:"POS=2"`
	EndDateTime time.Time `hl7:"POS=3"`
	Location CM_INTERNAL_LOCATION `hl7:"POS=4"`
}

// CM_EIP - Parent Order
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_EIP
type CM_EIP struct {
	ParentsPlacerOrderNumber string `hl7:"POS=1"`
	ParentsFillerOrderNumber string `hl7:"POS=2"`
}

// CM_RANGE - Range Of Values
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_RANGE
type CM_RANGE struct {
	LowValue string `hl7:"POS=1"`
	HighValue string `hl7:"POS=2"`
}

// MO - Money
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/MO
type MO struct {
	Quantity *int `hl7:"POS=1"`
	Denomination string `hl7:"POS=2"`
}

// CM_DLT - Delta Check
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_DLT
type CM_DLT struct {
	Range CM_RANGE `hl7:"POS=1"`
	NumericThreshold *int `hl7:"POS=2"`
	Change string `hl7:"POS=3"`
	LengthOfTimeDays *int `hl7:"POS=4"`
}

// CM_MSG - Message Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_MSG
type CM_MSG struct {
	MessageType string `hl7:"POS=1"`
	TriggerEvent string `hl7:"POS=2"`
}

// CQ - Composite Quantity With Units
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CQ
type CQ struct {
	Quantity *int `hl7:"POS=1"`
	Units string `hl7:"POS=2"`
}

// CM_CCD - Charge Time
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_CCD
type CM_CCD struct {
	WhenToCharge string `hl7:"POS=1"`
	DateTime time.Time `hl7:"POS=2"`
}

// FC - Financial Class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/FC
type FC struct {
	FinancialClass string `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// CM_BATCH_TOTAL - Cm For Batch Totals
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_BATCH_TOTAL
type CM_BATCH_TOTAL struct {
	BatchTotal1 *int `hl7:"POS=1"`
	BatchTotal2 *int `hl7:"POS=2"`
	 *int `hl7:"POS=3"`
}

// CM_LICENSE_NO - Cm For Driving License
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_LICENSE_NO
type CM_LICENSE_NO struct {
	LicenseNumber string `hl7:"POS=1"`
	IssuingStateProvinceCountry string `hl7:"POS=2"`
}

// CF - Coded Element with Formatted Values 
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CF
type CF struct {
	Identifier string `hl7:"POS=1"`
	FormattedText string `hl7:"POS=2"`
	NameOfCodingSystem string `hl7:"POS=3"`
	AlternateIdentifier string `hl7:"POS=4"`
	AlternateFormattedText string `hl7:"POS=5"`
	NameOfAlternateCodingSystem string `hl7:"POS=6"`
}

// CM_INTERNAL_LOCATION - Cm For Location Information In Hospital
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_INTERNAL_LOCATION
type CM_INTERNAL_LOCATION struct {
	NurseUnitStation string `hl7:"POS=1"`
	Room string `hl7:"POS=2"`
	Bed string `hl7:"POS=3"`
	FacilityID string `hl7:"POS=4"`
	BedStatus string `hl7:"POS=5"`
}

// CM_PLACER - Order Number Of The Client / The Contracting Authority
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PLACER
type CM_PLACER struct {
	UniquePlacerID string `hl7:"POS=1"`
	PlacerApplication string `hl7:"POS=2"`
}

// CM_SPD - Specialty
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_SPD
type CM_SPD struct {
	SpecialtyName string `hl7:"POS=1"`
	GoverningBoard string `hl7:"POS=2"`
	EligibleOrCertified string `hl7:"POS=3"`
	DateOfCertification time.Time `hl7:"POS=4;ATR=date"`
}

// CE - Coded Element
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CE
type CE struct {
	Identifier string `hl7:"POS=1"`
	Text string `hl7:"POS=2"`
	NameOfCodingSystem string `hl7:"POS=3"`
	AlternateIdentifier string `hl7:"POS=4"`
	AlternateText string `hl7:"POS=5"`
	NameOfAlternateCodingSystem string `hl7:"POS=6"`
}

// CM_POSITION - Seat Or Room / Table / Chair
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_POSITION
type CM_POSITION struct {
	Hall string `hl7:"POS=1"`
	Table string `hl7:"POS=2"`
	Chair string `hl7:"POS=3"`
}

// CM_LA1 - Location With Address Information
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_LA1
type CM_LA1 struct {
	DispenseDeliverToLocation CM_INTERNAL_LOCATION `hl7:"POS=1"`
	Location AD `hl7:"POS=2"`
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/XCN
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

// PN - Person Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/PN
type PN struct {
	FamilyName string `hl7:"POS=1"`
	GivenName string `hl7:"POS=2"`
	MiddleInitialOrName string `hl7:"POS=3"`
	SuffixEgJrOrIii string `hl7:"POS=4"`
	PrefixEgDr string `hl7:"POS=5"`
	DegreeEgMd string `hl7:"POS=6"`
}

// DLN - Driver's License Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/DLN
type DLN struct {
	DriversLicenseNumber string `hl7:"POS=1"`
	IssuingStateProvinceCountry string `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
}

// CM_DLD - Discharge Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_DLD
type CM_DLD struct {
	Code string `hl7:"POS=1"`
	Description string `hl7:"POS=2"`
}

// JCC - Job Code Class
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/JCC
type JCC struct {
	JobCode string `hl7:"POS=1"`
	JobClass string `hl7:"POS=2"`
}

// CM_GROUP_ID - Order Group Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_GROUP_ID
type CM_GROUP_ID struct {
	UniqueGroupID string `hl7:"POS=1"`
	PlacerApplicationID string `hl7:"POS=2"`
}

// CM_AUI - Authorization Information
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_AUI
type CM_AUI struct {
	AuthorizationNumber string `hl7:"POS=1"`
	Date time.Time `hl7:"POS=2;ATR=date"`
	Source string `hl7:"POS=3"`
}

// CM_OCD - Occurrence
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_OCD
type CM_OCD struct {
	OccurrenceCode string `hl7:"POS=1"`
	OccurrenceDate time.Time `hl7:"POS=2;ATR=date"`
}

// CM_PAT_ID_0192 - Patient IDWith Table 0192
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PAT_ID_0192
type CM_PAT_ID_0192 struct {
	PatientID string `hl7:"POS=1"`
	CheckDigit *int `hl7:"POS=2"`
	CheckDigitScheme string `hl7:"POS=3"`
	FacilityID string `hl7:"POS=4"`
	Type string `hl7:"POS=5"`
}

// CM_JOB_CODE - Job Title
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_JOB_CODE
type CM_JOB_CODE struct {
	JobCode string `hl7:"POS=1"`
	EmployeeClassification string `hl7:"POS=2"`
}

// CP - Composite Price
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CP
type CP struct {
	Price MO `hl7:"POS=1"`
	PriceType string `hl7:"POS=2"`
	FromValue *int `hl7:"POS=3"`
	ToValue *int `hl7:"POS=4"`
	RangeUnits CE `hl7:"POS=5"`
	RangeType string `hl7:"POS=6"`
}

// CM_DTN - Day Type And Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_DTN
type CM_DTN struct {
	DayType string `hl7:"POS=1"`
	NumberOfDays *int `hl7:"POS=2"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/XTN
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

// CM_FILLER - Number Of Processing Power Point
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_FILLER
type CM_FILLER struct {
	UniqueFillerID string `hl7:"POS=1"`
	FillerApplicationID string `hl7:"POS=2"`
}

// CX - Extended Composite ID With Check Digit
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CX
type CX struct {
	ID string `hl7:"POS=1"`
	CheckDigit string `hl7:"POS=2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=3"`
	AssigningAuthority HD `hl7:"POS=4"`
	IdentifierTypeCode string `hl7:"POS=5"`
	AssigningFacility HD `hl7:"POS=6"`
}

// CM_DIN - Activation Date
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_DIN
type CM_DIN struct {
	Date time.Time `hl7:"POS=1"`
	InstitutionName CE `hl7:"POS=2"`
}

// CM_OSD - Order Sequence
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_OSD
type CM_OSD struct {
	SequenceResultsFlag string `hl7:"POS=1"`
	PlacerOrderNumberEntityIdentifier string `hl7:"POS=2"`
	PlacerOrderNumberNamespaceID string `hl7:"POS=3"`
	FillerOrderNumberEntityIdentifier string `hl7:"POS=4"`
	FillerOrderNumberNamespaceID string `hl7:"POS=5"`
	SequenceConditionValue string `hl7:"POS=6"`
	MaximumNumberOfRepeats *int `hl7:"POS=7"`
}

// CM_VR - Value Qualifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_VR
type CM_VR struct {
	FirstDataCodeValue string `hl7:"POS=1"`
	LastDataCodeValue string `hl7:"POS=2"`
}

// AD - Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/AD
type AD struct {
	StreetAddress string `hl7:"POS=1"`
	OtherDesignation string `hl7:"POS=2"`
	City string `hl7:"POS=3"`
	StateOrProvince string `hl7:"POS=4"`
	ZipOrPostalCode string `hl7:"POS=5"`
	Country string `hl7:"POS=6"`
	Type string `hl7:"POS=7"`
	OtherGeographicDesignation string `hl7:"POS=8"`
}

// CK - Composite IDWith Check Digit
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CK
type CK struct {
	IDNumber *int `hl7:"POS=1"`
	CheckDigit *int `hl7:"POS=2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=3"`
	AssigningFacilityID string `hl7:"POS=4"`
}

// CM_OSP - Occurrence Span
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_OSP
type CM_OSP struct {
	OccurrenceSpanCode string `hl7:"POS=1"`
	OccurrenceSpanStartDate time.Time `hl7:"POS=2;ATR=date"`
	OccurrenceSpanStopDate time.Time `hl7:"POS=3;ATR=date"`
}

// PT - Processing Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/PT
type PT struct {
	ProcessingID string `hl7:"POS=1"`
	ProcessingMode string `hl7:"POS=2"`
}

// CM_PTA - Policy Type
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PTA
type CM_PTA struct {
	PolicyType string `hl7:"POS=1"`
	AmountClass string `hl7:"POS=2"`
	Amount *int `hl7:"POS=3"`
}

// CM_PIP - Privileges
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PIP
type CM_PIP struct {
	Privilege CE `hl7:"POS=1"`
	PrivilegeClass CE `hl7:"POS=2"`
	ExpirationDate time.Time `hl7:"POS=3;ATR=date"`
	ActivationDate time.Time `hl7:"POS=4;ATR=date"`
}

// CM_PRACTITIONER - Action Carried Out By
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PRACTITIONER
type CM_PRACTITIONER struct {
	ProcedurePractitionerID CN_PERSON `hl7:"POS=1"`
	ProcedurePractitionerType string `hl7:"POS=2"`
}

// CM_PEN - Penalty
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PEN
type CM_PEN struct {
	PenaltyID string `hl7:"POS=1"`
	PenaltyAmount *int `hl7:"POS=2"`
}

// TQ - Timing Quantity
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/TQ
type TQ struct {
	Quantity CQ `hl7:"POS=1"`
	Interval CM_RI `hl7:"POS=2"`
	Duration string `hl7:"POS=3"`
	StartDateTime time.Time `hl7:"POS=4"`
	EndDateTime time.Time `hl7:"POS=5"`
	Priority string `hl7:"POS=6"`
	Condition string `hl7:"POS=7"`
	Text string `hl7:"POS=8"`
	Conjunction string `hl7:"POS=9"`
	OrderSequencing CM_OSD `hl7:"POS=10"`
}

// CN_PERSON - Cn For Person
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CN_PERSON
type CN_PERSON struct {
	IDNumber string `hl7:"POS=1"`
	FamilyName string `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	MiddleInitialOrName string `hl7:"POS=4"`
	SuffixEgJrOrIii string `hl7:"POS=5"`
	PrefixEgDr string `hl7:"POS=6"`
	DegreeEgMd string `hl7:"POS=7"`
	SourceTableID string `hl7:"POS=8"`
}

// XON - Extended Composite Name And ID For Organizations
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/XON
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

// CM_PAT_ID - Patient ID
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PAT_ID
type CM_PAT_ID struct {
	PatientID string `hl7:"POS=1"`
	CheckDigit *int `hl7:"POS=2"`
	CheckDigitScheme string `hl7:"POS=3"`
	FacilityID string `hl7:"POS=4"`
	Type string `hl7:"POS=5"`
}

// XPN - Extended Person Name
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/XPN
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
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_RMC
type CM_RMC struct {
	RoomType string `hl7:"POS=1"`
	AmountType string `hl7:"POS=2"`
	CoverageAmount *int `hl7:"POS=3"`
}

// CM_PRL - Parent Result Link
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PRL
type CM_PRL struct {
	Obx3ObservationIdentifierOfParentResult string `hl7:"POS=1"`
	Obx4SubIDOfParentResult string `hl7:"POS=2"`
	Obx5ObservationResultsFromParent CE `hl7:"POS=3"`
}

// CM_RI - Interval
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_RI
type CM_RI struct {
	RepeatPattern string `hl7:"POS=1"`
	ExplicitTimeInterval string `hl7:"POS=2"`
}

// CM_SPS - Specimen Source
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_SPS
type CM_SPS struct {
	SpecimenSourceNameOrCode CE `hl7:"POS=1"`
	Additives string `hl7:"POS=2"`
	Freetext string `hl7:"POS=3"`
	BodySite CE `hl7:"POS=4"`
	SiteModifier CE `hl7:"POS=5"`
}

// TS - Time Stamp
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/TS
type TS struct {
	TimeOfAnEvent string `hl7:"POS=1"`
	DegreeOfPrecision string `hl7:"POS=2"`
}

// RP - Reference Pointer
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/RP
type RP struct {
	Pointer string `hl7:"POS=1"`
	ApplicationID string `hl7:"POS=2"`
	TypeOfData string `hl7:"POS=3"`
}

// CM_DDI - Daily Deductible
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_DDI
type CM_DDI struct {
	DelayDays string `hl7:"POS=1"`
	Amount *int `hl7:"POS=2"`
	NumberOfDays *int `hl7:"POS=3"`
}

// CM_ELD - Error
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_ELD
type CM_ELD struct {
	SegmentID string `hl7:"POS=1"`
	Sequence *int `hl7:"POS=2"`
	FieldPosition *int `hl7:"POS=3"`
	CodeIdentifyingError CE `hl7:"POS=4"`
}

// HD - Hierarchic Designator
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/HD
type HD struct {
	NamespaceID string `hl7:"POS=1"`
	UniversalID string `hl7:"POS=2"`
	UniversalIDType string `hl7:"POS=3"`
}

// COMP_ID_DIGIT - Composite IDW/chk Digit
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/COMP_ID_DIGIT
type COMP_ID_DIGIT struct {
	IDNumber *int `hl7:"POS=1"`
	CheckDigit *int `hl7:"POS=2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"POS=3"`
	AssigningFacilityID string `hl7:"POS=4"`
}

// CM_PCF - Pre-certification Required
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_PCF
type CM_PCF struct {
	PreCertificationPatientType string `hl7:"POS=1"`
	PreCertificationRequired string `hl7:"POS=2"`
	PreCertificationWindow time.Time `hl7:"POS=3"`
}

// CM_ABS_RANGE - Absolute Range
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_ABS_RANGE
type CM_ABS_RANGE struct {
	Range CM_RANGE `hl7:"POS=1"`
	NumericChange *int `hl7:"POS=2"`
	PercentPerChange *int `hl7:"POS=3"`
	Days *int `hl7:"POS=4"`
}

// EI - Entity Identifier
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"POS=1"`
	NamespaceID string `hl7:"POS=2"`
	UniversalID string `hl7:"POS=3"`
	UniversalIDType string `hl7:"POS=4"`
}

// PL - Person Location
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/PL
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

// CM_MOC - Charge To Practise
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_MOC
type CM_MOC struct {
	DollarAmount string `hl7:"POS=1"`
	ChargeCode string `hl7:"POS=2"`
}

// CM_UVC - Value Code And Amount
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_UVC
type CM_UVC struct {
	ValueCode string `hl7:"POS=1"`
	ValueAmount *int `hl7:"POS=2"`
}

// CN_PHYSICIAN - Cn For Physicians
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CN_PHYSICIAN
type CN_PHYSICIAN struct {
	PhysicianID string `hl7:"POS=1"`
	FamilyName string `hl7:"POS=2"`
	GivenName string `hl7:"POS=3"`
	MiddleInitialOrName string `hl7:"POS=4"`
	SuffixEgJrOrIii string `hl7:"POS=5"`
	PrefixEgDr string `hl7:"POS=6"`
	DegreeEgMd string `hl7:"POS=7"`
	SourceTableID string `hl7:"POS=8"`
}

// CM_FINANCE - Cm Of Finance
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/CM_FINANCE
type CM_FINANCE struct {
	FinancialClassID string `hl7:"POS=1"`
	EffectiveDate time.Time `hl7:"POS=2"`
}

// XAD - Extended Address
// https://hl7-definition-staging.caristix.com/v2/HL7v2.2/DataTypes/XAD
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

