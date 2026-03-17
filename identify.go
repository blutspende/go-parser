package parser

import (
	"regexp"

	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/instrumentenum"
	"github.com/blutspende/go-parser/errdef"
	"github.com/blutspende/go-parser/functions"
	"github.com/blutspende/go-parser/pconfig"
	"github.com/rs/zerolog/log"
)

func IdentifyMessageASTM(messageData []byte, config *pconfig.Configuration) (messageType instrumentenum.MessageType, err error) {
	// Init configuration
	err = pconfig.InitConfig(config)
	if err != nil {
		return "", err
	}
	// Verify protocol
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
	// Extract the signature of the message
	signature := ""
	for _, line := range lines {
		if len(line) > 0 {
			firstChar := string(line[0])
			// Filter out comments and manufacturer records
			if firstChar == "C" || firstChar == "M" {
				continue
			}
			// Collect first chars
			signature += firstChar
		}
	}
	// Set up the default ruleset and add custom rules from config
	rules := []pconfig.ASTMIdentifyRule{
		{"Query", "HQ+L", instrumentenum.MessageTypeQuery},
		{"Order", "H(PO+)+L", instrumentenum.MessageTypeOrder},
		{"Result", "H(P(OR+)+)+L", instrumentenum.MessageTypeResult},
		{"Result (multimessage)", "(H(P(OR+)+)+L)+", instrumentenum.MessageTypeResult},
	}
	rules = append(rules, config.CustomASTMIdentifyRules...)
	// Check the message signature with all the match rules to try to find a match
	for _, rule := range rules {
		regex := regexp.MustCompile("^" + rule.Regex + "$")
		if regex.MatchString(signature) {
			log.Debug().Str("signature", signature).Interface("rule", rule).Msg("message identified")
			return rule.Type, nil
		}
	}
	// Return the default message type unidentified with no error
	log.Debug().Str("signature", signature).Msg("message could not be identified")
	return instrumentenum.MessageTypeUnidentified, nil
}

func IdentifyMessageHL7(messageData []byte, config *pconfig.Configuration) (messageType string, protocolVersion string, err error) {
	// Init configuration
	err = pconfig.InitConfig(config)
	if err != nil {
		return "", "", err
	}
	// Verify protocol
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
	log.Debug().Str("MessageType", headerMsg.Header.MessageType).Str("ProtocolVersion", headerMsg.Header.ProtocolVersion).Msg("message identified")
	return headerMsg.Header.MessageType, headerMsg.Header.ProtocolVersion, nil
}
