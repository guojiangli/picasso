package kmetric

import (
	"github.com/prometheus/client_golang/prometheus"
)

// CounterVecOpts ...
type CounterVecoption struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Labels    []string
}

type counterVec struct {
	*prometheus.CounterVec
}

// NewCounterVec ...
func NewCounterVec(opts *CounterVecoption) *counterVec {
	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      opts.Name,
			Help:      opts.Help,
		}, opts.Labels)
	prometheus.MustRegister(vec)
	return &counterVec{
		CounterVec: vec,
	}
}

// Inc ...
func (counter *counterVec) Inc(labels ...string) {
	counter.WithLabelValues(labels...).Inc()
}

// Add ...
func (counter *counterVec) Add(v float64, labels ...string) {
	counter.WithLabelValues(labels...).Add(v)
}
