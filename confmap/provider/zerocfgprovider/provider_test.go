package zerocfgprovider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/confmap"
)

func TestNewFactory(t *testing.T) {
	p := NewFactory().Create(confmap.ProviderSettings{})
	_, ok := p.(*provider)
	assert.True(t, ok)
}

func TestProvider_Retrieve(t *testing.T) {
	p := &provider{}
	_, err := p.Retrieve(context.Background(), "", nil)
	assert.NoError(t, err)
}
