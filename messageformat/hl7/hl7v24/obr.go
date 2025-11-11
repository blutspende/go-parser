package hl7v24

import "time"

// HL7 v2.4 - OBR - Observation Reques
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OBR
type OBR struct {
	SetID                                int       `hl7:"POS=2;ATR=sequence" json:"SetID,omitempty"`
	PlacerOrderNumber                    EI        `hl7:"POS=3" json:"PlacerOrderNumber,omitempty"`
	FillerOrderNumber                    EI        `hl7:"POS=4" json:"FillerOrderNumber,omitempty"`
	UniversalServiceIdentifier           CE        `hl7:"POS=5" json:"UniversalServiceIdentifier,omitempty"`
	Priority                             string    `hl7:"POS=6" json:"Priority,omitempty"`
	RequestedDateTime                    time.Time `hl7:"POS=7" json:"RequestedDateTime,omitempty"`
	ObservationDateTime                  time.Time `hl7:"POS=8" json:"ObservationDateTime,omitempty"`
	ObservationEndDateTime               time.Time `hl7:"POS=9" json:"ObservationEndDateTime,omitempty"`
	CollectionVolume                     CQ        `hl7:"POS=10" json:"CollectionVolume,omitempty"`
	CollectorIdentifier                  []XCN     `hl7:"POS=11" json:"CollectorIdentifier,omitempty"`
	SpecimenActionCode                   string    `hl7:"POS=12" json:"SpecimenActionCode,omitempty"`
	DangerCode                           CE        `hl7:"POS=13" json:"DangerCode,omitempty"`
	RelevantClinicalInformation          string    `hl7:"POS=14" json:"RelevantClinicalInformation,omitempty"`
	SpecimenReceivedDateTime             time.Time `hl7:"POS=15" json:"SpecimenReceivedDateTime,omitempty"`
	SpecimenSource                       SPS       `hl7:"POS=16" json:"SpecimenSource,omitempty"`
	OrderingProvider                     []XCN     `hl7:"POS=17" json:"OrderingProvider,omitempty"`
	OrderCallbackPhoneNumber             []XTN     `hl7:"POS=18" json:"OrderCallbackPhoneNumber,omitempty"`
	PlacerField1                         string    `hl7:"POS=19" json:"PlacerField1,omitempty"`
	PlacerField2                         string    `hl7:"POS=20" json:"PlacerField2,omitempty"`
	FillerField1                         string    `hl7:"POS=21" json:"FillerField1,omitempty"`
	FillerField2                         string    `hl7:"POS=22" json:"FillerField2,omitempty"`
	ResultsRptStatusChngDateTime         time.Time `hl7:"POS=23" json:"ResultsRptStatusChngDateTime,omitempty"`
	ChargeToPractice                     MOC       `hl7:"POS=24" json:"ChargeToPractice,omitempty"`
	DiagnosticServiceSectionID           string    `hl7:"POS=25" json:"DiagnosticServiceSectionID,omitempty"`
	ResultStatus                         string    `hl7:"POS=26" json:"ResultStatus,omitempty"`
	ParentResult                         PRL       `hl7:"POS=27" json:"ParentResult,omitempty"`
	QuantityTiming                       []TQ      `hl7:"POS=28" json:"QuantityTiming,omitempty"`
	ResultCopiesTo                       []XCN     `hl7:"POS=29" json:"ResultCopiesTo,omitempty"`
	ParentNumber                         EIP       `hl7:"POS=30" json:"ParentNumber,omitempty"`
	TransportationMode                   string    `hl7:"POS=31" json:"TransportationMode,omitempty"`
	ReasonForStudy                       []CE      `hl7:"POS=32" json:"ReasonForStudy,omitempty"`
	PrincipalResultInterpreter           NDL       `hl7:"POS=33" json:"PrincipalResultInterpreter,omitempty"`
	AssistantResultInterpreter           []NDL     `hl7:"POS=34" json:"AssistantResultInterpreter,omitempty"`
	Technician                           []NDL     `hl7:"POS=35" json:"Technician,omitempty"`
	Transcriptionist                     []NDL     `hl7:"POS=36" json:"Transcriptionist,omitempty"`
	ScheduledDateTime                    time.Time `hl7:"POS=37" json:"ScheduledDateTime,omitempty"`
	NumberOfSampleContainers             *int      `hl7:"POS=38" json:"NumberOfSampleContainers,omitempty"`
	TransportLogisticsOfCollectedSample  []CE      `hl7:"POS=39" json:"TransportLogisticsOfCollectedSample,omitempty"`
	CollectorsComment                    []CE      `hl7:"POS=40" json:"CollectorsComment,omitempty"`
	TransportArrangementResponsibility   CE        `hl7:"POS=41" json:"TransportArrangementResponsibility,omitempty"`
	TransportArranged                    string    `hl7:"POS=42" json:"TransportArranged,omitempty"`
	EscortRequired                       string    `hl7:"POS=43" json:"EscortRequired,omitempty"`
	PlannedPatientTransportComment       []CE      `hl7:"POS=44" json:"PlannedPatientTransportComment,omitempty"`
	ProcedureCode                        CE        `hl7:"POS=45" json:"ProcedureCode,omitempty"`
	ProcedureCodeModifier                []CE      `hl7:"POS=46" json:"ProcedureCodeModifier,omitempty"`
	PlacerSupplementalServiceInformation []CE      `hl7:"POS=47" json:"PlacerSupplementalServiceInformation,omitempty"`
	FillerSupplementalSErviceInformation []CE      `hl7:"POS=48" json:"FillerSupplementalSErviceInformation,omitempty"`
}
