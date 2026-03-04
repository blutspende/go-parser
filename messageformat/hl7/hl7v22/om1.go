package hl7v22

import "time"

// OM1 - General - Fields That Apply To Most Observations
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/OM1
type OM1 struct {
	SegmentTypeID string `hl7:"POS=2"`
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=3"`
	ProducersTestObservationID CE `hl7:"POS=4"`
	PermittedDataTypes []string `hl7:"POS=5"`
	SpecimenRequired string `hl7:"POS=6"`
	ProducerID CE `hl7:"POS=7"`
	ObservationDescription string `hl7:"POS=8"`
	OtherTestObservationIdsForTheObservation CE `hl7:"POS=9"`
	OtherNames []string `hl7:"POS=10"`
	PreferredReportNameForTheObservation string `hl7:"POS=11"`
	PreferredShortNameOrMnemonicForObservation string `hl7:"POS=12"`
	PreferredLongNameForTheObservation string `hl7:"POS=13"`
	Orderability string `hl7:"POS=14"`
	IdentityOfInstrumentUsedToPerformThisStudy []CE `hl7:"POS=15"`
	CodedRepresentationOfMethod []CE `hl7:"POS=16"`
	Portable string `hl7:"POS=17"`
	ObservationProducingDepartmentSection []string `hl7:"POS=18"`
	TelephoneNumberOfSection string `hl7:"POS=19"`
	NatureOfTestObservation string `hl7:"POS=20"`
	ReportSubheader CE `hl7:"POS=21"`
	ReportDisplayOrder string `hl7:"POS=22"`
	DateTimeStampForAnyChangeInDefinitionForObs time.Time `hl7:"POS=23"`
	EffectiveDateTimeOfChange time.Time `hl7:"POS=24"`
	TypicalTurnAroundTime *int `hl7:"POS=25"`
	ProcessingTime *int `hl7:"POS=26"`
	ProcessingPriority []string `hl7:"POS=27"`
	ReportingPriority string `hl7:"POS=28"`
	OutsideSitesWhereObservationMayBePerformed []CE `hl7:"POS=29"`
	AddressOfOutsideSites AD `hl7:"POS=30"`
	PhoneNumberOfOutsideSite string `hl7:"POS=31"`
	ConfidentialityCode string `hl7:"POS=32"`
	ObservationsRequiredToInterpretTheObservation CE `hl7:"POS=33"`
	InterpretationOfObservations string `hl7:"POS=34"`
	ContraindicationsToObservations CE `hl7:"POS=35"`
	ReflexTestsObservations []CE `hl7:"POS=36"`
	RulesThatTriggerReflexTesting string `hl7:"POS=37"`
	FixedCannedMessage CE `hl7:"POS=38"`
	PatientPreparation string `hl7:"POS=39"`
	ProcedureMedication CE `hl7:"POS=40"`
	FactorsThatMayAffectTheObservation string `hl7:"POS=41"`
	TestObservationPerformanceSchedule []string `hl7:"POS=42"`
	DescriptionOfTestMethods string `hl7:"POS=43"`
}

