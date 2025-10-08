package hl7

import (
	"testing"

	"github.com/blutspende/go-parser"
	"github.com/stretchr/testify/assert"
)

func TestMessageIdentification(t *testing.T) {
	// Arrange
	data := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3`
	// Act
	messageType, protocolVersion, err := parser.IdentifyMessageHL7([]byte(data), config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "ORM^O01", messageType)
	assert.Equal(t, "2.3", protocolVersion)
}
