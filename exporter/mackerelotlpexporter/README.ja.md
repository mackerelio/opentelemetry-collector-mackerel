# Mackerel OTLP エクスポーター

🌎 日本語 | [English](./README.md)

OTLP（OpenTelemetry Protocol）を使用して、OpenTelemetryコレクターからメトリック・トレース・ログをMackerelに投稿するためのエクスポーターです。

MackerelはOTLPをネイティブサポートしており、OpenTelemetryコミュニティが提供する[OTLP gRPC Exporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlpexporter)や[OTLP HTTP Exporter](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlphttpexporter)を使用してテレメトリーを投稿することができます。

しかし、Mackerelはテレメトリーの種類ごと異なるエンドポイント・異なるOTLPのトランスポート種別を使用しています。複数のエクスポーターを併用する必要があり、手間がかかります。

Mackerel OTLP エクスポーターを使用することで、一つのエクスポーターでメトリック・トレース・ログを投稿することができます。

また、Mackerelの仕様に合わせて適切なタイムアウトやバッチの設定をデフォルトで適用します。パイプラインに[Batch Processor](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/batchprocessor)を入れる必要がありません。

## はじめかた

Mackerel OTLP エクスポーターを同梱したOpenTelemtryコレクターに以下のような設定を渡してください：

```yaml
exporters:
  mackerel_otlp:
```

> [!WARNING]
> v0.9.0から、設定ファイルで用いるコンポーネント名が`mackerelotlp`から`mackerel_otlp`に変わりました。これは[命名規約](https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/coding-guidelines.md)に準ずるための変更です。非推奨の古い名称も引き続きエイリアスとして使用できますが、将来のリリースで削除される可能性があります。

パイプラインも含めた設定ファイル全体の例を以下に示します:

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
    logs:
      receivers: [otlp]
      exporters: [mackerel_otlp]
```

`MACKEREL_APIKEY`環境変数にMackerelのAPIキー（書き込み可能）をセットしてからコレクターを起動してください。

## 高度な使い方

追加で以下の設定ができます:

- `mackerel_api_key`: MackerelのAPIキー（環境変数以外の方法で指定したいケースで使用してください）
- `timeout`: [タイムアウト設定](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#timeout)
- `sending_queue`: [送信キューの設定](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#sending-queue)
- `retry_on_failure`: [送信失敗時の再実行設定](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#retry-on-failure)

### 環境変数

- `OTELCOL_MACKEREL_PREFER_IPV4`: `1`などの値を設定すると、メトリクス用gRPCエンドポイントのホスト名をIPv4アドレスに名前解決してから接続します。DNSがAAAAレコードを返すもののIPv6の疎通がない環境でgRPC接続に失敗する場合に有用です。gRPCエンドポイントであるメトリックのエンドポイントにのみ影響します。
