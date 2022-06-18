package config

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/types"

type Config struct {
	Port types.Port `json:"port"`

	Server struct {
		TickRate byte `json:"tickRate"`
	} `json:"server"`
}
