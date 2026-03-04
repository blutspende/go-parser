package hl7v27

import "time"

// CNS - Clear Notification
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/CNS
type CNS struct {
	StartingNotificationReferenceNumber *int `hl7:"POS=2"`
	EndingNotificationReferenceNumber *int `hl7:"POS=3"`
	StartingNotificationDateTime time.Time `hl7:"POS=4"`
	EndingNotificationDateTime time.Time `hl7:"POS=5"`
	StartingNotificationCode CWE `hl7:"POS=6"`
	EndingNotificationCode CWE `hl7:"POS=7"`
}

