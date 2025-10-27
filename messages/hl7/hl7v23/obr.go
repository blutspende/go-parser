package hl7v23

import "time"

// OBR - Observation request segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OBR
type OBR struct {
	ObservationRequest                  string    `hl7:"POS=2" json:"observationRequest,omitempty"`
	PlacerOrderNumber                   EI        `hl7:"POS=3" json:"placerOrderNumber,omitempty"`
	FillerOrderNumber                   EI        `hl7:"POS=4" json:"fillerOrderNumber,omitempty"`
	UniversalServiceIdentifier          CE        `hl7:"POS=5" json:"universalServiceIdentifier,omitempty"`
	Priority                            string    `hl7:"POS=6" json:"priority,omitempty"`
	RequestedDateTime                   time.Time `hl7:"POS=7" json:"requestedDateTime,omitempty"`
	ObservationDateTime                 time.Time `hl7:"POS=8" json:"observationDateTime,omitempty"`
	ObservationEndDateTime              time.Time `hl7:"POS=9" json:"observationEndDateTime,omitempty"`
	CollectionVolume                    CQ        `hl7:"POS=10" json:"collectionVolume,omitempty"`
	CollectorIdentifier                 []XCN     `hl7:"POS=11" json:"collectorIdentifier,omitempty"`
	SpecimenActionCode                  string    `hl7:"POS=12" json:"specimenActionCode,omitempty"`
	DangerCode                          CE        `hl7:"POS=13" json:"dangerCode,omitempty"`
	RelevantClinicalInformation         string    `hl7:"POS=14" json:"relevantClinicalInformation,omitempty"`
	SpecimenReceivedDateTime            time.Time `hl7:"POS=15" json:"specimenReceivedDateTime,omitempty"`
	SpecimenSource                      CM_SPS    `hl7:"POS=16" json:"specimenSource,omitempty"`
	OrderingProvider                    []XCN     `hl7:"POS=17" json:"orderingProvider,omitempty"`
	OrderCallbackPhoneNumber            []XTN     `hl7:"POS=18" json:"orderCallbackPhoneNumber,omitempty"`
	PlacerField1                        string    `hl7:"POS=19" json:"placerField1,omitempty"`
	PlacerField2                        string    `hl7:"POS=20" json:"placerField2,omitempty"`
	FillerField1                        string    `hl7:"POS=21" json:"fillerField1,omitempty"`
	FillerField2                        string    `hl7:"POS=22" json:"fillerField2,omitempty"`
	ResultsRptStatusChngDateTime        time.Time `hl7:"POS=23" json:"resultsRptStatusChngDateTime,omitempty"`
	ChargeToPractice                    CM_MOC    `hl7:"POS=24" json:"chargeToPractice,omitempty"`
	DiagnosticServiceSectionID          string    `hl7:"POS=25" json:"diagnosticServiceSectionID,omitempty"`
	ResultStatus                        string    `hl7:"POS=26" json:"resultStatus,omitempty"`
	ParentResult                        CM_PRL    `hl7:"POS=27" json:"parentResult,omitempty"`
	QuantityTiming                      []TQ      `hl7:"POS=28" json:"quantityTiming,omitempty"`
	ResultCopiesTo                      []XCN     `hl7:"POS=29" json:"resultCopiesTo,omitempty"`
	ParentNumber                        CM_EIP    `hl7:"POS=30" json:"parentNumber,omitempty"`
	TransportationMode                  string    `hl7:"POS=31" json:"transportationMode,omitempty"`
	ReasonForStudy                      []CE      `hl7:"POS=32" json:"reasonForStudy,omitempty"`
	PrincipalResultInterpreter          CM_NDL    `hl7:"POS=33" json:"principalResultInterpreter,omitempty"`
	AssistantResultInterpreter          []CM_NDL  `hl7:"POS=34" json:"assistantResultInterpreter,omitempty"`
	Technician                          []CM_NDL  `hl7:"POS=35" json:"technician,omitempty"`
	Transcriptionist                    []CM_NDL  `hl7:"POS=36" json:"transcriptionist,omitempty"`
	ScheduledDateTime                   time.Time `hl7:"POS=37" json:"scheduledDateTime,omitempty"`
	NumberOfSampleContainers            *int      `hl7:"POS=38" json:"numberOfSampleContainers,omitempty"`
	TransportLogisticsOfCollectedSample []CE      `hl7:"POS=39" json:"transportLogisticsOfCollectedSample,omitempty"`
	CollectorsComment                   []CE      `hl7:"POS=40" json:"collectorsComment,omitempty"`
	TransportArrangementResponsibility  CE        `hl7:"POS=41" json:"transportArrangementResponsibility,omitempty"`
	TransportArranged                   string    `hl7:"POS=42" json:"transportArranged,omitempty"`
	EscortRequired                      string    `hl7:"POS=43" json:"escortRequired,omitempty"`
	PlannedPatientTransportComment      []CE      `hl7:"POS=44" json:"plannedPatientTransportComment,omitempty"`
}
