package mackerelotlpexporter

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

func createDefaultConfig() component.Config {
	return &Config{
		// overrides default exporter timeout config
		// because transport to Mackerel may take longer than 5 seconds,
		// which is default value in official OTLP exporters.
		TimeoutConfig: exporterhelper.TimeoutConfig{
			Timeout: 10 * time.Second,
		},
		QueueConfig: exporterhelper.NewDefaultQueueConfig(),
		RetryConfig: configretry.NewDefaultBackOffConfig(),
	}
}
