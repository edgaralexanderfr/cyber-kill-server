package config

type ConfigFactory struct{}

func (configFactory *ConfigFactory) NewConfig() ConfigInterface {
	return nil
}
