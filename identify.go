package parser

import (
	"regexp"

	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/instrumentenum"
	"github.com/blutspende/go-parser/errdef"
	"github.com/blutspende/go-parser/functions"
	"github.com/blutspende/go-parser/pconfig"
)

func IdentifyMessageASTM(messageData []byte, config *pconfig.Configuration) (messageType instrumentenum.MessageType, err error) {
	// Init configuration
	err = pconfig.InitConfig(config)
	if err != nil {
		return "", err
	}
	// Check for correct protocol
	if config.Protocol != pconfig.ASTM {
		return "", errdef.ErrIdentifyMessageWrongProtocol
	}
	// Convert encoding to UTF8
	utf8Data, err := encoding.ConvertFromEncodingToUTF8(messageData, config.Encoding)
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
	expressionMultimessageResult := "^((H(PM*C?M*(OM*C?M*(RM*C?M*)+)*)+)L?)+$"
	// Check the first characters against the regexes and return the message type
	switch {
	case regexp.MustCompile(expressionQuery).MatchString(firstChars):
		return instrumentenum.MessageTypeQuery, nil
	case regexp.MustCompile(expressionOrder).MatchString(firstChars):
		return instrumentenum.MessageTypeOrder, nil
	case regexp.MustCompile(expressionOrderAndResult).MatchString(firstChars):
		return instrumentenum.MessageTypeResult, nil
	case regexp.MustCompile(expressionManyOrderAndResult).MatchString(firstChars):
		return instrumentenum.MessageTypeResult, nil
	case regexp.MustCompile(expressionMultimessageResult).MatchString(firstChars):
		return instrumentenum.MessageTypeResult, nil
	}
	// If no match was found, return unknown
	return instrumentenum.MessageTypeUnidentified, err
}

func IdentifyMessageHL7(messageData []byte, config *pconfig.Configuration) (messageType string, protocolVersion string, err error) {
	// Init configuration
	err = pconfig.InitConfig(config)
	if err != nil {
		return "", "", err
	}
	// Check for correct protocol
	if config.Protocol != pconfig.HL7 {
		return "", "", errdef.ErrIdentifyMessageWrongProtocol
	}
	// Convert encoding to UTF8
	utf8Data, err := encoding.ConvertFromEncodingToUTF8(messageData, config.Encoding)
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
			MessageType     string `hl7:"POS=9"`
			ProtocolVersion string `hl7:"POS=12"`
		} `hl7:"TAG=MSH"`
	}
	var headerMsg HeaderMessage
	err = functions.ParseStruct(lines, &headerMsg, config)
	if err != nil {
		return "", "", err
	}
	// Check for missing data
	if headerMsg.Header.MessageType == "" {
		return "", "", errdef.ErrIdentifyMessageMissingMessageType
	}
	// Note: option to also error on missing protocol version
	//if headerMsg.Header.ProtocolVersion == "" {
	//	return "", "", errdef.ErrIdentifyMessageMissingProtocolVersion
	//}
	// Return the extracted message type and protocol version from the header struct
	return headerMsg.Header.MessageType, headerMsg.Header.ProtocolVersion, nil
}
