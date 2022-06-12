package config

type Config struct {
	Port int `json:"port"`

	Server struct {
		TickRate int `json:"tickRate"`
	} `json:"server"`
}
