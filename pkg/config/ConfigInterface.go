package config

type ConfigInterface interface {
	GetConfig(cacheable bool) *Config
}
