package hl7

import (
	"testing"

	"github.com/blutspende/go-parser/pconfig"
)

// Configuration struct for tests
var config *pconfig.Configuration

// Reset config to default values
func teardown() {
	config = pconfig.NewDefaultConfiguration(pconfig.HL7)
	_ = pconfig.InitConfig(config)
}

// Setup default config and run all tests
func TestMain(m *testing.M) {
	// Set up configuration
	teardown()
	// Run all tests
	m.Run()
}
