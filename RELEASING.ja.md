# リリース

## リリース手順

### 1. ルートモジュールのリリース

[Songmu/tagpr](https://github.com/Songmu/tagpr)を利用してルートモジュールのタグを打っています。ルートモジュールのバージョンは、 Mackerel OpenTelemetry Collector など、ライブラリではなく成果物を提供したいときのバージョンを表しています。

#### 1.1 バージョンを決める

tagpr が自動で Release for vx.y.z というタイトルの Pull Request（以後、リリース PR と呼びます）を作ります。デフォルトでは patch バージョンが上がることになっています。

リリース PR の description にあるマージされた PR 一覧を確認し、パッチバージョンを上げるままでいいのか、それともマイナーバージョンを上げた方がいいのかを判断してください。

マイナーバージョンを上げる場合には、リリース PR に tagpr:minor というラベルを付与した上で、CI for main branch を rerun してください。

#### 1.2 リリース PR のマージ

リリース PR をマージすると、自動でメインモジュールにタグが打たれます。また、GoReleaser が動き、 Mackerel OpenTelemetry Collector のビルド成果物が push されます。

### 2. サブモジュールのリリース

このリポジトリはモノレポ構成（複数の go.mod を持っている）になっています。例えば、[exporter/mackerelotlpexporter](./exporter/mackerelotlpexporter)の v0.1.0 をリリースする際には、`exporter/mackerelotlpexporter/v0.1.0`というタグを別途打つ必要があります。

cf.) [Managing module source - The Go Programming Language | Sourcing multiple modules in a single repository](https://go.dev/doc/modules/managing-source#multiple-module-source)

サブモジュールのバージョンは、ユーザーがこのリポジトリで提供されるコンポーネントを組み込んで OpenTelemetry Collector をビルドする際に使われます。

このオペレーションは現在 CI による自動化がされておらず、リリース作業者の手元でスクリプトを実行する必要があります。また、GPG キーを使用した commit への署名がセットアップされている必要があります。これは、[open-telemetry/opentelemetry-go-build-tools/multimod](https://github.com/open-telemetry/opentelemetry-go-build-tools/tree/main/multimod) に依存していることに起因します。

cf.) [コミットに署名する - GitHub Docs](https://docs.github.com/ja/authentication/managing-commit-signature-verification/signing-commits)

#### 2.1 手元にソースコードを用意する

```console
$ cd /path/to/opentelemetry-collector-mackerel
$ git switch main
$ git pull
```

#### 2.2 スクリプトを実行する

以下のようにスクリプトを実行すると、main branch の最新 commit に対して各サブモジュールに対応したタグが付与・push されます。

安全のため、手元に diff が残っていたり、現在のリモートリポジトリの main branch の commit hash と異なったりする場合には失敗して停止します。

```console
$ ./scripts/push-multimod-tags.sh
Using versioning file /Users/mackerelio/Repositories/hatena/mackerelio/opentelemetry-collector-mackerel/versions.yaml
Tagging commit 7b92a7623c53a4507c12aee395318909a4b78186:
exporter/mackerelotlpexporter/v0.2.0
confmap/provider/zerocfgprovider/v0.2.0
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 383 bytes | 383.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
To github.com:mackerelio/opentelemetry-collector-mackerel.git
 * [new tag]         exporter/mackerelotlpexporter/v0.2.0 -> exporter/mackerelotlpexporter/v0.2.0
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 392 bytes | 392.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
To github.com:mackerelio/opentelemetry-collector-mackerel.git
 * [new tag]         confmap/provider/zerocfgprovider/v0.2.0 -> confmap/provider/zerocfgprovider/v0.2.0
```
