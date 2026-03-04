package hl7v25

import "time"

// OM1 - General Segment
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/OM1
type OM1 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	ProducersServiceTestObservationID CE `hl7:"POS=3"`
	PermittedDataTypes []string `hl7:"POS=4"`
	SpecimenRequired string `hl7:"POS=5"`
	ProducerID CE `hl7:"POS=6"`
	ObservationDescription string `hl7:"POS=7"`
	OtherServiceTestObservationIdsForTheObservation CE `hl7:"POS=8"`
	OtherNames []string `hl7:"POS=9"`
	PreferredReportNameForTheObservation string `hl7:"POS=10"`
	PreferredShortNameOrMnemonicForObservation string `hl7:"POS=11"`
	PreferredLongNameForTheObservation string `hl7:"POS=12"`
	Orderability string `hl7:"POS=13"`
	IdentityOfInstrumentUsedToPerformThisStudy []CE `hl7:"POS=14"`
	CodedRepresentationOfMethod []CE `hl7:"POS=15"`
	PortableDeviceIndicator string `hl7:"POS=16"`
	ObservationProducingDepartmentSection []CE `hl7:"POS=17"`
	TelephoneNumberOfSection XTN `hl7:"POS=18"`
	NatureOfServiceTestObservation string `hl7:"POS=19"`
	ReportSubheader CE `hl7:"POS=20"`
	ReportDisplayOrder string `hl7:"POS=21"`
	DateTimeStampForAnyChangeInDefinitionForTheObservation time.Time `hl7:"POS=22"`
	EffectiveDateTimeOfChange time.Time `hl7:"POS=23"`
	TypicalTurnAroundTime *int `hl7:"POS=24"`
	ProcessingTime *int `hl7:"POS=25"`
	ProcessingPriority []string `hl7:"POS=26"`
	ReportingPriority string `hl7:"POS=27"`
	OutsideSitesWhereObservationMayBePerformed []CE `hl7:"POS=28"`
	AddressOfOutsideSites []XAD `hl7:"POS=29"`
	PhoneNumberOfOutsideSite XTN `hl7:"POS=30"`
	ConfidentialityCode CWE `hl7:"POS=31"`
	ObservationsRequiredToInterpretTheObservation CE `hl7:"POS=32"`
	InterpretationOfObservations string `hl7:"POS=33"`
	ContraindicationsToObservations CE `hl7:"POS=34"`
	ReflexTestsObservations []CE `hl7:"POS=35"`
	RulesThatTriggerReflexTesting string `hl7:"POS=36"`
	FixedCannedMessage CE `hl7:"POS=37"`
	PatientPreparation string `hl7:"POS=38"`
	ProcedureMedication CE `hl7:"POS=39"`
	FactorsThatMayAffectTheObservation string `hl7:"POS=40"`
	ServiceTestObservationPerformanceSchedule []string `hl7:"POS=41"`
	DescriptionOfTestMethods string `hl7:"POS=42"`
	KindOfQuantityObserved CE `hl7:"POS=43"`
	PointVersusInterval CE `hl7:"POS=44"`
	ChallengeInformation string `hl7:"POS=45"`
	RelationshipModifier CE `hl7:"POS=46"`
	TargetAnatomicSiteOfTest CE `hl7:"POS=47"`
	ModalityOfImagingMeasurement CE `hl7:"POS=48"`
}

