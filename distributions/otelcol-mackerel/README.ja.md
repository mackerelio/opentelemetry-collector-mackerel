# Mackerel OpenTelemetry コレクター

🌎 日本語 | [English](./README.md)

このディストリビューションは Mackerel ユーザーのために作られた OpenTelemetry コレクターのビルドです。

OpenTelemetry コレクターの設定ファイルのフォーマットが分からなくても、いくつかの環境変数をセットするだけでシンプルな設定のコレクターを起動できます。

## はじめかた

ローカルホスト上で OTLP 経由でテレメトリーを受信し、ホストのリソース属性を追加した上で Mackerel にエクスポートする OpenTelemetry コレクターを複雑な設定ファイルなしで起動できます。必要なのは Mackerel の API キーだけです。

### Docker

Docker Hub と GitHub Container Registry にてコンテナイメージを配布しています：

- [mackerel/otelcol-mackerel](https://hub.docker.com/r/mackerel/otelcol-mackerel)
- [ghcr.io/mackerelio/opentelemetry-collector-mackerel/otelcol-mackerel](https://github.com/mackerelio/opentelemetry-collector-mackerel/pkgs/container/opentelemetry-collector-mackerel%2Fotelcol-mackerel)

```
$ docker run -e MACKEREL_APIKEY=your_api_key mackerel/otelcol-mackerel:latest
2025-11-04T13:13:41.242Z        info    builders/builders.go:26 Development component. May change in the future.        {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "mackerel_otlp", "otelcol.component.kind": "exporter", "otelcol.signal": "metrics"}
2025-11-04T13:13:41.243Z        info    builders/builders.go:26 Development component. May change in the future.        {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "mackerel_otlp", "otelcol.component.kind": "exporter", "otelcol.signal": "traces"}
2025-11-04T13:13:41.244Z        info    service@v0.138.0/service.go:222 Starting otelcol-mackerel...    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "Version": "0.2.0", "NumCPU": 14}
2025-11-04T13:13:41.244Z        info    extensions/extensions.go:41     Starting extensions...  {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}}
2025-11-04T13:13:41.245Z        info    internal/resourcedetection.go:137       began detecting resource information    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "resourcedetection", "otelcol.component.kind": "processor", "otelcol.pipeline.id": "metrics", "otelcol.signal": "metrics"}
2025-11-04T13:13:41.246Z        info    internal/resourcedetection.go:188       detected resource information   {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "resourcedetection", "otelcol.component.kind": "processor", "otelcol.pipeline.id": "metrics", "otelcol.signal": "metrics", "resource": {"host.name":"2bbb1dcc8491","os.type":"linux"}}
2025-11-04T13:13:41.246Z        info    otlpreceiver@v0.138.0/otlp.go:121       Starting GRPC server    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "otlp", "otelcol.component.kind": "receiver", "endpoint": "127.0.0.1:4317"}
2025-11-04T13:13:41.246Z        info    otlpreceiver@v0.138.0/otlp.go:179       Starting HTTP server    {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}, "otelcol.component.id": "otlp", "otelcol.component.kind": "receiver", "endpoint": "127.0.0.1:4318"}
2025-11-04T13:13:41.246Z        info    service@v0.138.0/service.go:245 Everything is ready. Begin running and processing data. {"resource": {"service.instance.id": "ec2e6d20-2fb6-4017-b21e-cea7a01df4d7", "service.name": "otelcol-mackerel", "service.version": "0.2.0"}}
```

### Linux (deb パッケージ)

```sh
curl -fsSL https://mackerel.io/file/script/opentelemetry-collector-mackerel/setup-apt.sh | MACKEREL_APIKEY='<YOUR_API_KEY>' sh
```

<details>
<summary>マニュアルインストール</summary>

```console
$ # GitHubのリリースページから対応するアーキテクチャのasset URLを入手してください
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

</details>

### Linux (rpm パッケージ)

```sh
curl -fsSL https://mackerel.io/file/script/opentelemetry-collector-mackerel/setup-yum.sh | MACKEREL_APIKEY='<YOUR_API_KEY>' sh
```

<details>
<summary>マニュアルインストール</summary>

```console
$ # GitHubのリリースページから対応するアーキテクチャのasset URLを入手してください
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

</details>

## 追加設定

追加の環境変数をセットすることでデフォルトの設定を変更することができます。

| 環境変数                               | 説明                                                                      |
| -------------------------------------- | ------------------------------------------------------------------------- |
| `OTELCOL_MACKEREL_HOST`                | OTLP レシーバーが受信する IP アドレスやホスト名 (デフォルト値: localhost) |
| `OTELCOL_MACKEREL_SAMPLING_PERCENTAGE` | 指定したパーセンテージでトレースに確率的サンプリングを適用します          |

## 高度な使い方

Mackerel OpenTelemetry コレクターが提供するデフォルト設定を上書きして、標準の方法で設定を記述することもできます。

### Docker

```console
$ cat config.yaml
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
  mackerel_otlp:

service:
  pipelines:
    metrics:
      receivers: [otlp]
      exporters: [mackerel_otlp]
    traces:
      receivers: [otlp]
      exporters: [mackerel_otlp]
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

## コンポーネント

OpenTelemetry コミュニティが提供する OpenTelemetry コレクターコンポーネントの中から、特に有用なものをいくつか選択してバンドルしています。

もし Mackerel ユーザーの方でこのディストリビューションに追加してほしいコンポーネントがあれば、GitHub で Issue を開いてリクエストしてください。

### エクスポーター

| コンポーネント                          | 説明                   | ドキュメント                                                                                                       |
| --------------------------------------- | ---------------------- | ------------------------------------------------------------------------------------------------------------------ |
| `debug`                                 | Debug Exporter         | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/debugexporter)             |
| `nop`                                   | No-op Exporter         | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/nopexporter)               |
| `otlp_grpc` (alias: `otlp`)             | OTLP/gRPC Exporter     | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlpexporter)              |
| `otlp_http` (alias: `otlphttp`)         | OTLP/HTTP Exporter     | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlphttpexporter)          |
| `mackerel_otlp` (alias: `mackerelotlp`) | Mackerel OTLP Exporter | [Document](https://github.com/mackerelio/opentelemetry-collector-mackerel/tree/main/exporter/mackerelotlpexporter) |

### プロセッサー

| コンポーネント          | 説明                             | ドキュメント                                                                                                                    |
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

### レシーバー

| コンポーネント | 説明                  | ドキュメント                                                                                                         |
| -------------- | --------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `nop`          | No-op Receiver        | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/receiver/nopreceiver)                 |
| `otlp`         | OTLP Receiver         | [Document](https://github.com/open-telemetry/opentelemetry-collector/tree/main/receiver/otlpreceiver)                |
| `awsxray`      | AWS X-Ray Receiver    | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/awsxrayreceiver)     |
| `hostmetrics`  | Host Metrics Receiver | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver) |
| `httpcheck`    | HTTP Check Receiver   | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/httpcheckreceiver)   |
| `mysql`        | MySQL Receiver        | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/mysqlreceiver)       |
| `oracledb`     | Oracle DB Receiver    | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/oracledbreceiver)    |
| `postgresql`   | PostgreSQL Receiver   | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/postgresqlreceiver)  |
| `redis`        | Redis Receiver        | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/redisreceiver)       |

### コネクター

| コンポーネント | 説明                    | ドキュメント                                                                                                            |
| -------------- | ----------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `routing`      | Routing Connector       | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/connector/routingconnector)      |
| `servicegraph` | Service Graph Connector | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/connector/servicegraphconnector) |
| `spanmetrics`  | Span Metrics Connector  | [Document](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/connector/spanmetricsconnector)  |
