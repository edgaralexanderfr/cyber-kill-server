package config

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/types"
	"github.com/pmylund/go-cache"
)

type Config struct {
	Port types.Port `json:"port"`

	Server struct {
		TickRate byte `json:"tickRate"`
	} `json:"server"`
}

func (config Config) GetConfig(cacheable bool) *Config {
	if cacheable {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		config, found := c.Get("config")
		if found {
			return config.(*Config)
		} else {
			c.Set("config", &config, cache.NoExpiration)
		}
	}
	return &config
}
