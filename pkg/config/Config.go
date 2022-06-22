package config

import (
	"encoding/json"
	"os"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/types"
	"github.com/pmylund/go-cache"
)

var c = cache.New(cache.NoExpiration, cache.NoExpiration)

type Config struct {
	Port types.Port `json:"port"`

	Server struct {
		TickRate byte `json:"tickRate"`
	} `json:"server"`
}

func (config Config) GetConfig(cacheable bool) *Config {
	if &config == nil {
		configFile, err := os.Open("config.json")
		if err != nil {
			return nil
		}
		defer configFile.Close()
		json.NewDecoder(configFile).Decode(&config)
	}
	if cacheable {
		config, found := c.Get("config")
		if found {
			return config.(*Config)
		} else {
			c.Set("config", &config, cache.NoExpiration)
		}
	}
	return &config
}
