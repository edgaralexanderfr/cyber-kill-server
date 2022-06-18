package config

type ConfigFactoryInterface interface {
	NewConfig() ConfigInterface
}
