package parser

import (
	"errors"
	"regexp"

	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/messagetype"
	"github.com/blutspende/go-parser/functions"
	"github.com/blutspende/go-parser/models"
	"github.com/blutspende/go-parser/parserconfig"
)

var ErrIdentifyMessageWrongProtocol = errors.New("protocol mismatch between function and configuration")

func IdentifyMessageASTM(messageData []byte, config *parserconfig.Configuration) (messageType messagetype.MessageType, err error) {
	// Init configuration
	err = parserconfig.InitConfig(config)
	if err != nil {
		return "", err
	}
	// Check for correct protocol
	if config.Protocol != parserconfig.ASTM {
		return "", ErrIdentifyMessageWrongProtocol
	}
	// Convert encoding to UTF8
	utf8Data, err := encoding.ConvertFromEncodingToUtf8(messageData, config.Encoding)
	if err != nil {
		return "", err
	}
	// Split the message data into lines
	lines, err := functions.SliceLines(utf8Data, config)
	if err != nil {
		return "", err
	}
	// Extract the first characters from each line
	firstChars := ""
	for _, line := range lines {
		if len(line) > 0 {
			firstChars += string(line[0])
		}
	}
	// TODO: verify these regexes to be correct
	// Set up the possible message types regexes
	expressionQuery := "^(HQ+)+L?$"
	expressionOrder := "^(H(PM?C?M?OM?C?M?)+)+L?$"
	expressionOrderAndResult := "^(H(PM*C?M*OM*C?M*(RM*C?M*)+)+)+L?$"
	expressionManyOrderAndResult := "^(H(PM*C?M*(OM*C?M*(RM*C?M*)+)*)+)L?$"
	// Check the first characters against the regexes and return the message type
	switch {
	case regexp.MustCompile(expressionQuery).MatchString(firstChars):
		return messagetype.Query, nil
	case regexp.MustCompile(expressionOrder).MatchString(firstChars):
		return messagetype.Order, nil
	case regexp.MustCompile(expressionOrderAndResult).MatchString(firstChars):
		return messagetype.Result, nil
	case regexp.MustCompile(expressionManyOrderAndResult).MatchString(firstChars):
		return messagetype.Result, nil
	}
	// If no match was found return unknown
	return messagetype.Unidentified, err
}

type hl7Header struct {
	messageType     string `hl7:"8"`
	protocolVersion string `hl7:"11"`
}

func IdentifyMessageHL7(messageData []byte, config *parserconfig.Configuration) (messageType string, protocolVersion string, err error) {
	// Init configuration
	err = parserconfig.InitConfig(config)
	if err != nil {
		return "", "", err
	}
	// Check for correct protocol
	if config.Protocol != parserconfig.HL7 {
		return "", "", ErrIdentifyMessageWrongProtocol
	}
	// Convert encoding to UTF8
	utf8Data, err := encoding.ConvertFromEncodingToUtf8(messageData, config.Encoding)
	if err != nil {
		return "", "", err
	}
	// Split the message data into lines
	lines, err := functions.SliceLines(utf8Data, config)
	if err != nil {
		return "", "", err
	}
	// Parse the MSH segment as a single header line
	var header hl7Header
	annotation := models.StructAnnotation{
		StructName: "MSH",
	}
	_, err = functions.ParseLine(lines[0], header, annotation, 1, config)
	if err != nil {
		return "", "", err
	}
	// Save the protocol version to the config
	config.ProtocolVersion = header.protocolVersion
	// Return the extracted message type and protocol version from the header struct
	return header.messageType, header.protocolVersion, nil
}
