package hl7

import (
	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/timezone"
	"github.com/blutspende/go-parser"

	"testing"
	"time"

	"github.com/blutspende/go-parser/messages/hl7/hl7v23"
	"github.com/stretchr/testify/assert"
)

func TestMarshalMSH(t *testing.T) {
	// Arrange
	mshData := "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|0||ER|ER||8859/1|\u000d"
	filedata := mshData
	var err error
	var message hl7v23.ORM_O01
	// TODO: remove cross testing
	err = parser.Unmarshal([]byte(filedata), &message, config)
	config.Encoding = encoding.ASCII
	// Act
	var marshalledMessageBytes [][]byte
	marshalledMessageBytes, err = parser.Marshal(message, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mshData, string(marshalledMessageBytes[0]))
	// Teardown
	teardown()
}

func TestMarshalFromStruct(t *testing.T) {
	// Arrange
	mshData := "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20230203080903||ORM^O01|CID586246|P|2.3|0||ER|SU||8859/1|\u000d"
	trigger := hl7v23.ORM_O01{}
	msh := hl7v23.MSH{}
	msh.ReceivingApplication = hl7v23.HD{NamespaceId: "CIT"}
	msh.ReceivingFacility = hl7v23.HD{NamespaceId: "LAB"}
	msh.SendingApplication = hl7v23.HD{NamespaceId: "HL7_Host"}
	msh.SendingFacility = hl7v23.HD{NamespaceId: "HL7_Office"}
	msh.DateTimeOfMessage, _ = time.Parse("20060102150405", "20230203080903")
	msh.MessageType = "ORM"
	msh.MessageTriggerEvent = "O01"
	msh.MessageControlID = "CID586246"
	msh.ProcessingID = "P"
	msh.VersionID = "2.3"
	msh.SequenceNumber = 0
	msh.AcceptAcknowledgementType = "ER"
	msh.ApplicationAcknowledgementType = "SU"
	msh.CharacterSet = []string{"8859/1"}
	trigger.MSH = msh

	trigger.Order = append(trigger.Order, hl7v23.Order{})
	{
		orc := &trigger.Order[0].CommonOrderSegment
		{
			orc.OrderControl = "NW"
			orc.FillerOrderNumber.EntityIdentifier = "23071012"
			orc.DateTimeOfTransaction = time.Now()
			orc.OrderingProvider.ID = "AKB"
		}
		orb := &trigger.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment
		{
			orb.ObservationRequest = "1"
			orb.FillerOrderNumber.EntityIdentifier = "FIL4345"
			orb.UniversalServiceIdentifier.Identifier = "DNA-A"
			orb.UniversalServiceIdentifier.Text = "Loki A am DNA Strang"
			orb.RequestedDateTime = time.Now()
		}
	}
	trigger.Order = append(trigger.Order, hl7v23.Order{})
	{
		orc := &trigger.Order[1].CommonOrderSegment
		{
			orc.OrderControl = "NW"
			orc.PlacerOrderNumber.EntityIdentifier = "23071012"
			orc.DateTimeOfTransaction = time.Now()
			orc.OrderingProvider.ID = "AKB"
		}
		orb := &trigger.Order[1].OrderDetail.OrderDetailSegment.ObservationRequestSegment
		{
			orb.ObservationRequest = "2"
			orb.PlacerOrderNumber.EntityIdentifier = "FIL4345"
			orb.UniversalServiceIdentifier.Identifier = "DNA-DRQB"
			orb.UniversalServiceIdentifier.Text = "DRQB Loki"
			orb.RequestedDateTime = time.Now()
		}
	}
	pid := &trigger.Patient.PatientIdentification
	{
		pid.InternalID = append(pid.InternalID, hl7v23.CX{})
		pid.InternalID[0].Id = "01020304"
		pid.Name.FamilyName = "Nachnamäh"
		pid.Name.GivenName = "Vörname"
		pid.Sex = "U"
	}
	pv1 := &trigger.Patient.PatientVisit.PatientVisit
	{
		pv1.VisitNumber.Id = "VID001"
		pv1.PatientClass = "S"
	}
	config.Encoding = encoding.ASCII
	config.TimeZone = timezone.UTC
	// Act
	marshalledMessageBytes, err := parser.Marshal(trigger, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mshData, string(marshalledMessageBytes[0]))
	// Teardown
	teardown()
}

func TestMarshalPID(t *testing.T) {
	// Arrange
	mshData := "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|0||ER|ER||8859/1|\u000d"
	pidData := "PID|1|a^b~^c|00100M56016||Smith^Harry||19500412|M\u000d"
	filedata := mshData + pidData
	var message hl7v23.ORM_O01
	// TODO: remove cross testing
	err := parser.Unmarshal([]byte(filedata), &message, config)
	config.Encoding = encoding.ASCII
	// Act
	marshalledMessageBytes, err := parser.Marshal(message, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, pidData, string(marshalledMessageBytes[1]))
	// Teardown
	teardown()
}
