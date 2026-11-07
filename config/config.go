package config

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

//go:embed defaults/config.toml
var defaultConfig embed.FS

type Config struct {
	Application ApplicationConfig `toml:"application"`
	Controller  ControllerConfig  `toml:"controller"`
	Database    DatabaseConfig    `toml:"database"`
	Logging     LoggingConfig     `toml:"logging"`
	Network     NetworkConfig     `toml:"network"`
	Nodes       []NodeConfig      `toml:"nodes"`
}

type ApplicationConfig struct {
	Name        string `toml:"name"`
	Version     string `toml:"version"`
	Environment string `toml:"environment"`
	Debug       bool   `toml:"debug"`
}

type ControllerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type NodeConfig struct {
	Name    string `toml:"name"`
	Address string `toml:"address"`
	Port    int    `toml:"port"`
	Enabled bool   `toml:"enabled"`
}

type DatabaseConfig struct {
	Type string `toml:"type"`
	Path string `toml:"path"`
}

type LoggingConfig struct {
	Level  string `toml:"level"`
	Format string `toml:"format"`
}

type NetworkConfig struct {
	BindAddress      string `toml:"bind_address"`
	AdvertiseAddress string `toml:"advertise_address"`
}

func Load(path string) (*Config, error) {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		data, err := defaultConfig.ReadFile("defaults/config.toml")
		if err != nil {
			return nil, fmt.Errorf("failed to read default config: %w", err)
		}

		if err := os.WriteFile(path, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to create config file: %w", err)
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}
