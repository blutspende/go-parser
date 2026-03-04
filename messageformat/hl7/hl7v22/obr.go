package hl7v22

import "time"

// OBR - Observation Request
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/OBR
type OBR struct {
	SetID                             int         `hl7:"POS=2;ATR=sequence"`
	PlacerOrderNumber                 CM_PLACER   `hl7:"POS=3"`
	FillerOrderNumber                 CM_FILLER   `hl7:"POS=4"`
	UniversalServiceID                CE          `hl7:"POS=5"`
	Priority                          string      `hl7:"POS=6"`
	RequestedDateTime                 time.Time   `hl7:"POS=7"`
	ObservationDateTime               time.Time   `hl7:"POS=8"`
	ObservationEndDateTime            time.Time   `hl7:"POS=9"`
	CollectionVolume                  CQ          `hl7:"POS=10"`
	CollectorIdentifier               []CN_PERSON `hl7:"POS=11"`
	SpecimenActionCode                string      `hl7:"POS=12"`
	DangerCode                        CE          `hl7:"POS=13"`
	RelevantClinicalInformation       string      `hl7:"POS=14"`
	SpecimenReceivedDateTime          time.Time   `hl7:"POS=15"`
	SpecimenSource                    CM_SPS      `hl7:"POS=16"`
	OrderingProvider                  []CN_PERSON `hl7:"POS=17"`
	OrderCallbackPhoneNumber          []string    `hl7:"POS=18"`
	PlacerField1                      string      `hl7:"POS=19"`
	PlacerField2                      string      `hl7:"POS=20"`
	FillerField1                      string      `hl7:"POS=21"`
	FillerField2                      string      `hl7:"POS=22"`
	ResultsReportStatusChangeDateTime time.Time   `hl7:"POS=23"`
	ChargeToPractice                  CM_MOC      `hl7:"POS=24"`
	DiagnosticServiceSectionID        string      `hl7:"POS=25"`
	ResultStatus                      string      `hl7:"POS=26"`
	ParentResult                      CM_PRL      `hl7:"POS=27"`
	QuantityTiming                    []TQ        `hl7:"POS=28"`
	ResultCopiesTo                    []CN_PERSON `hl7:"POS=29"`
	ParentNumber                      CM_EIP      `hl7:"POS=30"`
	TransportationMode                string      `hl7:"POS=31"`
	ReasonForStudy                    []CE        `hl7:"POS=32"`
	PrincipalResultInterpreter        CM_NDL      `hl7:"POS=33"`
	AssistantResultInterpreter        []CM_NDL    `hl7:"POS=34"`
	Technician                        []CM_NDL    `hl7:"POS=35"`
	Transcriptionist                  []CM_NDL    `hl7:"POS=36"`
	ScheduledDateTime                 time.Time   `hl7:"POS=37"`
}
