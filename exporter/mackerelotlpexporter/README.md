# Mackerel OTLP Exporter

🌎 [日本語](./README.ja.md) | English

Export OpenTelemetry Metrics/Traces from OpenTelemetry Collector to [Mackerel](https://mackerel.io/) using OTLP (OpenTelemetry Protocol).

Mackerel natively supports OTLP, so we can send telemetry using the [OTLP gRPC Exporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlpexporter) or [OTLP HTTP Exporter](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlphttpexporter) provided by the OpenTelemetry community.

However, Mackerel requires different endpoints and OTLP transport types for each telemetry type. This necessitates setting up multiple exporters, which is inconvenient.

Using the Mackerel OTLP Exporter, you can send both traces and metrics with the same exporter.

In addition, appropriate timeout and batch configurations are applied by default to match Mackerel's specifications. You do not need to add the [Batch Processor](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/batchprocessor) to your pipeline.

## Getting Started

Pass the following configuration to the OpenTelemetry Collector bundled with the Mackerel OTLP Exporter:

```yaml
exporters:
  mackerel_otlp:
```

> [!WARNING]
> As of v0.9.0, the component type name has been changed from `mackerelotlp` to `mackerel_otlp` to follow with the [naming convention](https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/coding-guidelines.md). Although the deprecated name is still supported as an alias, it may be removed in a future release.

Here is an example of the whole settings:

```yaml
receivers:
  otlp:
    protocols:
      grpc:
      http:

exporters:
  mackerel_otlp:

service:
  pipelines:
    metrics:
      receivers: [otlp]
      exporters: [mackerel_otlp]
    traces:
      receivers: [otlp]
      exporters: [mackerel_otlp]
```

Set the Mackerel writable API key in the `MACKEREL_APIKEY` environment variable and run the OpenTelemetry Collector.

## Advanced Usage

You can additionally set the following configurations:

`mackerel_api_key`: Mackerel API key (use this if you want to specify this key via a method other than environment variables)

`timeout`: [Timeout configuration](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#timeout)

`sending_queue`: [Sending queue configuration](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#sending-queue)

`retry_on_failure`: [Retry-on-failure configuration](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#retry-on-failure)
