# Mackerel OpenTelemetry Collector

This distribution is an OpenTelemetry Collector built for Mackerel users.

Even if you don't know config file format fof OpenTelemetry Collector, you can start a simple collector with just a few environment variables.

## Components

We have selected components provided from the OpenTelemetry community that are particularly beneficial.

If you are a Mackerel user and would like to add OpenTelemetry Collector components to this distribution, please let us know by opening an issue.

### Exporters

| Component      | Description            | Document                                                                                                           |
| -------------- | ---------------------- | ------------------------------------------------------------------------------------------------------------------ |
| `debug`        | Debug Exporter         | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/debugexporter)             |
| `nop`          | No-op Exporter         | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/nopexporter)               |
| `otlp`         | OTLP/gRPC Exporter     | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlpexporter)              |
| `otlphttp`     | OTLP/HTTP Exporter     | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlphttpexporter)          |
| `mackerelotlp` | Mackerel OTLP Exporter | [Document](https://github.com/mackerelio/opentelemetry-collector-mackerel/tree/main/exporter/mackerelotlpexporter) |

### Processors

| Component               | Description                      | Document                                                                                                                        |
| ----------------------- | -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `batch`                 | Batch Processor                  | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/batchprocessor)                        |
| `memory_limiter`        | Memory Limiter Processor         | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/memorylimiterprocessor)                |
| `attributes`            | Attributes Processor             | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/attributesprocessor)           |
| `filter`                | Filter Processor                 | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/filterprocessor)               |
| `probabilistic_sampler` | Probabilistic Sampling Processor | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/probabilisticsamplerprocessor) |
| `resourcedetection`     | Resource Detection Processor     | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/resourcedetectionprocessor)    |
| `resource`              | Resource Processor               | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/resourceprocessor)             |
| `span`                  | Span Processor                   | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/spanprocessor)                 |
| `tail_sampling`         | Tail Sampling Processor          | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/tailsamplingprocessor)         |
| `transform`             | Transform Processor              | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/processor/transformprocessor/README.md)  |

### Receivers

| Component     | Description           | Document                                                                                                             |
| ------------- | --------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `nop`         | No-op Receiver        | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/receiver/nopreceiver)                 |
| `otlp`        | OTLP Receiver         | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/receiver/otlpreceiver)                |
| `hostmetrics` | Host Metrics Receiver | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver) |
| `httpcheck`   | HTTP Check Receiver   | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/httpcheckreceiver)   |
| `mysql`       | MySQL Receiver        | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/mysqlreceiver)       |
| `oracledb`    | Oracle DB Receiver    | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/oracledbreceiver)    |
| `postgresql`  | PostgreSQL Receiver   | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/postgresqlreceiver)  |
| `redis`       | Redis Receiver        | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/redisreceiver)       |

### Connectors

| Component      | Description             | Document                                                                                                                |
| -------------- | ----------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `routing`      | Routing Connector       | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/connector/routingconnector)      |
| `servicegraph` | Service Graph Connector | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/connector/servicegraphconnector) |
| `spanmetrics`  | Span Metrics Connector  | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/connector/spanmetricsconnector)  |
