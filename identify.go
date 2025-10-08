package parser

import (
	"regexp"

	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/messagetype"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/functions"
	"github.com/blutspende/go-parser/parserconfig"
)

func IdentifyMessageASTM(messageData []byte, config *parserconfig.Configuration) (messageType messagetype.MessageType, err error) {
	// Init configuration
	err = parserconfig.InitConfig(config)
	if err != nil {
		return "", err
	}
	// Check for correct protocol
	if config.Protocol != parserconfig.ASTM {
		return "", errmsg.ErrIdentifyMessageWrongProtocol
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

func IdentifyMessageHL7(messageData []byte, config *parserconfig.Configuration) (messageType string, protocolVersion string, err error) {
	// Init configuration
	err = parserconfig.InitConfig(config)
	if err != nil {
		return "", "", err
	}
	// Check for correct protocol
	if config.Protocol != parserconfig.HL7 {
		return "", "", errmsg.ErrIdentifyMessageWrongProtocol
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
	// Parse the MSH segment as a single header structure
	type HeaderMessage struct {
		Header struct {
			MessageType     string `hl7:"9"`
			ProtocolVersion string `hl7:"12"`
		} `hl7:"MSH"`
	}
	var headerMsg HeaderMessage
	lineIndex := 0
	err = functions.ParseStruct(lines, &headerMsg, &lineIndex, 1, 0, config)
	if err != nil {
		return "", "", err
	}
	// Check for missing data
	if headerMsg.Header.MessageType == "" {
		return "", "", errmsg.ErrIdentifyMessageMissingMessageType
	}
	if headerMsg.Header.ProtocolVersion == "" {
		return "", "", errmsg.ErrIdentifyMessageMissingProtocolVersion
	}
	// Save the protocol version to the config
	config.ProtocolVersion = headerMsg.Header.ProtocolVersion
	// Return the extracted message type and protocol version from the header struct
	return headerMsg.Header.MessageType, headerMsg.Header.ProtocolVersion, nil
}
