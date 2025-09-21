package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
)

type HTTPServerConfig struct {
	HTTPAddress     string `yaml:"http_address"`
	HandlerTimeout  int64  `yaml:"handler_timeout_sec"`
	ReadTimeoutSec  int64  `yaml:"read_timeout_sec"`
	WriteTimeoutSec int64  `yaml:"write_timeout_sec"`
	IdleTimeoutSec  int64  `yaml:"idle_timeout_sec"`
}

type Config struct {
	HTTPServer HTTPServerConfig `yaml:"http_server"`
}

type ConfigClient struct {
}

func NewConfigClient() *ConfigClient {
	return &ConfigClient{}
}

func (c *ConfigClient) getFileName() string {
	configDir := GetEnvironmentVariable("CONFIG_DIR")
	if configDir == "" {
		configDir = "./config"
	}
	switch Env() {
	case EnvDev:
		return fmt.Sprintf("%s/config_dev.yaml", configDir)
	case EnvStaging:
		return fmt.Sprintf("%s/config_staging.yaml", configDir)
	case EnvProduction:
		return fmt.Sprintf("%s/config_production.yaml", configDir)
	case EnvTest:
		return fmt.Sprintf("%s/config_test.yaml", configDir)
	default:
		panic("invalid environment")
	}
}

func (c *ConfigClient) LoadConfig() (*Config, error) {
	slog.Info("Loading config")
	// ignore error on .Load() since its only on local
	godotenv.Load()
	fileName := c.getFileName()
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	// Override config with environment variables
	// when we have api keys, etc.

	return config, nil

}

// Fx constructor
func NewConfig() *Config {
	configClient := NewConfigClient()
	config, err := configClient.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		panic(err)
	}
	return config
}
