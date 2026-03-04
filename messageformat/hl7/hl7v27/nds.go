package hl7v27

import "time"

// NDS - Notification Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/NDS
type NDS struct {
	NotificationReferenceNumber *int `hl7:"POS=2"`
	NotificationDateTime time.Time `hl7:"POS=3"`
	NotificationAlertSeverity CWE `hl7:"POS=4"`
	NotificationCode CWE `hl7:"POS=5"`
}

