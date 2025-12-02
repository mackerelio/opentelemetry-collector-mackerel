package mackerelotlpexporter

import (
	"errors"
	"os"
	"path/filepath"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

type Config struct {
	TimeoutConfig  exporterhelper.TimeoutConfig    `mapstructure:",squash"`
	QueueConfig    exporterhelper.QueueBatchConfig `mapstructure:"sending_queue"`
	RetryConfig    configretry.BackOffConfig       `mapstructure:"retry_on_failure"`
	MackerelApiKey configopaque.String             `mapstructure:"mackerel_api_key"`
	// MetricsEndpoint configurations are provided for testing purposes and may be modified or deprecated.
	MetricsEndpoint string `mapstructure:"metrics_endpoint"`
	// TracesEndpoint configurations are provided for testing purposes and may be modified or deprecated.
	TracesEndpoint string `mapstructure:"traces_endpoint"`
}

func (cfg *Config) Validate() error {
	if err := cfg.TimeoutConfig.Validate(); err != nil {
		return err
	}
	if err := cfg.QueueConfig.Validate(); err != nil {
		return err
	}
	if err := cfg.RetryConfig.Validate(); err != nil {
		return err
	}
	if _, err := cfg.mackerelApiKey(); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) mackerelApiKey() (configopaque.String, error) {
	if cfg.MackerelApiKey != "" {
		return cfg.MackerelApiKey, nil
	}
	if v := os.Getenv("MACKEREL_APIKEY"); v != "" {
		return configopaque.String(v), nil
	} else if v := os.Getenv("MACKEREL_API_KEY"); v != "" {
		return configopaque.String(v), nil
	} else if v, err := readCredentialFrom("mackerel-apikey"); err == nil {
		return configopaque.String(v), nil
	} else {
		return "", errors.New("Mackerel API key must be specified") //nolint:staticcheck // Mackerel is proper noun
	}
}

func readCredentialFrom(name string) (string, error) {
	dir := os.Getenv("CREDENTIALS_DIRECTORY")
	if dir == "" {
		return "", os.ErrNotExist
	}
	file := filepath.Join(dir, name)
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
