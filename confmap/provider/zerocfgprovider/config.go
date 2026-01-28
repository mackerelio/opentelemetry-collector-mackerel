package zerocfgprovider

import (
	"net"
	"slices"

	"github.com/caarlos0/env/v11"
)

type configGenerator struct {
	*cfgEnvs
	receivers                   map[string]any
	processors                  map[string]any
	metricsPipelineReceiverIDs  []string
	metricsPipelineProcessorIDs *processorIDs
	tracesPipelineReceiverIDs   []string
	tracesPipelineProcessorIDs  *processorIDs
}

type cfgEnvs struct {
	Host               string   `env:"OTELCOL_MACKEREL_HOST" envDefault:"localhost"`
	SamplingPercentage *float64 `env:"OTELCOL_MACKEREL_SAMPLING_PERCENTAGE"`
}

func newConfigGenerator() *configGenerator {
	return &configGenerator{
		cfgEnvs:                     nil,
		receivers:                   map[string]any{},
		processors:                  map[string]any{},
		metricsPipelineReceiverIDs:  []string{},
		metricsPipelineProcessorIDs: &processorIDs{},
		tracesPipelineReceiverIDs:   []string{},
		tracesPipelineProcessorIDs:  &processorIDs{},
	}
}

func (g *configGenerator) Generate() (map[string]any, error) {
	envs, err := env.ParseAs[cfgEnvs]()
	if err != nil {
		return nil, err
	}
	g.cfgEnvs = &envs

	g.addOTLPReceiver()
	g.addResourceDetectionProcessor()
	if g.SamplingPercentage != nil {
		g.addProbabilisticSamplingProcessor()
	}

	cfg := map[string]any{
		"receivers":  g.receivers,
		"processors": g.processors,
		"exporters": map[string]any{
			"mackerel_otlp": nil,
		},
		"service": map[string]any{
			"pipelines": map[string]any{
				"metrics": map[string]any{
					"receivers":  g.metricsPipelineReceiverIDs,
					"processors": g.metricsPipelineProcessorIDs.GeneratePipeline(),
					"exporters":  []string{"mackerel_otlp"},
				},
				"traces": map[string]any{
					"receivers":  g.tracesPipelineReceiverIDs,
					"processors": g.tracesPipelineProcessorIDs.GeneratePipeline(),
					"exporters":  []string{"mackerel_otlp"},
				},
			},
		},
	}
	return cfg, nil
}

func (g *configGenerator) addOTLPReceiver() {
	const id = "otlp"
	g.receivers[id] = map[string]any{
		"protocols": map[string]any{
			"grpc": map[string]any{
				"endpoint": net.JoinHostPort(g.Host, "4317"),
			},
			"http": map[string]any{
				"endpoint": net.JoinHostPort(g.Host, "4318"),
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

func (g *configGenerator) addProbabilisticSamplingProcessor() {
	const id = "probabilistic_sampler"
	g.processors[id] = map[string]any{
		"sampling_percentage": g.SamplingPercentage,
	}
	g.tracesPipelineProcessorIDs.samplingProcessorIDs = append(g.tracesPipelineProcessorIDs.samplingProcessorIDs, id)
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
