# MDOT Test Collector

Build binaries to test the components provided by Mackerel OpenTelemetry Collector Distro.

This distribution artifact is not intended for user use.

## Build and Run

```sh
cd /path/to/distributions/otelcol-test
go mod tidy
go generate ./...
./_build/otelcol-mackerel-test --config ./config.yaml
```
