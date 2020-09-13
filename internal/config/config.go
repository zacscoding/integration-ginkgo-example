package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server   *ServerConfig   `yaml:"server"`
	Database *DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Endpoint string `yaml:"endpoint"`
	Pool     struct {
		MaxOpen     int `yaml:"maxOpen"`
		MaxIdle     int `yaml:"maxIdle"`
		MaxLifeTime int `yaml:"maxLifeTime"`
	} `yaml:"pool"`
	CreateTable bool `yaml:"createTable"`
}

// Load load config from given config file path
func Load(path string) (*Config, error) {
	c := Config{
		Server:   &ServerConfig{},
		Database: &DatabaseConfig{},
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
