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

		Web struct {
			Port types.Port `json:"port"`
			Dir  string     `json:"dir"`
		} `json:"web"`
	} `json:"server"`

	Map struct {
		Width  uint16 `json:"width"`
		Height uint16 `json:"height"`
		Size   uint16 `json:"size"`
	} `json:"map"`
}

func (config Config) GetConfig(cacheable bool) *Config {
	cfg, found := c.Get("config")
	if cacheable && found {
		return cfg.(*Config)
	}
	configFile, err := os.Open("config.json")
	if err != nil {
		return nil
	}
	defer configFile.Close()
	newCfg := Config{}
	json.NewDecoder(configFile).Decode(&newCfg)
	if cacheable {
		c.Set("config", newCfg, cache.NoExpiration)
	}
	return &newCfg
}
