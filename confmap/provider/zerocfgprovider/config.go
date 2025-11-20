package zerocfgprovider

import (
	"net"
	"os"
	"slices"
)

type configGenerator struct {
	receivers                   map[string]any
	processors                  map[string]any
	metricsPipelineReceiverIDs  []string
	metricsPipelineProcessorIDs *processorIDs
	tracesPipelineReceiverIDs   []string
	tracesPipelineProcessorIDs  *processorIDs
}

func newConfigGenerator() *configGenerator {
	return &configGenerator{
		receivers:                   map[string]any{},
		processors:                  map[string]any{},
		metricsPipelineReceiverIDs:  []string{},
		metricsPipelineProcessorIDs: &processorIDs{},
		tracesPipelineReceiverIDs:   []string{},
		tracesPipelineProcessorIDs:  &processorIDs{},
	}
}

func (g *configGenerator) Generate() map[string]any {
	g.addOTLPReceiver()
	g.addResourceDetectionProcessor()

	cfg := map[string]any{
		"receivers":  g.receivers,
		"processors": g.processors,
		"exporters": map[string]any{
			"mackerelotlp": nil,
		},
		"service": map[string]any{
			"pipelines": map[string]any{
				"metrics": map[string]any{
					"receivers":  g.metricsPipelineReceiverIDs,
					"processors": g.metricsPipelineProcessorIDs.GeneratePipeline(),
					"exporters":  []string{"mackerelotlp"},
				},
				"traces": map[string]any{
					"receivers":  g.tracesPipelineReceiverIDs,
					"processors": g.tracesPipelineProcessorIDs.GeneratePipeline(),
					"exporters":  []string{"mackerelotlp"},
				},
			},
		},
	}
	return cfg
}

func (g *configGenerator) addOTLPReceiver() {
	const id = "otlp"
	host := os.Getenv("OTELCOL_MACKEREL_HOST")
	if host == "" {
		host = "localhost"
	}
	g.receivers[id] = map[string]any{
		"protocols": map[string]any{
			"grpc": map[string]any{
				"endpoint": net.JoinHostPort(host, "4317"),
			},
			"http": map[string]any{
				"endpoint": net.JoinHostPort(host, "4318"),
			},
		},
	}
	g.metricsPipelineReceiverIDs = append(g.metricsPipelineReceiverIDs, id)
	g.tracesPipelineReceiverIDs = append(g.tracesPipelineReceiverIDs, id)
}

func (g *configGenerator) addResourceDetectionProcessor() {
	const id = "resourcedetection"
	g.processors[id] = map[string]any{
		// TODO: adapt settings for specific environments such as cloud and container environments
		"detectors": []string{"env", "system"},
		"system": map[string]any{
			"hostname_sources": []string{"os"},
		},
	}
	g.metricsPipelineProcessorIDs.sendingSourceProcessorIDs = append(g.metricsPipelineProcessorIDs.sendingSourceProcessorIDs, id)
	g.tracesPipelineProcessorIDs.sendingSourceProcessorIDs = append(g.tracesPipelineProcessorIDs.sendingSourceProcessorIDs, id)
}

type processorIDs struct {
	memoryLimiterProcessorIDs []string
	samplingProcessorIDs      []string
	sendingSourceProcessorIDs []string
	otherProcessorIDs         []string
}

func (p *processorIDs) GeneratePipeline() []string {
	// best practices are provided for the execution order of processors
	// cf.) https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor#recommended-processors
	return slices.Concat(
		p.memoryLimiterProcessorIDs,
		p.samplingProcessorIDs,
		p.sendingSourceProcessorIDs,
		p.otherProcessorIDs,
	)
}
