package mackerelotlpexporter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

func TestCreateDefaultConfig(t *testing.T) {
	defaultConfig := createDefaultConfig()
	assert.NoError(t, componenttest.CheckConfigStruct(defaultConfig))

	cfg := defaultConfig.(*Config)
	assert.Equal(t, 10*time.Second, cfg.TimeoutConfig.Timeout)

	queueCfg := cfg.QueueConfig.Get()
	require.NotNil(t, queueCfg)

	batchCfg := queueCfg.Batch.Get()
	require.NotNil(t, batchCfg)
	assert.Equal(t, int64(3_500_000), batchCfg.MaxSize)
	assert.Equal(t, exporterhelper.RequestSizerTypeBytes, batchCfg.Sizer)
}
