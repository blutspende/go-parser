package hl7v28

import "time"

// OM1 - General Segment
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/OM1
type OM1 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	ProducersServiceTestObservationID CWE `hl7:"POS=3"`
	PermittedDataTypes []string `hl7:"POS=4"`
	SpecimenRequired string `hl7:"POS=5"`
	ProducerID CWE `hl7:"POS=6"`
	ObservationDescription string `hl7:"POS=7"`
	OtherServiceTestObservationIdsForTheObservation CWE `hl7:"POS=8"`
	OtherNames1 []string `hl7:"POS=9"`
	PreferredReportNameForTheObservation string `hl7:"POS=10"`
	PreferredShortNameOrMnemonicForTheObservation string `hl7:"POS=11"`
	PreferredLongNameForTheObservation string `hl7:"POS=12"`
	Orderability string `hl7:"POS=13"`
	IdentityOfInstrumentUsedToPerformThisStudy []CWE `hl7:"POS=14"`
	CodedRepresentationOfMethod []CWE `hl7:"POS=15"`
	PortableDeviceIndicator string `hl7:"POS=16"`
	ObservationProducingDepartmentSection []CWE `hl7:"POS=17"`
	TelephoneNumberOfSection XTN `hl7:"POS=18"`
	NatureOfServiceTestObservation CWE `hl7:"POS=19"`
	ReportSubheader CWE `hl7:"POS=20"`
	ReportDisplayOrder string `hl7:"POS=21"`
	DateTimeStampForAnyChangeInDefinitionForTheObservation time.Time `hl7:"POS=22"`
	EffectiveDateTimeOfChange time.Time `hl7:"POS=23"`
	TypicalTurnAroundTime *int `hl7:"POS=24"`
	ProcessingTime *int `hl7:"POS=25"`
	ProcessingPriority []string `hl7:"POS=26"`
	ReportingPriority string `hl7:"POS=27"`
	OutsideSitesWhereObservationMayBePerformed []CWE `hl7:"POS=28"`
	AddressOfOutsideSites []XAD `hl7:"POS=29"`
	PhoneNumberOfOutsideSite XTN `hl7:"POS=30"`
	ConfidentialityCode CWE `hl7:"POS=31"`
	ObservationsRequiredToInterpretThisObservation CWE `hl7:"POS=32"`
	InterpretationOfObservations string `hl7:"POS=33"`
	ContraindicationsToObservations CWE `hl7:"POS=34"`
	ReflexTestsObservations []CWE `hl7:"POS=35"`
	RulesThatTriggerReflexTesting string `hl7:"POS=36"`
	FixedCannedMessage CWE `hl7:"POS=37"`
	PatientPreparation string `hl7:"POS=38"`
	ProcedureMedication CWE `hl7:"POS=39"`
	FactorsThatMayAffectTheObservation string `hl7:"POS=40"`
	ServiceTestObservationPerformanceSchedule []string `hl7:"POS=41"`
	DescriptionOfTestMethods string `hl7:"POS=42"`
	KindOfQuantityObserved CWE `hl7:"POS=43"`
	PointVersusInterval CWE `hl7:"POS=44"`
	ChallengeInformation string `hl7:"POS=45"`
	RelationshipModifier CWE `hl7:"POS=46"`
	TargetAnatomicSiteOfTest CWE `hl7:"POS=47"`
	ModalityOfImagingMeasurement CWE `hl7:"POS=48"`
	ExclusiveTest string `hl7:"POS=49"`
	DiagnosticServSectID string `hl7:"POS=50"`
	TaxonomicClassificationCode CWE `hl7:"POS=51"`
	OtherNames2 []string `hl7:"POS=52"`
}

