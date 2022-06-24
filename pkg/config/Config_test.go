package config

import (
	"testing"
)

var configFactory ConfigFactoryInterface = &ConfigFactory{}
var config ConfigInterface = configFactory.NewConfig()

func TestConfigGetConfig(t *testing.T) {
	if config == nil {
		t.Fatalf(`configFactory.NewConfig() = nil, want ConfigInterface`)
	}

	t.Logf(`Validating config.GetConfig(true):`)

	conf := config.GetConfig(true)

	if conf == nil {
		t.Fatalf(`config.GetConfig(true) = nil, want *Config`)
	}

	validateConfig(t, conf)

	t.Logf(`Validating config.GetConfig(true) one more time:`)

	conf = config.GetConfig(true)

	if conf == nil {
		t.Fatalf(`config.GetConfig(true) = nil, want *Config`)
	}

	validateConfig(t, conf)

	t.Logf(`Validating config.GetConfig(false):`)

	conf = config.GetConfig(false)

	if conf == nil {
		t.Fatalf(`config.GetConfig(false) = nil, want *Config`)
	}

	validateConfig(t, conf)
}

func validateConfig(t *testing.T, conf *Config) {
	if conf.Port == 0 {
		t.Fatalf(`*Config.Port is zero, want non-zero`)
	}

	if conf.Server.TickRate == 0 {
		t.Fatalf(`*Config.Server.TickRate is zero, want non-zero`)
	}
}
