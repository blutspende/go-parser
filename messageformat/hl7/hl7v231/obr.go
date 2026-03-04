package hl7v231

import "time"

// OBR - Observation request segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/OBR
type OBR struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	PlacerOrderNumber EI `hl7:"POS=3"`
	FillerOrderNumber EI `hl7:"POS=4"`
	UniversalServiceID CE `hl7:"POS=5"`
	PriorityObr string `hl7:"POS=6"`
	RequestedDateTime time.Time `hl7:"POS=7"`
	ObservationDateTime time.Time `hl7:"POS=8"`
	ObservationEndDateTime time.Time `hl7:"POS=9"`
	CollectionVolume CQ `hl7:"POS=10"`
	CollectorIdentifier []XCN `hl7:"POS=11"`
	SpecimenActionCode string `hl7:"POS=12"`
	DangerCode CE `hl7:"POS=13"`
	RelevantClinicalInfo string `hl7:"POS=14"`
	SpecimenReceivedDateTime time.Time `hl7:"POS=15"`
	SpecimenSource SPS `hl7:"POS=16"`
	OrderingProvider []XCN `hl7:"POS=17"`
	OrderCallbackPhoneNumber []XTN `hl7:"POS=18"`
	PlacerField1 string `hl7:"POS=19"`
	PlacerField2 string `hl7:"POS=20"`
	FillerField1 string `hl7:"POS=21"`
	FillerField2 string `hl7:"POS=22"`
	ResultsRptStatusChangeDateTime time.Time `hl7:"POS=23"`
	ChargeToPractice MOC `hl7:"POS=24"`
	DiagnosticServSectID string `hl7:"POS=25"`
	ResultStatus string `hl7:"POS=26"`
	ParentResult PRL `hl7:"POS=27"`
	QuantityTiming []TQ `hl7:"POS=28"`
	ResultCopiesTo []XCN `hl7:"POS=29"`
	ParentNumber EIP `hl7:"POS=30"`
	TransportationMode string `hl7:"POS=31"`
	ReasonForStudy []CE `hl7:"POS=32"`
	PrincipalResultInterpreter NDL `hl7:"POS=33"`
	AssistantResultInterpreter []NDL `hl7:"POS=34"`
	Technician []NDL `hl7:"POS=35"`
	Transcriptionist []NDL `hl7:"POS=36"`
	ScheduledDateTime time.Time `hl7:"POS=37"`
	NumberOfSampleContainers *int `hl7:"POS=38"`
	TransportLogisticsOfCollectedSample []CE `hl7:"POS=39"`
	CollectorsComment []CE `hl7:"POS=40"`
	TransportArrangementResponsibility CE `hl7:"POS=41"`
	TransportArranged string `hl7:"POS=42"`
	EscortRequired string `hl7:"POS=43"`
	PlannedPatientTransportComment []CE `hl7:"POS=44"`
	ProcedureCode CE `hl7:"POS=45"`
	ProcedureCodeModifier []CE `hl7:"POS=46"`
}

