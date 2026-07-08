package mackerelotlpexporter

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type mockMetricsReceiver struct {
	pmetricotlp.UnimplementedGRPCServer
	requestsCounter *atomic.Int64
	mux             sync.Mutex
	lastMetadata    metadata.MD
}

func (r *mockMetricsReceiver) Export(ctx context.Context, req pmetricotlp.ExportRequest) (pmetricotlp.ExportResponse, error) {
	r.requestsCounter.Add(1)
	r.mux.Lock()
	defer r.mux.Unlock()
	md, _ := metadata.FromIncomingContext(ctx)
	r.lastMetadata = md
	return pmetricotlp.NewExportResponse(), nil
}

func TestSendMetrics(t *testing.T) {
	t.Parallel()

	ln, err := net.Listen("tcp", "localhost:")
	require.NoError(t, err)
	t.Cleanup(func() { _ = ln.Close })

	srv := grpc.NewServer()
	receiver := &mockMetricsReceiver{
		requestsCounter: new(atomic.Int64),
		mux:             sync.Mutex{},
		lastMetadata:    nil,
	}
	pmetricotlp.RegisterGRPCServer(srv, receiver)
	go func() { _ = srv.Serve(ln) }()
	t.Cleanup(srv.GracefulStop)

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	// Disable queuing to ensure that we execute the request when calling ConsumeTraces
	// otherwise we will not see any errors.
	cfg.QueueConfig = configoptional.None[exporterhelper.QueueBatchConfig]()
	cfg.MetricsEndpoint = ln.Addr().String()
	cfg.MackerelApiKey = "dummy"
	cfg.InSecure = true
	set := exportertest.NewNopSettings(factory.Type())

	exporter, err := factory.CreateMetrics(t.Context(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, exporter)
	t.Cleanup(func() { assert.NoError(t, exporter.Shutdown(t.Context())) })

	host := componenttest.NewNopHost()
	require.NoError(t, exporter.Start(t.Context(), host))

	metrics := pmetric.NewMetrics()
	require.NoError(t, exporter.ConsumeMetrics(t.Context(), metrics))
	assert.Equal(t, int64(1), receiver.requestsCounter.Load())
	require.NotNil(t, receiver.lastMetadata)
	assert.Equal(t, []string([]string{"dummy"}), receiver.lastMetadata.Get("Mackerel-Api-Key"))
}

func TestSendTraces(t *testing.T) {
	t.Parallel()

	requestsCounter := new(atomic.Int64)
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/traces", func(w http.ResponseWriter, r *http.Request) {
		requestsCounter.Add(1)

		assert.Equal(t, "*/*", r.Header.Get("Accept"))
		assert.Equal(t, "dummy", r.Header.Get("Mackerel-Api-Key"))

		resp := ptraceotlp.NewExportResponse()
		bytes, err := resp.MarshalProto()
		assert.NoError(t, err)
		w.Header().Set("Content-Type", "application/x-protobuf")
		_, err = w.Write(bytes)
		assert.NoError(t, err)
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	// Disable queuing to ensure that we execute the request when calling ConsumeTraces
	// otherwise we will not see any errors.
	cfg.QueueConfig = configoptional.None[exporterhelper.QueueBatchConfig]()
	cfg.TracesEndpoint = srv.URL
	cfg.MackerelApiKey = "dummy"
	cfg.InSecure = true
	set := exportertest.NewNopSettings(factory.Type())

	exporter, err := factory.CreateTraces(t.Context(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, exporter)
	t.Cleanup(func() { assert.NoError(t, exporter.Shutdown(t.Context())) })

	host := componenttest.NewNopHost()
	require.NoError(t, exporter.Start(t.Context(), host))

	traces := ptrace.NewTraces()
	require.NoError(t, exporter.ConsumeTraces(t.Context(), traces))
	assert.Equal(t, int64(1), requestsCounter.Load())
}

func TestSendLogs(t *testing.T) {
	t.Parallel()

	requestsCounter := new(atomic.Int64)
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/logs", func(w http.ResponseWriter, r *http.Request) {
		requestsCounter.Add(1)

		assert.Equal(t, "*/*", r.Header.Get("Accept"))
		assert.Equal(t, "dummy", r.Header.Get("Mackerel-Api-Key"))

		w.WriteHeader(http.StatusOK)
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	// Disable queuing to ensure that we execute the request when calling ConsumeLogs
	// otherwise we will not see any errors.
	cfg.QueueConfig = configoptional.None[exporterhelper.QueueBatchConfig]()
	cfg.LogsEndpoint = srv.URL
	cfg.MackerelApiKey = "dummy"
	cfg.InSecure = true
	set := exportertest.NewNopSettings(factory.Type())

	exporter, err := factory.CreateLogs(t.Context(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, exporter)
	t.Cleanup(func() { assert.NoError(t, exporter.Shutdown(t.Context())) })

	host := componenttest.NewNopHost()
	require.NoError(t, exporter.Start(t.Context(), host))

	logs := plog.NewLogs()
	require.NoError(t, exporter.ConsumeLogs(t.Context(), logs))
	assert.Equal(t, int64(1), requestsCounter.Load())
}

func TestResolveIPv4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		endpoint         string
		wantOriginalAddr string
		wantIPv4         bool
		wantErr          bool
	}{
		{
			name:             "hostname is resolved to IPv4",
			endpoint:         "localhost:4317",
			wantOriginalAddr: "localhost:4317",
			wantIPv4:         true,
		},
		{
			name:             "IP address is returned unchanged",
			endpoint:         "127.0.0.1:4317",
			wantOriginalAddr: "",
			wantIPv4:         false,
		},
		{
			name:             "IPv6 address is returned unchanged",
			endpoint:         "[::1]:4317",
			wantOriginalAddr: "",
			wantIPv4:         false,
		},
		{
			name:     "no port in endpoint",
			endpoint: "localhost",
			wantErr:  true,
		},
		{
			name:     "unresolvable hostname",
			endpoint: "unresolvable.invalid:4317",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			resolved, originalAddr, err := resolveIPv4(t.Context(), tt.endpoint)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantOriginalAddr, originalAddr)
			if tt.wantIPv4 {
				host, _, err := net.SplitHostPort(resolved)
				require.NoError(t, err)
				ip := net.ParseIP(host)
				require.NotNil(t, ip)
				assert.NotNil(t, ip.To4(), "expected IPv4 address, got %s", host)
			} else {
				assert.Equal(t, tt.endpoint, resolved)
			}
		})
	}
}
