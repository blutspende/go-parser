package hl7

import (
	"testing"

	"github.com/blutspende/go-parser/parserconfig"
)

// Configuration struct for tests
var config *parserconfig.Configuration

// Reset config to default values
func teardown() {
	config = parserconfig.NewDefaultConfiguration(parserconfig.HL7)
	_ = parserconfig.InitConfig(config)
}

// Setup default config and run all tests
func TestMain(m *testing.M) {
	// Set up configuration
	teardown()
	// Run all tests
	m.Run()
}
