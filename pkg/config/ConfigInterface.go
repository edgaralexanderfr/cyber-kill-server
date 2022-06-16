package config

type ConfigInterface interface {
	GetConfig(cache bool) *Config
}
