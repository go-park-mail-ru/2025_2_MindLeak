package server

import "time"

type Config struct {
	BindAddr     string        `toml:"bind_addr"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:     ":8090",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
}
