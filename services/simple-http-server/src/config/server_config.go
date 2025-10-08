package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func LoadServerConfig(path string) (*ServerConfig, error) {
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
