package kmetric

import "github.com/prometheus/client_golang/prometheus"

// SummaryVecOpts ...
type SummaryVecoption struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Labels    []string
}

type summaryVec struct {
	*prometheus.SummaryVec
}

// Build ...
func NewSummaryVec(opts *SummaryVecoption) *summaryVec {
	vec := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      opts.Name,
			Help:      opts.Help,
		}, opts.Labels)
	prometheus.MustRegister(vec)
	return &summaryVec{
		SummaryVec: vec,
	}
}

// Observe ...
func (summary *summaryVec) Observe(v float64, labels ...string) {
	summary.WithLabelValues(labels...).Observe(v)
}
