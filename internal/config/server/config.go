package server

import "time"

type Config struct {
	BindAddr     string        `toml:"bind_addr"`
	LogLevel     string        `toml:"log_level"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:     ":8080",
		LogLevel:     "debug",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
