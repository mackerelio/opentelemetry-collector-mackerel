package zerocfgprovider

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.yaml.in/yaml/v3"
)

func TestConfigGenerator_Generate(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		g := newConfigGenerator()
		rawCfg := g.Generate()
		gotYAMLCfg, err := yaml.Marshal(rawCfg)
		require.NoError(t, err)
		wantYAMLCfg, err := os.ReadFile("./testdata/00_default_config.yaml")
		require.NoError(t, err)
		assert.YAMLEq(t, string(wantYAMLCfg), string(gotYAMLCfg))
	})

	t.Run("with OTELCOL_MACKEREL_HOST env", func(t *testing.T) {
		t.Setenv("OTELCOL_MACKEREL_HOST", "otel-collector.test")
		g := newConfigGenerator()
		rawCfg := g.Generate()
		gotYAMLCfg, err := yaml.Marshal(rawCfg)
		require.NoError(t, err)
		wantYAMLCfg, err := os.ReadFile("./testdata/01_with_otelcol_host_config.yaml")
		require.NoError(t, err)
		assert.YAMLEq(t, string(wantYAMLCfg), string(gotYAMLCfg))
	})
}
