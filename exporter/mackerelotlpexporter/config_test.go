package mackerelotlpexporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	t.Run("valid if Mackerel API key is provided via config", func(t *testing.T) {
		cfg := &Config{
			MackerelApiKey: "dummyapikey",
		}
		assert.NoError(t, cfg.Validate())
	})
	t.Run("valid if Mackerel API key is provided via MACKEREL_APIKEY env", func(t *testing.T) {
		cfg := &Config{}
		t.Setenv("MACKEREL_APIKEY", "dummyapikey")
		assert.NoError(t, cfg.Validate())
	})
	t.Run("valid if Mackerel API key is provided via MACKEREL_API_KEY env", func(t *testing.T) {
		cfg := &Config{}
		t.Setenv("MACKEREL_API_KEY", "dummyapikey")
		assert.NoError(t, cfg.Validate())
	})
	t.Run("valid if Mackerel API keys are provided via MACKEREL_APIKEY and MACKEREL_API_KEY env", func(t *testing.T) {
		cfg := &Config{}
		t.Setenv("MACKEREL_APIKEY", "dummyapikey")
		t.Setenv("MACKEREL_API_KEY", "dummyapikey")
		assert.NoError(t, cfg.Validate())
	})
	t.Run("invalid if no Mackerel API keys are provided", func(t *testing.T) {
		cfg := &Config{}
		assert.Error(t, cfg.Validate())
	})
}
