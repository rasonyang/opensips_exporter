package processors

import (
	"github.com/VoIPGRID/opensips_exporter/opensips"
	"github.com/prometheus/client_golang/prometheus"
)

// URIProcessor metrics related to SIP URI processing.
// doc: http://www.opensips.org/html/docs/modules/1.11.x/uri.html
// src: https://github.com/OpenSIPS/opensips/blob/1.11/modules/uri/uri_mod.c#L191
type URIProcessor struct {
	statistics map[string]opensips.Statistic
}

var uriLabelNames = []string{}
var uriMetrics = map[string]metric{
	"positive checks": newMetric("uri", "positive_checks", "Amount of positive URI checks.", uriLabelNames, prometheus.CounterValue),
	"negative_checks": newMetric("uri", "negative_checks", "Amount of negative URI checks.", uriLabelNames, prometheus.CounterValue),
}

func init() {
	for metric := range uriMetrics {
		Processors[metric] = uriProcessorFunc
	}
	Processors["uri:"] = uriProcessorFunc
}

// Describe implements prometheus.Collector.
func (p URIProcessor) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range uriMetrics {
		ch <- metric.Desc
	}
}

// Collect implements prometheus.Collector.
func (p URIProcessor) Collect(ch chan<- prometheus.Metric) {
	for key, metric := range uriMetrics {
		ch <- prometheus.MustNewConstMetric(
			metric.Desc,
			metric.ValueType,
			p.statistics[key].Value,
		)
	}
}

func uriProcessorFunc(s map[string]opensips.Statistic) prometheus.Collector {
	return &URIProcessor{
		statistics: s,
	}
}
