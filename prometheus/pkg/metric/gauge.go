package metric

import "github.com/prometheus/client_golang/prometheus"

var _ GaugeVec = (*promGaugeVec)(nil)

type GaugeVecOpts Opts

type promGaugeVec struct {
	gauge *prometheus.GaugeVec
}

func NewGaugeVec(cfg *GaugeVecOpts) GaugeVec {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
		}, cfg.Labels)
	// 注册监控指标
	prometheus.MustRegister(vec)
	return &promGaugeVec{
		gauge: vec,
	}
}

func (gauge *promGaugeVec) Set(v float64, labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Set(v)
}

func (gauge *promGaugeVec) Inc(labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Inc()
}

func (gauge *promGaugeVec) Dec(labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Dec()
}

func (gauge *promGaugeVec) Add(v float64, labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Add(v)
}

func (gauge *promGaugeVec) Sub(v float64, labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Sub(v)
}
