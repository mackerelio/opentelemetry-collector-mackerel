package mackerelotlpexporter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
)

func TestCreateDefaultConfig(t *testing.T) {
	defaultConfig := createDefaultConfig()
	assert.NoError(t, componenttest.CheckConfigStruct(defaultConfig))

	cfg := defaultConfig.(*Config)
	assert.Equal(t, 10*time.Second, cfg.TimeoutConfig.Timeout)
}
