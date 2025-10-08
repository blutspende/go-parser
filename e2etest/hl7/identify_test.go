package hl7

import (
	"testing"

	"github.com/blutspende/go-parser"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/parserconfig"
	"github.com/stretchr/testify/assert"
)

func TestIdentifyMessage_AllOk(t *testing.T) {
	// Arrange
	data := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3`
	// Act
	messageType, protocolVersion, err := parser.IdentifyMessageHL7([]byte(data), config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "ORM^O01", messageType)
	assert.Equal(t, "2.3", protocolVersion)
}

func TestIdentifyMessage_MissingMessageType(t *testing.T) {
	// Arrange
	data := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155|||20110926125155|P|2.3`
	// Act
	messageType, protocolVersion, err := parser.IdentifyMessageHL7([]byte(data), config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrIdentifyMessageMissingMessageType.Error())
	assert.Equal(t, "", messageType)
	assert.Equal(t, "", protocolVersion)
}
func TestIdentifyMessage_MissingProtocolVersion(t *testing.T) {
	// Arrange
	data := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01`
	// Act
	messageType, protocolVersion, err := parser.IdentifyMessageHL7([]byte(data), config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrIdentifyMessageMissingProtocolVersion.Error())
	assert.Equal(t, "", messageType)
	assert.Equal(t, "", protocolVersion)
}
func TestIdentifyMessage_WrongProtocol(t *testing.T) {
	// Arrange
	data := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3`
	config.Protocol = parserconfig.ASTM
	// Act
	messageType, protocolVersion, err := parser.IdentifyMessageHL7([]byte(data), config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrIdentifyMessageWrongProtocol.Error())
	assert.Equal(t, "", messageType)
	assert.Equal(t, "", protocolVersion)
	// Teardown
	teardown()
}
