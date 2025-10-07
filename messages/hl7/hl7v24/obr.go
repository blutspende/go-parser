package hl7v24

import "time"

// HL7 v2.4 - OBR - Observation Reques
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OBR
type OBR struct {
	SetID                                int       `hl7:"1" json:"SetID,omitempty"`
	PlacerOrderNumber                    EI        `hl7:"2" json:"PlacerOrderNumber,omitempty"`
	FillerOrderNumber                    EI        `hl7:"3" json:"FillerOrderNumber,omitempty"`
	UniversalServiceIdentifier           CE        `hl7:"4" json:"UniversalServiceIdentifier,omitempty"`
	Priority                             string    `hl7:"5" json:"Priority,omitempty"`
	RequestedDateTime                    time.Time `hl7:"6" json:"RequestedDateTime,omitempty"`
	ObservationDateTime                  time.Time `hl7:"7" json:"ObservationDateTime,omitempty"`
	ObservationEndDateTime               time.Time `hl7:"8" json:"ObservationEndDateTime,omitempty"`
	CollectionVolume                     CQ        `hl7:"9" json:"CollectionVolume,omitempty"`
	CollectorIdentifier                  []XCN     `hl7:"10" json:"CollectorIdentifier,omitempty"`
	SpecimenActionCode                   string    `hl7:"11" json:"SpecimenActionCode,omitempty"`
	DangerCode                           CE        `hl7:"12" json:"DangerCode,omitempty"`
	RelevantClinicalInformation          string    `hl7:"13" json:"RelevantClinicalInformation,omitempty"`
	SpecimenReceivedDateTime             time.Time `hl7:"14" json:"SpecimenReceivedDateTime,omitempty"`
	SpecimenSource                       SPS       `hl7:"15" json:"SpecimenSource,omitempty"`
	OrderingProvider                     []XCN     `hl7:"16" json:"OrderingProvider,omitempty"`
	OrderCallbackPhoneNumber             []XTN     `hl7:"17" json:"OrderCallbackPhoneNumber,omitempty"`
	PlacerField1                         string    `hl7:"18" json:"PlacerField1,omitempty"`
	PlacerField2                         string    `hl7:"19" json:"PlacerField2,omitempty"`
	FillerField1                         string    `hl7:"20" json:"FillerField1,omitempty"`
	FillerField2                         string    `hl7:"21" json:"FillerField2,omitempty"`
	ResultsRptStatusChngDateTime         time.Time `hl7:"22" json:"ResultsRptStatusChngDateTime,omitempty"`
	ChargeToPractice                     MOC       `hl7:"23" json:"ChargeToPractice,omitempty"`
	DiagnosticServiceSectionID           string    `hl7:"24" json:"DiagnosticServiceSectionID,omitempty"`
	ResultStatus                         string    `hl7:"25" json:"ResultStatus,omitempty"`
	ParentResult                         PRL       `hl7:"26" json:"ParentResult,omitempty"`
	QuantityTiming                       []TQ      `hl7:"27" json:"QuantityTiming,omitempty"`
	ResultCopiesTo                       []XCN     `hl7:"28" json:"ResultCopiesTo,omitempty"`
	ParentNumber                         EIP       `hl7:"29" json:"ParentNumber,omitempty"`
	TransportationMode                   string    `hl7:"30" json:"TransportationMode,omitempty"`
	ReasonForStudy                       []CE      `hl7:"31" json:"ReasonForStudy,omitempty"`
	PrincipalResultInterpreter           NDL       `hl7:"32" json:"PrincipalResultInterpreter,omitempty"`
	AssistantResultInterpreter           []NDL     `hl7:"33" json:"AssistantResultInterpreter,omitempty"`
	Technician                           []NDL     `hl7:"34" json:"Technician,omitempty"`
	Transcriptionist                     []NDL     `hl7:"35" json:"Transcriptionist,omitempty"`
	ScheduledDateTime                    time.Time `hl7:"36" json:"ScheduledDateTime,omitempty"`
	NumberOfSampleContainers             int       `hl7:"37" json:"NumberOfSampleContainers,omitempty"`
	TransportLogisticsOfCollectedSample  []CE      `hl7:"38" json:"TransportLogisticsOfCollectedSample,omitempty"`
	CollectorsComment                    []CE      `hl7:"39" json:"CollectorsComment,omitempty"`
	TransportArrangementResponsibility   CE        `hl7:"40" json:"TransportArrangementResponsibility,omitempty"`
	TransportArranged                    string    `hl7:"41" json:"TransportArranged,omitempty"`
	EscortRequired                       string    `hl7:"42" json:"EscortRequired,omitempty"`
	PlannedPatientTransportComment       []CE      `hl7:"43" json:"PlannedPatientTransportComment,omitempty"`
	ProcedureCode                        CE        `hl7:"44" json:"ProcedureCode,omitempty"`
	ProcedureCodeModifier                []CE      `hl7:"45" json:"ProcedureCodeModifier,omitempty"`
	PlacerSupplementalServiceInformation []CE      `hl7:"46" json:"PlacerSupplementalServiceInformation,omitempty"`
	FillerSupplementalSErviceInformation []CE      `hl7:"47" json:"FillerSupplementalSErviceInformation,omitempty"`
}
