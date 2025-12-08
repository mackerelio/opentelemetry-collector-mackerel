package mackerelotlpexporter

import (
	"errors"
	"os"

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
	// Generally, a deb package should successfully start the service it provides upon installation.
	// For now, this exporter needs an API key which prevents us from adhering to this rule.
	//
	// However, Validate() should not return an error.
	// If it does, the collector process will exit,
	// and systemd will then attempt to restart the failed process repeatedly.
	// In this scenario, since the error cannot be resolved due to the missing API key,
	// the unit eventually reach its start limit.
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
	} else {
		return "", errors.New("Mackerel API key must be specified") //nolint:staticcheck // Mackerel is proper noun
	}
}
