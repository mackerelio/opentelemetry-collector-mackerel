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
}

func (cfg *Config) Validate() error {
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
	} else {
		return "", errors.New("Mackerel API key must be specified") //nolint:staticcheck // Mackerel is proper noun
	}
}
