package hl7v23

import "time"

// OBR - Observation request segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OBR
type OBR struct {
	ObservationRequest         string `hl7:"2" json:"observationRequest,omitempty"`
	PlacerOrderNumber          EI     `hl7:"3" json:"placerOrderNumber,omitempty"`
	FillerOrderNumber          EI     `hl7:"4" json:"fillerOrderNumber,omitempty"`
	UniversalServiceIdentifier CE     `hl7:"5" json:"universalServiceIdentifier,omitempty"`
	Priority                   string `hl7:"6" json:"priority,omitempty"`
	// TODO: is longdate really correct here?
	RequestedDateTime                   time.Time `hl7:"7,longdate" json:"requestedDateTime,omitempty"`
	ObservationDateTime                 time.Time `hl7:"8,longdate" json:"observationDateTime,omitempty"`
	ObservationEndDateTime              time.Time `hl7:"9,longdate" json:"observationEndDateTime,omitempty"`
	CollectionVolume                    CQ        `hl7:"10" json:"collectionVolume,omitempty"`
	CollectorIdentifier                 []XCN     `hl7:"11" json:"collectorIdentifier,omitempty"`
	SpecimenActionCode                  string    `hl7:"12" json:"specimenActionCode,omitempty"`
	DangerCode                          CE        `hl7:"13" json:"dangerCode,omitempty"`
	RelevantClinicalInformation         string    `hl7:"14" json:"relevantClinicalInformation,omitempty"`
	SpecimenReceivedDateTime            time.Time `hl7:"15" json:"specimenReceivedDateTime,omitempty"`
	SpecimenSource                      CM_SPS    `hl7:"16" json:"specimenSource,omitempty"`
	OrderingProvider                    []XCN     `hl7:"17" json:"orderingProvider,omitempty"`
	OrderCallbackPhoneNumber            []XTN     `hl7:"18" json:"orderCallbackPhoneNumber,omitempty"`
	PlacerField1                        string    `hl7:"19" json:"placerField1,omitempty"`
	PlacerField2                        string    `hl7:"20" json:"placerField2,omitempty"`
	FillerField1                        string    `hl7:"21" json:"fillerField1,omitempty"`
	FillerField2                        string    `hl7:"22" json:"fillerField2,omitempty"`
	ResultsRptStatusChngDateTime        time.Time `hl7:"23" json:"resultsRptStatusChngDateTime,omitempty"`
	ChargeToPractice                    CM_MOC    `hl7:"24" json:"chargeToPractice,omitempty"`
	DiagnosticServiceSectionID          string    `hl7:"25" json:"diagnosticServiceSectionID,omitempty"`
	ResultStatus                        string    `hl7:"26" json:"resultStatus,omitempty"`
	ParentResult                        CM_PRL    `hl7:"27" json:"parentResult,omitempty"`
	QuantityTiming                      []TQ      `hl7:"28" json:"quantityTiming,omitempty"`
	ResultCopiesTo                      []XCN     `hl7:"29" json:"resultCopiesTo,omitempty"`
	ParentNumber                        CM_EIP    `hl7:"30" json:"parentNumber,omitempty"`
	TransportationMode                  string    `hl7:"31" json:"transportationMode,omitempty"`
	ReasonForStudy                      []CE      `hl7:"32" json:"reasonForStudy,omitempty"`
	PrincipalResultInterpreter          CM_NDL    `hl7:"33" json:"principalResultInterpreter,omitempty"`
	AssistantResultInterpreter          []CM_NDL  `hl7:"34" json:"assistantResultInterpreter,omitempty"`
	Technician                          []CM_NDL  `hl7:"35" json:"technician,omitempty"`
	Transcriptionist                    []CM_NDL  `hl7:"36" json:"transcriptionist,omitempty"`
	ScheduledDateTime                   time.Time `hl7:"37" json:"scheduledDateTime,omitempty"`
	NumberOfSampleContainers            int       `hl7:"38" json:"numberOfSampleContainers,omitempty"`
	TransportLogisticsOfCollectedSample []CE      `hl7:"39" json:"transportLogisticsOfCollectedSample,omitempty"`
	CollectorsComment                   []CE      `hl7:"40" json:"collectorsComment,omitempty"`
	TransportArrangementResponsibility  CE        `hl7:"41" json:"transportArrangementResponsibility,omitempty"`
	TransportArranged                   string    `hl7:"42" json:"transportArranged,omitempty"`
	EscortRequired                      string    `hl7:"43" json:"escortRequired,omitempty"`
	PlannedPatientTransportComment      []CE      `hl7:"44" json:"plannedPatientTransportComment,omitempty"`
}
