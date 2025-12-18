package mackerelotlpexporter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

func TestCreateDefaultConfig(t *testing.T) {
	defaultConfig := createDefaultConfig()
	assert.NoError(t, componenttest.CheckConfigStruct(defaultConfig))

	cfg := defaultConfig.(*Config)
	assert.Equal(t, 10*time.Second, cfg.TimeoutConfig.Timeout)
	queueCfg := cfg.QueueConfig.GetOrInsertDefault()
	batchCfg := queueCfg.Batch.GetOrInsertDefault()
	assert.Equal(t, exporterhelper.RequestSizerTypeBytes, batchCfg.Sizer)
	assert.Equal(t, int64(5_000_000), batchCfg.MaxSize)
}
