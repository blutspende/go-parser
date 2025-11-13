package pconfig

import (
	"errors"
)

func NewDefaultConfiguration(protocol Protocol) *Configuration {
	var newConfig Configuration
	switch protocol {
	case ASTM:
		newConfig = DefaultConfigurationASTM
	case HL7:
		newConfig = DefaultConfigurationHL7
	default:
		return nil
	}
	return &newConfig
}

func InitConfig(config *Configuration) (err error) {
	if config == nil {
		return errors.New("nil configuration")
	}
	if config.Protocol != ASTM && config.Protocol != HL7 {
		return errors.New("invalid protocol")
	}
	if config.Delimiters.Field == "" ||
		config.Delimiters.Repeat == "" ||
		config.Delimiters.Component == "" ||
		config.Delimiters.Escape == "" ||
		(config.Protocol == HL7 && config.Delimiters.SubComponent == "") {
		return errors.New("missing delimiter(s)")
	}
	config.TimeLocation, err = config.TimeZone.GetLocation()
	return err
}
