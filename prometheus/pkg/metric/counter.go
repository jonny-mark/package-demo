package metric

import "github.com/prometheus/client_golang/prometheus"

var _ CounterVec = (*promCounterVec)(nil)

type CounterVecOpts Opts

// promCounterVec counter vec.
type promCounterVec struct {
	counter *prometheus.CounterVec
}

func NewCounterVec(cfg *CounterVecOpts) CounterVec {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
		}, cfg.Labels)
	// 注册监控指标
	prometheus.MustRegister(vec)
	return &promCounterVec{
		counter: vec,
	}
}

func (counter *promCounterVec) Inc(labels ...string) {
	counter.counter.WithLabelValues(labels...).Inc()
}

func (counter *promCounterVec) Add(v float64, labels ...string) {
	counter.counter.WithLabelValues(labels...).Add(v)
}
