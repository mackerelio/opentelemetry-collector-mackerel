package mackerelotlpexporter

import (
	"time"

	"github.com/mackerelio/opentelemetry-collector-mackerel/exporter/mackerelotlpexporter/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/xexporter"
)

func NewFactory() exporter.Factory {
	return xexporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		xexporter.WithTraces(createTraces, metadata.TracesStability),
		xexporter.WithMetrics(createMetrics, metadata.MetricsStability),
	)
}

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
