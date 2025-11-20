# Mackerel OpenTelemetry Collector

This distribution is an OpenTelemetry Collector built for Mackerel users.

Even if you don't know config file format fof OpenTelemetry Collector, you can start a simple collector with just a few environment variables.

## Getting Started

You can start a OpenTelemetry Collector that receives telemetry via OTLP on the localhost domain, adds host resource attributes, and exports the data to Mackerel without a complex configuration file. All you need is your Mackerel API key.

### Docker

We publish images on Docker Hub and GitHub Container Registry:

- [mackerel/otelcol-mackerel](https://hub.docker.com/r/mackerel/otelcol-mackerel)
- [ghcr.io/mackerelio/opentelemetry-collector-mackerel/otelcol-mackerel](https://github.com/mackerelio/opentelemetry-collector-mackerel/pkgs/container/opentelemetry-collector-mackerel%2Fotelcol-mackerel)

```
$ docker run -e MACKEREL_APIKEY=your_api_key mackerel/otelcol-mackerel:latest
2025-11-04T13:13:41.242Z        info    builders/builders.go:26 Development component. May change in the future.        {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "mackerelotlp", "otelcol.component.kind": "exporter", "otelcol.signal": "metrics"}
2025-11-04T13:13:41.243Z        info    builders/builders.go:26 Development component. May change in the future.        {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "mackerelotlp", "otelcol.component.kind": "exporter", "otelcol.signal": "traces"}
2025-11-04T13:13:41.244Z        info    service@v0.138.0/service.go:222 Starting otelcol-mackerel...    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "Version": "0.2.0", "NumCPU": 14}
2025-11-04T13:13:41.244Z        info    extensions/extensions.go:41     Starting extensions...  {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}}
2025-11-04T13:13:41.245Z        info    internal/resourcedetection.go:137       began detecting resource information    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "resourcedetection", "otelcol.component.kind": "processor", "otelcol.pipeline.id": "metrics", "otelcol.signal": "metrics"}
2025-11-04T13:13:41.246Z        info    internal/resourcedetection.go:188       detected resource information   {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "resourcedetection", "otelcol.component.kind": "processor", "otelcol.pipeline.id": "metrics", "otelcol.signal": "metrics", "resource": {"host.name":"2bbb1dcc8491","os.type":"linux"}}
2025-11-04T13:13:41.246Z        info    otlpreceiver@v0.138.0/otlp.go:121       Starting GRPC server    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "otlp", "otelcol.component.kind": "receiver", "endpoint": "127.0.0.1:4317"}
2025-11-04T13:13:41.246Z        info    otlpreceiver@v0.138.0/otlp.go:179       Starting HTTP server    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "otlp", "otelcol.component.kind": "receiver", "endpoint": "127.0.0.1:4318"}
2025-11-04T13:13:41.246Z        info    service@v0.138.0/service.go:245 Everything is ready. Begin running and processing data. {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}}
```

### Linux (deb Package)

```console
$ # Get the asset URL for the corresponding architecture from the GitHub release page.
$ sudo apt install https://github.com/mackerelio/opentelemetry-collector-mackerel/releases/download/v0.2.0/otelcol-mackerel_0.2.0_linux_amd64.deb
$ echo "MACKEREL_APIKEY=your_api_key" | sudo tee -a /etc/otelcol-mackerel/otelcol-mackerel.conf
$ sudo systemctl status otelcol-mackerel.service  --no-pager --lines=0
● otelcol-mackerel.service - Mackerel OpenTelemetry Collector
     Loaded: loaded (/usr/lib/systemd/system/otelcol-mackerel.service; enabled; preset: disabled)
     Active: active (running) since Tue 2025-11-04 11:17:21 JST; 10h ago
 Invocation: 90d09569ad9147d798cb41e55c14b717
   Main PID: 30011 (otelcol-mackere)
      Tasks: 14 (limit: 10643)
     Memory: 31.9M (peak: 34.4M)
        CPU: 4.152s
     CGroup: /system.slice/otelcol-mackerel.service
             └─30011 /usr/bin/otelcol-mackerel --config=mackerel:default
```

### Linux (rpm Package)

```console
$ # Get the asset URL for the corresponding architecture from the GitHub release page.
$ sudo apt install https://github.com/mackerelio/opentelemetry-collector-mackerel/releases/download/v0.2.0/otelcol-mackerel_0.2.0_linux_amd64.rpm
$ echo "MACKEREL_APIKEY=your_api_key" | sudo tee -a /etc/otelcol-mackerel/otelcol-mackerel.conf
$ sudo systemctl status otelcol-mackerel.service  --no-pager --lines=0
● otelcol-mackerel.service - Mackerel OpenTelemetry Collector
     Loaded: loaded (/usr/lib/systemd/system/otelcol-mackerel.service; enabled; preset: disabled)
     Active: active (running) since Tue 2025-11-04 11:17:21 JST; 10h ago
 Invocation: 90d09569ad9147d798cb41e55c14b717
   Main PID: 30011 (otelcol-mackere)
      Tasks: 14 (limit: 10643)
     Memory: 31.9M (peak: 34.4M)
        CPU: 4.152s
     CGroup: /system.slice/otelcol-mackerel.service
             └─30011 /usr/bin/otelcol-mackerel --config=mackerel:default
```

## Options

You can modify the config from the default by setting additional environment variables.

| Environment Variable    | Description                                                     |
| ----------------------- | --------------------------------------------------------------- |
| `OTELCOL_MACKEREL_HOST` | bind address or hostname for OTLP receiver (default: localhost) |

## Advanced Usage

You can override the default configuration provided by Mackerel OpenTelemetry Collector.

### Docker

```console
$ cat config.yaml
receivers:
  otlp:
    protocols:
      grpc:
      http:

exporters:
  mackerelotlp:

service:
  pipelines:
    metrics:
      receivers: [otlp]
      exporters: [mackerelotlp]
    traces:
      receivers: [otlp]
      exporters: [mackerelotlp]
$ docker run -e MACKEREL_APIKEY=your_api_key --mount type=bind,src=./config.yaml,dst=/home/nonroot/config.yaml mackerel/otelcol-mackerel:latest --config /home/nonroot/config.yaml validate
$ docker run -e MACKEREL_APIKEY=your_api_key --mount type=bind,src=./config.yaml,dst=/home/nonroot/config.yaml mackerel/otelcol-mackerel:latest --config /home/nonroot/config.yaml
```

### Linux

```console
$ cat /etc/otelcol-mackerel/config.yaml
receivers:
  otlp:
    protocols:
      grpc:
      http:

exporters:
  mackerelotlp:

service:
  pipelines:
    metrics:
      receivers: [otlp]
      exporters: [mackerelotlp]
    traces:
      receivers: [otlp]
      exporters: [mackerelotlp]
$ MACKEREL_APIKEY="your_api_key" otelcol-mackerel --config=/etc/otelcol-mackerel/config.yaml validate
$ sudo vim /etc/otelcol-mackerel/otelcol-mackerel.conf
```

```diff
-OTELCOL_MACKEREL_OPTIONS="--config=mackerel:default"
+OTELCOL_MACKEREL_OPTIONS="--config=/etc/otelcol-mackerel/config.yaml"
```

```console
$ sudo systemctl restart otelcol-mackerel.service
```

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
