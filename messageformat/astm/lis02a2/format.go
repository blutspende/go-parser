package lis02a2

import "time"

// Standard LIS02-A2 record and message format declarations to use for (un)marshaling
// These structures can be used directly or as a base for custom implementations
// Basis for the protocol definition: https://samson-rus.com/wp-content/files/LIS2-A2.pdf
// The commended section numbers below (eg: 5.6.1) refer to the document's sections above
// Page 30-33 provides information about the logical structure of messageformat

// Record substructures //

type StandardUniversalTestID struct { //5.6.1
	UniversalTestID       string `astm:"POS=1"`
	UniversalTestIDName   string `astm:"POS=2"`
	UniversalTestIDType   string `astm:"POS=3"`
	ManufacturersTestType string `astm:"POS=4"`
}
type ExtendedUniversalTestID struct { //5.6.1
	UniversalTestID       string `astm:"POS=1"`
	UniversalTestIDName   string `astm:"POS=2"`
	UniversalTestIDType   string `astm:"POS=3"`
	ManufacturersTestType string `astm:"POS=4"`
	ManufacturersTestName string `astm:"POS=5"`
	ManufacturersTestCode string `astm:"POS=6"`
	TestCode              string `astm:"POS=7"`
}

// Record structures //

type Header struct {
	MessageControlID        string    `astm:"POS=3"`  // 6.3
	AccessPassword          string    `astm:"POS=4"`  // 6.4
	SenderNameOrID          string    `astm:"POS=5"`  // 6.5
	SenderStreetAddress     string    `astm:"POS=6"`  // 6.6
	Reserved                string    `astm:"POS=7"`  // 6.7
	SenderTelephone         string    `astm:"POS=8"`  // 6.8
	CharacteristicsOfSender string    `astm:"POS=9"`  // 6.9
	ReceiverID              string    `astm:"POS=10"` // 6.10
	Comment                 string    `astm:"POS=11"` // 6.11
	ProcessingID            string    `astm:"POS=12"` // 6.12
	Version                 string    `astm:"POS=13"` // 6.13
	DateAndTime             time.Time `astm:"POS=14"` // 6.14
}
type Patient struct {
	PracticeAssignedPatientID          string    `astm:"POS=3"`          // 7.3
	LabAssignedPatientID               string    `astm:"POS=4"`          // 7.4
	ID3                                string    `astm:"POS=5"`          // 7.5
	LastName                           string    `astm:"POS=6.1"`        // 7.6.1
	FirstName                          string    `astm:"POS=6.2"`        // 7.6.2
	MothersMaidenName                  string    `astm:"POS=7"`          // 7.7
	DOB                                time.Time `astm:"POS=8;ATR=date"` // 7.8
	Gender                             string    `astm:"POS=9"`          // 7.9
	Race                               string    `astm:"POS=10"`         // 7.10
	Address                            string    `astm:"POS=11"`         // 7.11
	F12                                string    `astm:"POS=12"`         // 7.12
	Telephone                          string    `astm:"POS=13"`         // 7.13
	AttendingPhysicianID               string    `astm:"POS=14"`         // 7.14
	SpecialField1                      string    `astm:"POS=15"`         // 7.15
	SpecialField2                      string    `astm:"POS=16"`         // 7.16
	Height                             string    `astm:"POS=17"`         // 7.17
	Weight                             string    `astm:"POS=18"`         // 7.18
	SuspectedDiagnosis                 string    `astm:"POS=19"`         // 7.19
	ActiveMedication                   string    `astm:"POS=20"`         // 7.20
	Diet                               string    `astm:"POS=21"`         // 7.21
	PracticeField1                     string    `astm:"POS=22"`         // 7.22
	PracticeField2                     string    `astm:"POS=23"`         // 7.23
	AdmissionAndDischargeDates         string    `astm:"POS=24"`         // 7.24
	AdmissionStatus                    string    `astm:"POS=25"`         // 7.25
	Location                           string    `astm:"POS=26"`         // 7.26
	NatureOfAlternativeDiagnosticCodes string    `astm:"POS=27"`         // 7.27
	AlternativeDiagnosticCodes         string    `astm:"POS=28"`         // 7.28
	Religion                           string    `astm:"POS=29"`         // 7.29
	MaritalStatus                      string    `astm:"POS=30"`         // 7.30
	IsolationStatus                    string    `astm:"POS=31"`         // 7.31
	Language                           string    `astm:"POS=32"`         // 7.32
	HospitalService                    string    `astm:"POS=33"`         // 7.33
	HospitalInstitution                string    `astm:"POS=34"`         // 7.34
	DosageCategory                     string    `astm:"POS=35"`         // 7.35
}
type Order struct {
	SpecimenID                   string                  `astm:"POS=3"`    // 8.4.3
	InstrumentSpecimenID         string                  `astm:"POS=4"`    // 8.4.4
	UniversalTestID              StandardUniversalTestID `astm:"POS=5"`    // 8.4.5
	Priority                     string                  `astm:"POS=6"`    // 8.4.6
	RequestedOrderDateTime       time.Time               `astm:"POS=7"`    // 8.4.7
	SpecimenCollectionDateTime   time.Time               `astm:"POS=8"`    // 8.4.8
	CollectionEndTime            time.Time               `astm:"POS=9"`    // 8.4.9
	CollectionVolume             string                  `astm:"POS=10"`   // 8.4.10
	CollectionID                 string                  `astm:"POS=11"`   // 8.4.11
	ActionCode                   string                  `astm:"POS=12"`   // 8.4.12
	DangerCode                   string                  `astm:"POS=13"`   // 8.4.13
	RelevantClinicalInformation  string                  `astm:"POS=14"`   // 8.4.14
	DateTimeSpecimenReceived     string                  `astm:"POS=15"`   // 8.4.15
	SpecimenType                 string                  `astm:"POS=16.1"` // 8.4.16
	SpecimenSource               string                  `astm:"POS=16.2"` // 8.4.16
	OrderingPhysician            string                  `astm:"POS=17"`   // 8.4.17
	PhysicianTelephone           string                  `astm:"POS=18"`   // 8.4.18
	UserField1                   string                  `astm:"POS=19"`   // 8.4.19
	UserField2                   string                  `astm:"POS=20"`   // 8.4.20
	LaboratoryField1             string                  `astm:"POS=21"`   // 8.4.21
	LaboratoryField2             string                  `astm:"POS=22"`   // 8.4.22
	DateTimeResultsReported      time.Time               `astm:"POS=23"`   // 8.4.23
	InstrumentCharge             string                  `astm:"POS=24"`   // 8.4.24
	InstrumentSectionID          string                  `astm:"POS=25"`   // 8.4.25
	ReportType                   string                  `astm:"POS=26"`   // 8.4.26
	Reserved                     string                  `astm:"POS=27"`   // 8.4.27
	LocationOfSpecimenCollection string                  `astm:"POS=28"`   // 8.4.28
	NosocomialInfectionFlag      string                  `astm:"POS=29"`   // 8.4.29
	SpecimenService              string                  `astm:"POS=30"`   // 8.4.30
	SpecimenInstitution          string                  `astm:"POS=31"`   // 8.4.31
}
type Result struct {
	UniversalTestID                          ExtendedUniversalTestID `astm:"POS=3"`    // 9.3
	DataMeasurementValue                     string                  `astm:"POS=4.1"`  // 9.4
	InitialMeasurementValue                  string                  `astm:"POS=4.2"`  // 9.4
	MeasurementValueOfDevice                 string                  `astm:"POS=4.3"`  // 9.4
	Units                                    string                  `astm:"POS=5"`    // 9.5
	ReferenceRange                           string                  `astm:"POS=6"`    // 9.6
	ResultAbnormalFlag                       string                  `astm:"POS=7"`    // 9.7
	NatureOfAbnormalTesting                  string                  `astm:"POS=8"`    // 9.8
	ResultStatus                             string                  `astm:"POS=9"`    // 9.9
	DateOfChangeInInstrumentNormativeTesting time.Time               `astm:"POS=10"`   // 9.10
	OperatorIDPerformed                      string                  `astm:"POS=11.1"` // 9.11
	OperatorIDVerified                       string                  `astm:"POS=11.2"` // 9.11
	DateTimeTestStarted                      time.Time               `astm:"POS=12"`   // 9.12
	DateTimeCompleted                        time.Time               `astm:"POS=13"`   // 9.13
	InstrumentIdentification                 string                  `astm:"POS=14"`   // 9.14
}
type Query struct {
	StartingRangeIDNumber           string `astm:"POS=3"`  // 11.3
	EndingRangeIDNumber             string `astm:"POS=4"`  // 11.4
	UniversalTestID                 string `astm:"POS=5"`  // 11.5
	NatureOfRequestTimeLimits       string `astm:"POS=6"`  // 11.6
	BeginningRequestResultsDateTime string `astm:"POS=7"`  // 11.7
	EndingRequestResultsDateTime    string `astm:"POS=8"`  // 11.8
	RequestingPhysicianName         string `astm:"POS=9"`  // 11.9
	RequestingPhysicianTelephone    string `astm:"POS=10"` // 11.10
	UserField1                      string `astm:"POS=11"` // 11.11
	UserField2                      string `astm:"POS=12"` // 11.12
	RequestInformationStatus        string `astm:"POS=13"` // 11.13
}
type Comment struct {
	CommentSource string `astm:"POS=3"` // 10.3
	CommentText   string `astm:"POS=4"` // 10.4
	CommentType   string `astm:"POS=5"` // 10.5
}
type Manufacturer struct {
	F3  string `astm:"POS=3"`  // 14.3
	F4  string `astm:"POS=4"`  // 14.4
	F5  string `astm:"POS=5"`  // 14.5
	F6  string `astm:"POS=6"`  // 14.6
	F7  string `astm:"POS=7"`  // 14.7
	F8  string `astm:"POS=8"`  // 14.8
	F9  string `astm:"POS=9"`  // 14.9
	F10 string `astm:"POS=10"` // 14.10
	F11 string `astm:"POS=11"` // 14.11
	F12 string `astm:"POS=12"` // 14.12
	F13 string `astm:"POS=13"` // 14.13
	F14 string `astm:"POS=14"` // 14.14
}
type Terminator struct { //Hasta la vista...
	TerminatorCode string `astm:"POS=3"` // 12.3
}

// Message structures //

type PatientGroup struct {
	Patient     Patient      `astm:"TAG=P"`
	Comments    []Comment    `astm:"TAG=C;ATR=optional"`
	OrderGroups []OrderGroup `astm:"GROUP"`
}
type OrderGroup struct {
	Order        Order         `astm:"TAG=O"`
	ResultGroups []ResultGroup `astm:"GROUP"`
}
type ResultGroup struct {
	Result   Result    `astm:"TAG=R"`
	Comments []Comment `astm:"TAG=C;ATR=optional"`
}
type PatientOrder struct {
	Patient Patient `astm:"TAG=P"`
	Orders  []Order `astm:"TAG=O"`
}

// Messages //

type ResultMessage struct {
	Header        Header         `astm:"TAG=H"`
	Manufacturer  Manufacturer   `astm:"TAG=M;ATR=optional"`
	PatientGroups []PatientGroup `astm:"GROUP"`
	Terminator    Terminator     `astm:"TAG=L"`
}
type ResultMultiMessage struct {
	ResultMessages []ResultMessage `astm:"GROUP"`
}
type QueryMessage struct {
	Header     Header     `astm:"TAG=H"`
	Queries    []Query    `astm:"TAG=Q"`
	Terminator Terminator `astm:"TAG=L"`
}
type OrderMessage struct {
	Header        Header         `astm:"TAG=H"`
	PatientOrders []PatientOrder `astm:"GROUP"`
	Terminator    Terminator     `astm:"TAG=L"`
}
