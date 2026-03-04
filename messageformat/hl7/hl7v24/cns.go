package hl7v24

import "time"

// CNS - Clear Notification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CNS
type CNS struct {
	StartingNotificationReferenceNumber *int `hl7:"POS=2"`
	EndingNotificationReferenceNumber *int `hl7:"POS=3"`
	StartingNotificationDateTime time.Time `hl7:"POS=4"`
	EndingNotificationDateTime time.Time `hl7:"POS=5"`
	StartingNotificationCode CE `hl7:"POS=6"`
	EndingNotificationCode CE `hl7:"POS=7"`
}

