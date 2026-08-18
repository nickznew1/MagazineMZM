package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
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

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg, err
}
