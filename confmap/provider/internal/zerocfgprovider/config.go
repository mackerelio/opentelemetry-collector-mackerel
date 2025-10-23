package zerocfgprovider

import "slices"

type configGenerator struct {
	processors                  map[string]any
	metricsPipelineProcessorIDs *processorIDs
	tracesPipelineProcessorIDs  *processorIDs
}

func newConfigGenerator() *configGenerator {
	return &configGenerator{
		processors:                  map[string]any{},
		metricsPipelineProcessorIDs: &processorIDs{},
		tracesPipelineProcessorIDs:  &processorIDs{},
	}
}

func (g *configGenerator) Generate() map[string]any {
	g.addResourceDetectionProcessor()

	cfg := map[string]any{
		"receivers": map[string]any{
			"otlp": map[string]any{
				"protocols": map[string]any{
					"grpc": nil,
					"http": nil,
				},
			},
		},
		"processors": g.processors,
		"exporters": map[string]any{
			"mackerelotlp": nil,
		},
		"service": map[string]any{
			"pipelines": map[string]any{
				"metrics": map[string]any{
					"receivers":  []string{"otlp"},
					"processors": g.metricsPipelineProcessorIDs.GeneratePipeline(),
					"exporters":  []string{"mackerelotlp"},
				},
				"traces": map[string]any{
					"receivers":  []string{"otlp"},
					"processors": g.tracesPipelineProcessorIDs.GeneratePipeline(),
					"exporters":  []string{"mackerelotlp"},
				},
			},
		},
	}
	return cfg
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
	batchProcessorIDs         []string
	otherProcessorIDs         []string
}

func (p *processorIDs) GeneratePipeline() []string {
	// best practices are provided for the execution order of processors
	// cf.) https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor#recommended-processors
	return slices.Concat(
		p.memoryLimiterProcessorIDs,
		p.samplingProcessorIDs,
		p.sendingSourceProcessorIDs,
		p.batchProcessorIDs,
		p.otherProcessorIDs,
	)
}
