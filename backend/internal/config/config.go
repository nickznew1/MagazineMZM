package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Env          string        `yaml:"env"`
	DbConfig     []DatabaseCfg `yaml:"database"`
	ServerConfig []ServerCfg   `yaml:"server"`
	ClientConfig []ClientCfg   `yaml:"client"`
}

type DatabaseCfg struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Ssl      string `yaml:"ssl"`
	User     string `yaml:"user"`
}

type ServerCfg struct {
	Port string `yaml:"port"`
}

type ClientCfg struct {
	Url string `yaml:"api_url"`
}

func MustLoad(path string) *Config {
	var cfg Config

	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("File doesn't exist on %s path", path)
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Error when try to parse config: %s", err)
	}

	return &cfg
}
