package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Database *DatabaseConfig `yaml:"database"`
	Server   *ServerConfig   `yaml:"server"`
	Logger   *LoggerConfig   `yaml:"logger"`
	User     *UserConfig     `yaml:"user"`
}

// DatabaseConfig ...
type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}

// ServerConfig ...
type ServerConfig struct {
	Port string `yaml:"port"`
}

// LoggerConfig ...
type LoggerConfig struct {
	Level string `yaml:"level"`
}

// UserConfig ...
type UserConfig struct {
	Salt      string `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

// ReadConfig ...
func ReadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var conf Config

	if err := yaml.NewDecoder(file).Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
