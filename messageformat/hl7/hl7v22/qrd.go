package hl7v22

import "time"

// QRD - Query Definition
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/QRD
type QRD struct {
	QueryDateTime time.Time `hl7:"POS=2"`
	QueryFormatCode string `hl7:"POS=3"`
	QueryPriority string `hl7:"POS=4"`
	QueryID string `hl7:"POS=5"`
	DeferredResponseType string `hl7:"POS=6"`
	DeferredResponseDateTime time.Time `hl7:"POS=7"`
	QuantityLimitedRequest CQ `hl7:"POS=8"`
	WhoSubjectFilter []string `hl7:"POS=9"`
	WhatSubjectFilter []string `hl7:"POS=10"`
	WhatDepartmentDataCode []string `hl7:"POS=11"`
	WhatDataCodeValueQualifier []CM_VR `hl7:"POS=12"`
	QueryResultsLevel string `hl7:"POS=13"`
}

