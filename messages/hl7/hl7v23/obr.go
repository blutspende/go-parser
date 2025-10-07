package hl7v23

import "time"

// OBR - Observation request segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OBR
type OBR struct {
	ObservationRequest                  string    `hl7:"1" json:"observationRequest,omitempty"`
	PlacerOrderNumber                   EI        `hl7:"2" json:"placerOrderNumber,omitempty"`
	FillerOrderNumber                   EI        `hl7:"3" json:"fillerOrderNumber,omitempty"`
	UniversalServiceIdentifier          CE        `hl7:"4" json:"universalServiceIdentifier,omitempty"`
	Priority                            string    `hl7:"5" json:"priority,omitempty"`
	RequestedDateTime                   time.Time `hl7:"6" json:"requestedDateTime,omitempty"`
	ObservationDateTime                 time.Time `hl7:"7" json:"observationDateTime,omitempty"`
	ObservationEndDateTime              time.Time `hl7:"8" json:"observationEndDateTime,omitempty"`
	CollectionVolume                    CQ        `hl7:"9" json:"collectionVolume,omitempty"`
	CollectorIdentifier                 []XCN     `hl7:"10" json:"collectorIdentifier,omitempty"`
	SpecimenActionCode                  string    `hl7:"11" json:"specimenActionCode,omitempty"`
	DangerCode                          CE        `hl7:"12" json:"dangerCode,omitempty"`
	RelevantClinicalInformation         string    `hl7:"13" json:"relevantClinicalInformation,omitempty"`
	SpecimenReceivedDateTime            time.Time `hl7:"14" json:"specimenReceivedDateTime,omitempty"`
	SpecimenSource                      CM_SPS    `hl7:"15" json:"specimenSource,omitempty"`
	OrderingProvider                    []XCN     `hl7:"16" json:"orderingProvider,omitempty"`
	OrderCallbackPhoneNumber            []XTN     `hl7:"17" json:"orderCallbackPhoneNumber,omitempty"`
	PlacerField1                        string    `hl7:"18" json:"placerField1,omitempty"`
	PlacerField2                        string    `hl7:"19" json:"placerField2,omitempty"`
	FillerField1                        string    `hl7:"20" json:"fillerField1,omitempty"`
	FillerField2                        string    `hl7:"21" json:"fillerField2,omitempty"`
	ResultsRptStatusChngDateTime        time.Time `hl7:"22" json:"resultsRptStatusChngDateTime,omitempty"`
	ChargeToPractice                    CM_MOC    `hl7:"23" json:"chargeToPractice,omitempty"`
	DiagnosticServiceSectionID          string    `hl7:"24" json:"diagnosticServiceSectionID,omitempty"`
	ResultStatus                        string    `hl7:"25" json:"resultStatus,omitempty"`
	ParentResult                        CM_PRL    `hl7:"26" json:"parentResult,omitempty"`
	QuantityTiming                      []TQ      `hl7:"27" json:"quantityTiming,omitempty"`
	ResultCopiesTo                      []XCN     `hl7:"28" json:"resultCopiesTo,omitempty"`
	ParentNumber                        CM_EIP    `hl7:"29" json:"parentNumber,omitempty"`
	TransportationMode                  string    `hl7:"30" json:"transportationMode,omitempty"`
	ReasonForStudy                      []CE      `hl7:"31" json:"reasonForStudy,omitempty"`
	PrincipalResultInterpreter          CM_NDL    `hl7:"32" json:"principalResultInterpreter,omitempty"`
	AssistantResultInterpreter          []CM_NDL  `hl7:"33" json:"assistantResultInterpreter,omitempty"`
	Technician                          []CM_NDL  `hl7:"34" json:"technician,omitempty"`
	Transcriptionist                    []CM_NDL  `hl7:"35" json:"transcriptionist,omitempty"`
	ScheduledDateTime                   time.Time `hl7:"36" json:"scheduledDateTime,omitempty"`
	NumberOfSampleContainers            int       `hl7:"37" json:"numberOfSampleContainers,omitempty"`
	TransportLogisticsOfCollectedSample []CE      `hl7:"38" json:"transportLogisticsOfCollectedSample,omitempty"`
	CollectorsComment                   []CE      `hl7:"39" json:"collectorsComment,omitempty"`
	TransportArrangementResponsibility  CE        `hl7:"40" json:"transportArrangementResponsibility,omitempty"`
	TransportArranged                   string    `hl7:"41" json:"transportArranged,omitempty"`
	EscortRequired                      string    `hl7:"42" json:"escortRequired,omitempty"`
	PlannedPatientTransportComment      []CE      `hl7:"43" json:"plannedPatientTransportComment,omitempty"`
}
