package hl7v23

import "time"

// OBX - Observation segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OBX
type OBX struct {
	SetID                        int       `hl7:"POS=2;ATR=sequence" json:"setID,omitempty"`
	ValueType                    string    `hl7:"POS=3" json:"valueType,omitempty"`
	ObservationIdentifier        CE        `hl7:"POS=4" json:"observationIdentifier,omitempty"`
	ObservationSubID             string    `hl7:"POS=5" json:"observationSubID,omitempty"`
	ObservationValue             []string  `hl7:"POS=6" json:"observationValue,omitempty"`
	Units                        CE        `hl7:"POS=7" json:"units,omitempty"`
	ReferenceRange               string    `hl7:"POS=8" json:"referenceRange,omitempty"`
	AbnormalFlags                []string  `hl7:"POS=9" json:"abnormalFlags,omitempty"`
	Probability                  *float32  `hl7:"POS=10" json:"probability,omitempty"`
	NatureOfAbnormalTest         []string  `hl7:"POS=11" json:"natureOfAbnormalTest,omitempty"`
	ResultStatus                 string    `hl7:"POS=12" json:"resultStatus,omitempty"`
	DateLastObservedNormalValues time.Time `hl7:"POS=13" json:"dateLastObservedNormalValues,omitempty"`
	UserDefinedAccessChecks      string    `hl7:"POS=14" json:"userDefinedAccessChecks,omitempty"`
	DateTimeOfObservation        time.Time `hl7:"POS=15" json:"dateTimeOfObservation,omitempty"`
	ProducersID                  CE        `hl7:"POS=16" json:"producersID,omitempty"`
	ResponsibleObserver          XCN       `hl7:"POS=17" json:"responsibleObserver,omitempty"`
	ObservationMethod            []CE      `hl7:"POS=18" json:"observationMethod,omitempty"`
}
