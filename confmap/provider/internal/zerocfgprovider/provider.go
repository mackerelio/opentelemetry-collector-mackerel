package zerocfgprovider

import (
	"context"

	"go.opentelemetry.io/collector/confmap"
)

const schemeName = "mackerel"

type provider struct{}

func newProvider(confmap.ProviderSettings) confmap.Provider {
	return &provider{}
}

func NewFactory() confmap.ProviderFactory {
	return confmap.NewProviderFactory(newProvider)
}

func (p *provider) Scheme() string {
	return schemeName
}

func (p *provider) Retrieve(_ context.Context, _ string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
	configGenerator := newConfigGenerator()
	rawCfg := configGenerator.Generate()
	return confmap.NewRetrieved(rawCfg)
}

func (p *provider) Shutdown(context.Context) error {
	return nil
}
