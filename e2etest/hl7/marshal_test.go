package hl7

import (
	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/timezone"
	"github.com/blutspende/bloodlab-common/utils"
	"github.com/blutspende/go-parser"
	"github.com/blutspende/go-parser/enums/notation"

	"testing"
	"time"

	"github.com/blutspende/go-parser/messages/hl7/hl7v23"
	"github.com/stretchr/testify/assert"
)

func TestMarshalFromStruct(t *testing.T) {
	// Arrange
	dt, _ := time.Parse("20060102150405", "20230203080903")
	source := hl7v23.ORM_O01{
		MSH: hl7v23.MSH{
			ReceivingApplication:           hl7v23.HD{NamespaceId: "CIT"},
			ReceivingFacility:              hl7v23.HD{NamespaceId: "LAB"},
			SendingApplication:             hl7v23.HD{NamespaceId: "HL7_Host"},
			SendingFacility:                hl7v23.HD{NamespaceId: "HL7_Office"},
			DateTimeOfMessage:              dt,
			MessageType:                    "ORM",
			MessageTriggerEvent:            "O01",
			MessageControlID:               "CID586246",
			ProcessingID:                   "P",
			VersionID:                      "2.3",
			SequenceNumber:                 nil,
			AcceptAcknowledgementType:      "ER",
			ApplicationAcknowledgementType: "SU",
			CharacterSet:                   []string{"8859/1"},
		},
		Order: []hl7v23.Order{
			{
				CommonOrderSegment: hl7v23.ORC{
					OrderControl: "NW",
					FillerOrderNumber: hl7v23.EI{
						EntityIdentifier: "23071012",
					},
					DateTimeOfTransaction: dt,
					OrderingProvider: hl7v23.XCN{
						ID: "AKB",
					},
				},
				OrderDetail: hl7v23.OrderDetail{
					OrderDetailSegment: hl7v23.OrderDetailSegment{
						ObservationRequestSegment: hl7v23.OBR{
							ObservationRequest: "1",
							FillerOrderNumber: hl7v23.EI{
								EntityIdentifier: "FIL4345",
							},
							UniversalServiceIdentifier: hl7v23.CE{
								Identifier: "DNA-A",
								Text:       "Loki A am DNA Strang",
							},
							RequestedDateTime: dt,
						},
					},
				},
			},
			{
				CommonOrderSegment: hl7v23.ORC{
					OrderControl: "NW",
					FillerOrderNumber: hl7v23.EI{
						EntityIdentifier: "23071012",
					},
					DateTimeOfTransaction: dt,
					OrderingProvider: hl7v23.XCN{
						ID: "AKB",
					},
				},
				OrderDetail: hl7v23.OrderDetail{
					OrderDetailSegment: hl7v23.OrderDetailSegment{
						ObservationRequestSegment: hl7v23.OBR{
							ObservationRequest: "2",
							FillerOrderNumber: hl7v23.EI{
								EntityIdentifier: "FIL4345",
							},
							UniversalServiceIdentifier: hl7v23.CE{
								Identifier: "DNA-DRQB",
								Text:       "DRQB Loki",
							},
							RequestedDateTime: dt,
						},
					},
				},
			},
		},
		Patient: hl7v23.Patient{
			PatientIdentification: hl7v23.PID{
				InternalID: []hl7v23.CX{
					{
						Id: "01020304",
					},
				},
				Name: hl7v23.XPN{
					FamilyName: "Nachnamäh",
					GivenName:  "Vörname",
				},
				Sex: "U",
			},
			PatientVisit: hl7v23.PatientVisit{
				PatientVisit: hl7v23.PV1{
					VisitNumber: hl7v23.CX{
						Id: "VID001",
					},
					PatientClass: "S",
				},
			},
		},
	}
	expected := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20230203080903||ORM^O01|CID586246|P|2.3|||ER|SU||8859/1
PID|0||01020304||Nachnamäh^Vörname|||U
PD1
PV1||S|||||||||||||||||VID001
PV2|
GT1|
ORC|NW||23071012||||||20230203080903|||AKB
OBR|1||FIL4345|DNA-A^Loki A am DNA Strang||20230203080903
RQD|
RQ1|
RQ1|
ODS|
ODT|
BLG|
ORC|NW||23071012||||||20230203080903|||AKB
OBR|2||FIL4345|DNA-DRQB^DRQB Loki||20230203080903
RQD|
RQ1|
RQ1|
ODS|
ODT|
BLG|`
	config.Encoding = encoding.ASCII
	config.TimeZone = timezone.UTC
	config.Notation = notation.Short
	// Act
	marshaled, err := parser.Marshal(source, config)
	result := string(utils.ConvertBytes2Dto1D(marshaled))
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	// Teardown
	teardown()
}

func TestMarshalPID(t *testing.T) {
	// Arrange
	dt, _ := time.Parse("20060102150405", "20110926125155")
	dob, _ := time.Parse("20060102", "19500412")
	source := hl7v23.ORM_O01{
		MSH: hl7v23.MSH{
			ReceivingApplication:           hl7v23.HD{NamespaceId: "CIT"},
			ReceivingFacility:              hl7v23.HD{NamespaceId: "LAB"},
			SendingApplication:             hl7v23.HD{NamespaceId: "HL7_Host"},
			SendingFacility:                hl7v23.HD{NamespaceId: "HL7_Office"},
			DateTimeOfMessage:              dt.UTC(),
			MessageType:                    "ORM",
			MessageTriggerEvent:            "O01",
			MessageControlID:               "CID586246",
			ProcessingID:                   "P",
			VersionID:                      "2.3",
			SequenceNumber:                 nil,
			AcceptAcknowledgementType:      "ER",
			ApplicationAcknowledgementType: "ER",
			CharacterSet:                   []string{"8859/1"},
		},
		Patient: hl7v23.Patient{
			PatientIdentification: hl7v23.PID{
				ID: 1,
				InternalID: []hl7v23.CX{
					{
						Id:         "a",
						CheckDigit: "b",
					},
					{
						CheckDigit: "c",
					},
				},
				AlternateID: []hl7v23.CX{
					{
						Id: "00100M56016",
					},
				},
				Name: hl7v23.XPN{
					FamilyName: "Smith",
					GivenName:  "Harry",
				},
				DOB: dob,
				Sex: "M",
			},
		},
	}
	expected0 := "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926145155||ORM^O01|CID586246|P|2.3|||ER|ER||8859/1"
	expected1 := "PID|1||a^b~^c|00100M56016|Smith^Harry||19500412|M"
	config.Encoding = encoding.ASCII
	// Act
	result, err := parser.Marshal(source, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expected0, string(result[0]))
	assert.Equal(t, expected1, string(result[1]))
	// Teardown
	teardown()
}
