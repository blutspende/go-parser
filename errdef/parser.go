package errdef

import "errors"

// IdentifyMessage
var (
	ErrIdentifyMessageWrongProtocol          = errors.New("protocol mismatch between function and configuration")
	ErrIdentifyMessageMissingMessageType     = errors.New("missing message type from hl7 header")
	ErrIdentifyMessageMissingProtocolVersion = errors.New("missing protocol version from hl7 header")
	ErrIdentifyInvalidRegex                  = errors.New("invalid regex in identify rule")
)
