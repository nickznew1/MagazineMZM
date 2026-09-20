package core_transport_http_cors

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr string `envconfig:"ADDR"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("FRONTEND", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()

	if err != nil {
		err = fmt.Errorf("get frontend-server config: %w", err)
		panic(err)
	}
	return config
}
