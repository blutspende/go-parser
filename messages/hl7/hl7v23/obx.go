package hl7v23

import "time"

// OBX - Observation segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OBX
type OBX struct {
	SetID                        int       `hl7:"1,sequence" json:"setID,omitempty"`
	ValueType                    string    `hl7:"2" json:"valueType,omitempty"`
	ObservationIdentifier        CE        `hl7:"3" json:"observationIdentifier,omitempty"`
	ObservationSubID             string    `hl7:"4" json:"observationSubID,omitempty"`
	ObservationValue             []string  `hl7:"5" json:"observationValue,omitempty"`
	Units                        CE        `hl7:"6" json:"units,omitempty"`
	ReferenceRange               string    `hl7:"7" json:"referenceRange,omitempty"`
	AbnormalFlags                []string  `hl7:"8" json:"abnormalFlags,omitempty"`
	Probability                  float32   `hl7:"9" json:"probability,omitempty"`
	NatureOfAbnormalTest         []string  `hl7:"10" json:"natureOfAbnormalTest,omitempty"`
	ResultStatus                 string    `hl7:"11" json:"resultStatus,omitempty"`
	DateLastObservedNormalValues time.Time `hl7:"12,longdate" json:"dateLastObservedNormalValues,omitempty"`
	UserDefinedAccessChecks      string    `hl7:"13" json:"userDefinedAccessChecks,omitempty"`
	DateTimeOfObservation        time.Time `hl7:"14,longdate" json:"dateTimeOfObservation,omitempty"`
	ProducersID                  CE        `hl7:"15" json:"producersID,omitempty"`
	ResponsibleObserver          XCN       `hl7:"16" json:"responsibleObserver,omitempty"`
	ObservationMethod            []CE      `hl7:"17" json:"observationMethod,omitempty"`
}
