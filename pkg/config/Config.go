package config

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/t"

type Config struct {
	Port t.Port `json:"port"`

	Server struct {
		TickRate byte `json:"tickRate"`
	} `json:"server"`
}
