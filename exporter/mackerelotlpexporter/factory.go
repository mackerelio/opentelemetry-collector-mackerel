package mackerelotlpexporter

import (
	"time"

	"github.com/mackerelio/opentelemetry-collector-mackerel/exporter/mackerelotlpexporter/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/xexporter"
)

const (
	defaultTimeout           = 10 * time.Second
	defaultBatchMaxSizeBytes = 5_000_000 // 5MB
	defaultMetricsEndpoint   = "otlp.mackerelio.com:4317"
	defaultTracesEndpoint    = "https://otlp-vaxila.mackerelio.com"
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
	queueConfig := exporterhelper.NewDefaultQueueConfig()

	// overrides default exporter queue batch config
	// because Vaxila endpoint does not accept requests larger than 6MB.
	queueBatchConfig := queueConfig.Batch.GetOrInsertDefault()
	queueBatchConfig.Sizer = exporterhelper.RequestSizerTypeBytes
	queueBatchConfig.MaxSize = defaultBatchMaxSizeBytes

	queueConfig.Batch = configoptional.Default(*queueBatchConfig)

	return &Config{
		// overrides default exporter timeout config
		// because transport to Mackerel may take longer than 5 seconds,
		// which is default value in official OTLP exporters.
		TimeoutConfig: exporterhelper.TimeoutConfig{
			Timeout: defaultTimeout,
		},
		QueueConfig:     queueConfig,
		RetryConfig:     configretry.NewDefaultBackOffConfig(),
		MetricsEndpoint: defaultMetricsEndpoint,
		TracesEndpoint:  defaultTracesEndpoint,
	}
}
