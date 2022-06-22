package config

import (
	"encoding/json"
	"os"
)

type ConfigFactory struct{}

func (configFactory *ConfigFactory) NewConfig() ConfigInterface {
	var c Config
	configFile, err := os.Open("config.json")
	if err != nil {
		return c
	}
	defer configFile.Close()
	json.NewDecoder(configFile).Decode(&c)
	return c
}
