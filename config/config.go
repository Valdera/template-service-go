package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config struct to implement model of inbox configuration
type Config struct {
	GRPCPort string `envconfig:"GRPC_PORT" default:"9090"`
	RESTPort string `envconfig:"REST_PORT" default:"8080"`
	HOST     string `envconfig:"HOST" default:""`
}

// Get to get defined configuration
func Get() Config {
	_ = godotenv.Load()
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
