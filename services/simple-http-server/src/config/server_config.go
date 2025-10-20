package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const serverPort = "8080"

const path = "config/server_config.yaml"

type ServerConfig struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func LoadServerConfig() (*ServerConfig, error) {
	var c ServerConfig
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(file, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (cfg *ServerConfig) GetServerPortAsString() string {
	if cfg.Server.Port == "" {
		return serverPort
	}
	return cfg.Server.Port
}
