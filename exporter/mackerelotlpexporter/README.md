# Mackerel OTLP Exporter

Export OpenTelemetry Metrics/Traces from OpenTelemetry Collector to [Mackerel](https://mackerel.io/) using OTLP (OpenTelemetry Protocol).

Mackerel natively supports OTLP, so we can send telemetry using the [OTLP gRPC Exporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlpexporter) or [OTLP HTTP Exporter](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlphttpexporter) provided by the OpenTelemetry community.

However, Mackerel requires different endpoints and OTLP transport types for each telemetry type. This necessitates setting up multiple exporters, which is inconvenient.

Using the Mackerel OTLP Exporter, you can send both traces and metrics with the same exporter.

## Getting Started

Pass the following configuration to the OpenTelemetry Collector bundled with the Mackerel OTLP Exporter:

```yaml
exporters:
  mackerelotlp:
```

Here is an example of the whole settings:

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
      http:
        endpoint: 0.0.0.0:4318

processors:
  batch:
    send_batch_size: 5000
    send_batch_max_size: 5000

exporters:
  mackerelotlp:

service:
  pipelines:
    metrics:
      receivers: [otlp]
      processors: [batch]
      exporters: [mackerelotlp]
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [mackerelotlp]
```

Set the Mackerel writable API key in the `MACKEREL_APIKEY` environment variable and run the OpenTelemetry Collector.
